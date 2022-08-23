package main

import (
	"bandolier/application"
	"bandolier/controllers"
	"bandolier/domain/showbank"
	"bandolier/domain/venueshow"
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

	mongoDatabase := mongoClient.Database("projections")
	availableSlotsRepo := mongodb.NewAvailableSlotsRepository(mongoDatabase)
	showDetailRepo := mongodb.NewShowDetailRepository(mongoDatabase)
	showBankRepo := mongodb.NewBankBalanceRepository(mongoDatabase)
	schedulingDispatcher := getShowDispatcher(schedulingEventStore)
	bankDispatcher := getBankDispatcher(bankEventStore)
	commandStore := infrastructure.NewEsCommandStore(schedulingEventStore, esdbClient, serde, schedulingDispatcher)

	showArchiver := application.NewShowArchiverProcessManager(
		commandStore,
	)

	err = venueshow.RegisterTypes(typeMapper)
	if err != nil {
		panic(err)
	}

	err = showbank.RegisterTypes(typeMapper)
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

	//commandStore := infrastructure.NewEsCommandStore(schedulingEventStore, esdbClient, serde, schedulingDispatcher)

	bookingController := controllers.NewBookingController(availableSlotsRepo, showDetailRepo, schedulingDispatcher, schedulingEventStore)
	bankController := controllers.NewBankController(bankDispatcher, bankEventStore)
	e := echo.New()
	bookingController.Register(e)
	bankController.Register(e)

	err = commandStore.Start()
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
	showRepository := venueshow.NewEventStoreShowRepository(aggregateStore)
	handlers := venueshow.NewHandlers(showRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func getBankDispatcher(eventStore infrastructure.EventStore) *infrastructure.Dispatcher {
	aggregateStore := infrastructure.NewEsAggregateStore(eventStore, 5)
	showRepository := showbank.NewEventStoreBankRepository(aggregateStore)
	handlers := showbank.NewHandlers(showRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func createMongoClient() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoConnectionString))
}
