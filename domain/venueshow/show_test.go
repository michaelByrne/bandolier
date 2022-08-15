package venueshow_test

import (
	"bandolier/domain/venueshow"
	"bandolier/domain/venueshow/commands"
	"bandolier/domain/venueshow/events"
	infrastructure "bandolier/infrastructure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestShowAggregate(t *testing.T) {
	store := infrastructure.NewFakeAggregateStore()
	registry := venueshow.NewEventStoreShowRepository(store)
	venueID := "8cf31856-dc3e-497a-ade5-3ef6978603af"
	firstSlotID := "90e23890-85e0-45a7-a6d0-059d7c280534"
	secondSlotID := "44bcd742-4798-4316-b731-586dfa07f8af"
	venue := venueshow.NewVenue(venueID, "Floristree")
	firstStart := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	secondStart := time.Date(2018, time.January, 1, 0, 30, 0, 0, time.UTC)
	showID := venueshow.NewShowID(venueID, firstStart)
	artist := venueshow.NewArtist("77452a30-5873-4e60-afc3-68a14bf7ea60", "Celebration", false)
	oneHundredCents := 100

	a := ShowTests{
		AggregateTests:  infrastructure.NewAggregateTests(store),
		venue:           venue,
		firstStart:      firstStart,
		secondStart:     secondStart,
		showID:          showID,
		firstSlotID:     firstSlotID,
		secondSlotID:    secondSlotID,
		artist:          artist,
		oneHour:         time.Hour,
		thirtyMinutes:   time.Minute * time.Duration(30),
		oneHundredCents: oneHundredCents,
	}

	a.RegisterHandlers(venueshow.NewHandlers(registry))

	t.Run("ShowShouldBeScheduled", a.ShowShouldBeScheduled)
	t.Run("ShouldBeScheduledWithManySlots", a.ShouldBeScheduledWithManySlots)
	t.Run("SlotShouldBeScheduled", a.SlotShouldBeScheduled)
	t.Run("OverlappingSlotShouldNotBeScheduled", a.OverlappingSlotShouldNotBeScheduled)
	t.Run("ShouldBookSlot", a.ShouldBookSlot)
	t.Run("ShouldNotBookSlotWithBadID", a.ShouldNotBookSlotWithBadID)
	t.Run("ShouldNotDoubleBookSlot", a.ShouldNotDoubleBookSlot)
	t.Run("ShouldNotBookSlotWithoutAScheduledShow", a.ShouldNotBookSlotWithoutAScheduledShow)
	t.Run("ShouldCancel", a.ShouldCancel)
	t.Run("ShouldCancelShowWithSlotBookings", a.ShouldCancelShowWithSlotBookings)
	t.Run("ShouldArchiveShow", a.ShouldArchiveShow)
	t.Run("ShouldNotArchiveArchivedShow", a.ShouldNotArchiveArchivedShow)
	t.Run("ShouldNotArchiveShowThatDoesNotExist", a.ShouldNotArchiveShowThatDoesNotExist)

}

func (s ShowTests) ShowShouldBeScheduled(t *testing.T) {
	s.Given()
	s.When(commands.NewScheduleShow(s.venue.ID, s.venue.Name, s.firstStart, []commands.ScheduledSlot{}))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), changes[0])
	})

}

func (s ShowTests) ShouldBeScheduledWithManySlots(t *testing.T) {
	var slots []commands.ScheduledSlot
	slots = make([]commands.ScheduledSlot, 30)
	tenMinutes := time.Minute * 10
	for i, _ := range slots {
		slots[i] = commands.NewScheduledSlot(s.firstStart.Add(tenMinutes*time.Duration(i)), tenMinutes)
	}

	s.Given()
	s.When(commands.NewScheduleShow(s.venue.ID, s.venue.Name, s.firstStart, slots))
	s.Then(func(changes []interface{}, err error) {
		assert.NoError(t, err)
		assert.Equal(t, events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), changes[0])
		assert.Len(t, changes, 31)
	})
}

func (s ShowTests) ShouldNotBeScheduledIfShowCancelled(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewShowScheduleCancelled(s.showID.Value))
	s.When(commands.NewScheduleShow(s.venue.ID, s.venue.Name, s.firstStart, []commands.ScheduledSlot{}))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.ShowScheduleAlreadyCancelledError{}, err)
	})
}

func (s ShowTests) SlotShouldBeScheduled(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart))
	s.When(commands.NewScheduleSlot(s.firstSlotID, s.firstStart, s.oneHour, s.venue.ID))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour), changes[0])
	})
}

