package showbank_test

import (
	"bandolier/domain/showbank"
	"bandolier/domain/showbank/commands"
	"bandolier/domain/showbank/events"
	"bandolier/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewBank(t *testing.T) {
	store := infrastructure.NewFakeAggregateStore()
	repo := showbank.NewEventStoreBankRepository(store)
	showID := "05b2f393-7d9e-48e8-be24-48e8d2fc114e"
	artistID := "1c2afcd6-0989-4151-8576-6aaf2614f6d4"
	oneDollar := 100
	twoDollars := 200
	fiveDollars := 500
	tenDollars := 1000
	zeroDollars := 0

	a := BankTests{
		AggregateTests: infrastructure.NewAggregateTests(store),
		showID:         showID,
		oneDollar:      oneDollar,
		twoDollars:     twoDollars,
		fiveDollars:    fiveDollars,
		tenDollars:     tenDollars,
		zeroDollars:    zeroDollars,
		artistID:       artistID,
	}

	a.RegisterHandlers(showbank.NewHandlers(repo))

	t.Run("ShouldOpenBankWithPresale", a.ShouldOpenBankWithPresale)
	t.Run("ShouldNotOpenBankIfAlreadyOpened", a.ShouldNotOpenBankIfAlreadyOpened)
	t.Run("ShouldPayDoor", a.ShouldPayDoor)
	t.Run("ShouldNotPayDoorIfBankNotOpened", a.ShouldNotPayDoorIfBankNotOpened)
	t.Run("ShouldReceiveCovers", a.ShouldReceiveCovers)
	t.Run("ShouldReceiveCoversAndPayDoor", a.ShouldReceiveCoversAndPayDoor)
	t.Run("ShouldPayArtist", a.ShouldPayArtist)
	t.Run("ShouldPayArtistAndPayDoor", a.ShouldPayArtistAndPayDoor)
}

func (b BankTests) ShouldOpenBankWithPresale(t *testing.T) {
	b.Given()
	b.When(commands.NewOpenBank(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewBankOpened(b.showID, b.oneDollar), changes[0])
	})
}

func (b BankTests) ShouldNotOpenBankIfAlreadyOpened(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.oneDollar))
	b.When(commands.NewOpenBank(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, showbank.BankAlreadyOpenedError{}, err)
	})
}

func (b BankTests) ShouldPayDoor(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.oneDollar))
	b.When(commands.NewPayDoor(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewDoorPaid(b.showID, b.oneDollar, b.zeroDollars), changes[0])
	})
}

func (b BankTests) ShouldNotPayDoorIfBankNotOpened(t *testing.T) {
	b.Given()
	b.When(commands.NewPayDoor(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, showbank.BankNotOpenedError{}, err)
	})
}

func (b BankTests) ShouldReceiveCovers(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.oneDollar))
	b.When(commands.NewReceiveCovers(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewCoversReceived(b.showID, b.oneDollar, b.twoDollars), changes[0])
	})
}

func (b BankTests) ShouldReceiveCoversAndPayDoor(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.oneDollar), events.NewCoversReceived(b.showID, b.oneDollar, b.twoDollars))
	b.When(commands.NewPayDoor(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewDoorPaid(b.showID, b.oneDollar, b.oneDollar), changes[0])
	})
}

func (b BankTests) ShouldPayArtist(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.oneDollar))
	b.When(commands.NewPayArtist(b.oneDollar, b.showID, b.artistID))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewArtistPaid(b.oneDollar, b.zeroDollars, b.showID, b.artistID), changes[0])
	})
}

func (b BankTests) ShouldPayArtistAndPayDoor(t *testing.T) {
	b.Given(events.NewBankOpened(b.showID, b.twoDollars), events.NewArtistPaid(b.oneDollar, b.oneDollar, b.showID, b.artistID))
	b.When(commands.NewPayDoor(b.showID, b.oneDollar))
	b.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewDoorPaid(b.showID, b.oneDollar, b.zeroDollars), changes[0])
	})
}

type BankTests struct {
	infrastructure.AggregateTests

	showID      string
	artistID    string
	oneDollar   int
	twoDollars  int
	fiveDollars int
	tenDollars  int
	zeroDollars int
}
