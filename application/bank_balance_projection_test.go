package application_test

import (
	"bandolier/application"
	"bandolier/domain/showbank/events"
	"bandolier/infrastructure"
	"bandolier/infrastructure/mongodb"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestBankBalanceProjection(t *testing.T) {
	showID := "bd2a6fad-89aa-48f2-b080-383addd8448a"
	oneDollar := 100
	fiveDollars := 500
	twentyDollars := 2000
	tenDollars := 1000

	balance := events.NewBankOpened(showID, tenDollars)
	m := map[string]interface{}{}

	b, err := json.Marshal(balance)
	require.NoError(t, err)

	err = json.Unmarshal(b, &m)
	require.NoError(t, err)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost"))
	defer client.Disconnect(context.TODO())
	assert.NoError(t, err)

	p := &BankBalanceTests{
		HandlerTests: infrastructure.NewHandlerTests(t),

		showID:        showID,
		oneDollar:     oneDollar,
		fiveDollars:   fiveDollars,
		twentyDollars: twentyDollars,
		tenDollars:    tenDollars,
	}

	p.SetHandlerFactory(func() infrastructure.EventHandler {
		p.repository = mongodb.NewBankBalanceRepository(client.Database("cool-db"))
		return application.NewBankBalanceProjection(p.repository)
	})

	t.Run("ShouldOpenBankWithPresale", p.ShouldOpenBankWithPresale)
	t.Run("ShouldUpdateBalanceWhenCoversReceived", p.ShouldUpdateBalanceWhenCoversReceived)
	t.Run("ShouldUpdateBalanceWhenDoorPaid", p.ShouldUpdateBalanceWhenDoorPaid)
	t.Run("ShouldUpdateBalanceWhenDoorPaidAndCoversReceived", p.ShouldUpdateBalanceWhenDoorPaidAndCoversReceived)
}

func (b *BankBalanceTests) ShouldOpenBankWithPresale(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.tenDollars))

	balance, err := b.repository.GetBalance(b.showID)
	require.NoError(t, err)

	b.Then(b.tenDollars, balance.BalanceInCents)
}

func (b *BankBalanceTests) ShouldUpdateBalanceWhenCoversReceived(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.tenDollars), events.NewCoversReceived(b.showID, b.fiveDollars, b.twentyDollars))

	balance, err := b.repository.GetBalance(b.showID)
	require.NoError(t, err)

	b.Then(b.twentyDollars, balance.BalanceInCents)
}

func (b *BankBalanceTests) ShouldUpdateBalanceWhenDoorPaid(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.tenDollars), events.NewDoorPaid(b.showID, b.fiveDollars, b.fiveDollars))

	balance, err := b.repository.GetBalance(b.showID)
	require.NoError(t, err)

	b.Then(b.fiveDollars, balance.BalanceInCents)
}

func (b *BankBalanceTests) ShouldUpdateBalanceWhenDoorPaidAndCoversReceived(t *testing.T) {
	b.Given(
		events.NewBankOpened(b.showID, b.tenDollars),
		events.NewDoorPaid(b.showID, b.fiveDollars, b.fiveDollars),
		events.NewCoversReceived(b.showID, b.oneDollar, b.tenDollars),
	)

	balance, err := b.repository.GetBalance(b.showID)
	require.NoError(t, err)

	b.Then(b.tenDollars, balance.BalanceInCents)
}

type BankBalanceTests struct {
	infrastructure.HandlerTests

	repository *mongodb.BankBalanceRepository

	showID        string
	oneDollar     int
	fiveDollars   int
	tenDollars    int
	twentyDollars int
}
