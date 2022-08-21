package main

import (
	"bandolier/application"
	"bandolier/controllers"
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
	eventStore := infrastructure.NewEsEventStore(esdbClient, "scheduling", serde)
	mongoDatabase := mongoClient.Database("projections")
	availableSlotsRepo := mongodb.NewAvailableSlotsRepository(mongoDatabase)
	showDetailRepo := mongodb.NewShowDetailRepository(mongoDatabase)

	err = venueshow.RegisterTypes(typeMapper)
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
	)

	err = subManager.Start(context.TODO())
	if err != nil {
		panic(err)
	}

	dispatcher := getDispatcher(eventStore)
	//commandStore := infrastructure.NewEsCommandStore(eventStore, esdbClient, serde, dispatcher)

	bookingController := controllers.NewBookingController(availableSlotsRepo, showDetailRepo, dispatcher, eventStore)
	e := echo.New()
	bookingController.Register(e)

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

func getDispatcher(eventStore infrastructure.EventStore) *infrastructure.Dispatcher {
	aggregateStore := infrastructure.NewEsAggregateStore(eventStore, 5)
	showRepository := venueshow.NewEventStoreShowRepository(aggregateStore)
	handlers := venueshow.NewHandlers(showRepository)
	cmdHandlerMap := infrastructure.NewCommandHandlerMap(handlers)
	dispatcher := infrastructure.NewDispatcher(cmdHandlerMap)
	return &dispatcher
}

func createMongoClient() (*mongo.Client, error) {
	return mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoConnectionString))
}
