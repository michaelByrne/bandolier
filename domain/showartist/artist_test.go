package showartist

import (
	"bandolier/domain/showartist/commands"
	"bandolier/domain/showartist/events"
	"bandolier/domain/showvenue"
	"bandolier/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestArtistAggregate(t *testing.T) {
	store := infrastructure.NewFakeAggregateStore()
	registry := NewEventStoreArtistRepository(store)
	artistID := NewArtistID("Yo La Tengo", time.Now()).Value
	showID := showvenue.NewShowID("73b4e6dd-35b7-430f-8c5c-45f6f485914b", time.Now()).Value
	fiveDollars := 500
	tenDollars := 1000
	fifteenDollars := 1500
	now := time.Now()

	a := ArtistTests{
		AggregateTests: infrastructure.NewAggregateTests(store),
		artistID:       artistID,
		showID:         showID,
		fiveDollars:    fiveDollars,
		tenDollars:     tenDollars,
		fifteenDollars: fifteenDollars,
		date:           now,
	}

	a.RegisterHandlers(NewHandlers(registry))

	t.Run("ArtistShouldBeCreated", a.ArtistShouldBeCreated)
	t.Run("ArtistShouldNotBeCreatedTwice", a.ArtistShouldNotBeCreatedTwice)
	t.Run("ArtistShouldBePaid", a.ArtistShouldBePaid)
	t.Run("ArtistShouldNotBePaidTwiceForTheSameShow", a.ArtistShouldNotBePaidTwiceForTheSameShow)
	t.Run("ArtistShouldBePaidForMultipleShows", a.ArtistShouldBePaidForMultipleShows)
}

func (a ArtistTests) ArtistShouldBeCreated(t *testing.T) {
	a.Given()
	a.When(
		commands.NewCreateArtist("Yo La Tengo", a.date),
	)
	a.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewArtistCreated("Yo La Tengo", a.artistID), changes[0])
	})
}

func (a ArtistTests) ArtistShouldNotBeCreatedTwice(t *testing.T) {
	a.Given(
		events.NewArtistCreated("Yo La Tengo", a.artistID),
	)
	a.When(
		commands.NewCreateArtist("Yo La Tengo", a.date),
	)
	a.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, ArtistAlreadyCreatedError{}, err)
	})
}

func (a ArtistTests) ArtistShouldBePaid(t *testing.T) {
	a.Given(
		events.NewArtistCreated("Yo La Tengo", a.artistID),
	)
	a.When(
		commands.NewPayArtist(a.artistID, a.fiveDollars, a.showID),
	)
	a.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewPaid(a.fiveDollars, a.showID, a.artistID), changes[0])
	})
}

func (a ArtistTests) ArtistShouldNotBePaidTwiceForTheSameShow(t *testing.T) {
	a.Given(
		events.NewArtistCreated("Yo La Tengo", a.artistID),
		events.NewPaid(a.fiveDollars, a.showID, a.artistID),
	)
	a.When(
		commands.NewPayArtist(a.artistID, a.tenDollars, a.showID),
	)
	a.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, ArtistAlreadyPaidForShowError{}, err)
	})
}

func (a ArtistTests) ArtistShouldBePaidForMultipleShows(t *testing.T) {
	a.Given(
		events.NewArtistCreated("Yo La Tengo", a.artistID),
		events.NewPaid(a.fiveDollars, a.showID, a.artistID),
	)
	a.When(
		commands.NewPayArtist(a.artistID, a.tenDollars, "73b4e6dd-35b7-430f-8c5c-45f6f485914b"),
	)
	a.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewPaid(a.tenDollars, "73b4e6dd-35b7-430f-8c5c-45f6f485914b", a.artistID), changes[0])
	})
}

type ArtistTests struct {
	infrastructure.AggregateTests

	artistID       string
	showID         string
	fiveDollars    int
	tenDollars     int
	fifteenDollars int
	date           time.Time
}