func (s ShowTests) OverlappingSlotShouldNotBeScheduled(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour))
	s.When(commands.NewScheduleSlot(s.showID.Value, s.secondStart, s.thirtyMinutes, s.venue.ID))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.SlotOverlapsError{}, err)
	})
}

func (s ShowTests) ShouldBookSlot(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour))
	s.When(commands.NewBookSlot(s.firstSlotID, s.artist.ID, s.venue.ID, s.firstStart, s.artist.Name, false))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewSlotBooked(s.firstSlotID, s.showID.Value, s.artist.ID, s.artist.Name, false), changes[0])
	})
}

func (s ShowTests) ShouldNotBookSlotWithBadID(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour))
	s.When(commands.NewBookSlot("b97db0bf-1341-47eb-8f73-3bfef1188e29", s.artist.ID, s.venue.ID, s.firstStart, s.artist.Name, false))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.SlotDoesNotExistError{}, err)
	})
}

func (s ShowTests) ShouldNotDoubleBookSlot(t *testing.T) {
	s.Given(
		events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart),
		events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour),
		events.NewSlotBooked(s.firstSlotID, s.showID.Value, s.artist.ID, s.artist.Name, false),
	)
	s.When(commands.NewBookSlot(s.firstSlotID, s.artist.ID, s.venue.ID, s.firstStart, s.artist.Name, false))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.SlotAlreadyBookedError{}, err)
	})
}

func (s ShowTests) ShouldNotBookSlotWithoutAScheduledShow(t *testing.T) {
	s.Given()
	s.When(commands.NewBookSlot(s.firstSlotID, s.artist.ID, s.venue.ID, s.firstStart, s.artist.Name, false))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.ShowNotScheduledError{}, err)
	})
}

func (s ShowTests) ShouldCancel(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour))
	s.When(commands.NewCancel(s.venue.ID, s.firstStart))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewSlotScheduleCancelled(s.firstSlotID), changes[0])
	})
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewShowScheduleCancelled(s.showID.Value), changes[1])
	})
}

func (s ShowTests) ShouldCancelShowWithSlotBookings(t *testing.T) {
	s.Given(
		events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart),
		events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour),
		events.NewSlotBooked(s.firstSlotID, s.showID.Value, s.artist.ID, s.artist.Name, false),
	)
	s.When(commands.NewCancel(s.venue.ID, s.firstStart))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewSlotBookingCancelled(s.firstSlotID), changes[0])
	})
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewSlotScheduleCancelled(s.firstSlotID), changes[1])
	})
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewShowScheduleCancelled(s.showID.Value), changes[2])
	})
}

func (s ShowTests) ShouldArchiveShow(t *testing.T) {
	s.Given(events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart), events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour))
	s.When(commands.NewArchiveShow(s.venue.ID, s.firstStart, s.oneHundredCents))
	s.Then(func(changes []interface{}, err error) {
		require.NoError(t, err)
		assert.Equal(t, events.NewShowArchived(s.showID.Value, s.oneHundredCents), changes[0])
	})
}

func (s ShowTests) ShouldNotArchiveArchivedShow(t *testing.T) {
	s.Given(
		events.NewShowScheduled(s.showID.Value, s.venue.ID, s.firstStart),
		events.NewSlotScheduled(s.firstSlotID, s.showID.Value, s.firstStart, s.oneHour),
		events.NewShowArchived(s.showID.Value, s.oneHundredCents),
	)
	s.When(commands.NewArchiveShow(s.venue.ID, s.firstStart, s.oneHundredCents))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.ShowAlreadyArchivedError{}, err)
	})
}

func (s ShowTests) ShouldNotArchiveShowThatDoesNotExist(t *testing.T) {
	s.Given()
	s.When(commands.NewArchiveShow(s.venue.ID, s.firstStart, s.oneHundredCents))
	s.Then(func(changes []interface{}, err error) {
		require.Error(t, err)
		assert.Equal(t, venueshow.ShowNotScheduledError{}, err)
	})
}

type ShowTests struct {
	infrastructure.AggregateTests

	showID          venueshow.ShowID
	venue           venueshow.Venue
	artist          venueshow.Artist
	firstSlotID     string
	secondSlotID    string
	firstStart      time.Time
	secondStart     time.Time
	oneHour         time.Duration
	thirtyMinutes   time.Duration
	oneHundredCents int
}
