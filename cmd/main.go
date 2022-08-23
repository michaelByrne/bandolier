package main

import (
	"bandolier/application"
	"bandolier/controllers"
	"bandolier/domain/showartist"
	"bandolier/domain/showbank"
	"bandolier/domain/showvenue"
	"bandolier/eventsourcing"
	"bandolier/infrastructure"
	"bandolier/infrastructure/mongodb"
	"bandolier/infrastructure/projections"
	"context"
	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const (
	EsdbConnectionString  = "esdb://localhost:2113?tls=false"
	MongoConnectionString = "mongodb://localhost"
)

func main() {
	esdbClient, err := createESDBClient()
	if err != nil {
		panic(err)
	}

	mongoClient, err := createMongoClient()
	if err != nil {
		panic(err)
	}

	typeMapper := eventsourcing.NewTypeMapper()
	serde := infrastructure.NewEsEventSerde(typeMapper)
	schedulingEventStore := infrastructure.NewEsEventStore(esdbClient, "scheduling", serde)
	bankEventStore := infrastructure.NewEsEventStore(esdbClient, "bank", serde)
	artistEventStore := infrastructure.NewEsEventStore(esdbClient, "artist", serde)

	mongoDatabase := mongoClient.Database("projections")
	availableSlotsRepo := mongodb.NewAvailableSlotsRepository(mongoDatabase)
	showDetailRepo := mongodb.NewShowDetailRepository(mongoDatabase)
	showBankRepo := mongodb.NewBankBalanceRepository(mongoDatabase)
	schedulingDispatcher := getShowDispatcher(schedulingEventStore)
	bankDispatcher := getBankDispatcher(bankEventStore)
	artistDispatcher := getArtistDispatcher(artistEventStore)
	bankCommandStore := infrastructure.NewEsCommandStore(schedulingEventStore, esdbClient, serde, bankDispatcher)

	showArchiver := application.NewShowArchiverProcessManager(
		bankCommandStore,
	)

	err = showvenue.RegisterTypes(typeMapper)
	if err != nil {
		panic(err)
	}

	err = showbank.RegisterTypes(typeMapper)
	if err != nil {
		panic(err)
	}

	err = showartist.RegisterTypes(typeMapper)
	if err != nil {
		panic(err)
	}

	subManager := projections.NewSubscriptionManager(
		esdbClient,
		infrastructure.NewEsCheckpointStore(esdbClient, "DaySubscription", serde),
		serde,
		"$all",
		projections.NewProjector(application.NewAvailableSlotsProjection(availableSlotsRepo)),
		projections.NewProjector(application.NewShowDetailProjection(showDetailRepo)),
		projections.NewProjector(showArchiver),
		projections.NewProjector(application.NewBankBalanceProjection(showBankRepo)),
	)

	err = subManager.Start(context.TODO())
	if err != nil {
		panic(err)
	}

	//bankCommandStore := infrastructure.NewEsCommandStore(schedulingEventStore, esdbClient, serde, schedulingDispatcher)

	bookingController := controllers.NewBookingController(availableSlotsRepo, showDetailRepo, schedulingDispatcher, schedulingEventStore)
	bankController := controllers.NewBankController(bankDispatcher, bankEventStore, showBankRepo)
	artistController := controllers.NewArtistController(artistDispatcher, artistEventStore)

	e := echo.New()

	bookingController.Register(e)
	bankController.Register(e)
	artistController.Register(e)

	err = bankCommandStore.Start()
	if err != nil {
		panic(err)
	}

	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":5001"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Bandolier!")
}

func createESDBClient() (*esdb.Client, error) {
	settings, err := esdb.ParseConnectionString(EsdbConnectionString)
	if err != nil {
		return nil, err
	}

	db, err := esdb.NewClient(settings)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getShowDispatcher(eventStore infrastructure.EventStore) *infrastructure.Dispatcher {
	aggregateStore := infrastructure.NewEsAggregateStore(eventStore, 5)
	showRepository := showvenue.NewEventStoreShowRepository(aggregateStore)
	handlers := showvenue.NewHandlers(showRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func getBankDispatcher(eventStore infrastructure.EventStore) *infrastructure.Dispatcher {
	aggregateStore := infrastructure.NewEsAggregateStore(eventStore, 5)
	bankRepository := showbank.NewEventStoreBankRepository(aggregateStore)
	handlers := showbank.NewHandlers(bankRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func getArtistDispatcher(eventStore infrastructure.EventStore) *infrastructure.Dispatcher {
	aggregateStore := infrastructure.NewEsAggregateStore(eventStore, 5)
	artistRepository := showartist.NewEventStoreArtistRepository(aggregateStore)
	handlers := showartist.NewHandlers(artistRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func createMongoClient() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoConnectionString))
}
