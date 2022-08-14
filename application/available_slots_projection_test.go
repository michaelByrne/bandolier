package application_test

import (
	"bandolier/application"
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/inmemory"
	projections2 "bandolier/infrastructure/projections"
	"testing"
	"time"
)

func TestNewAvailableSlotsProjection(t *testing.T) {
	r := inmemory.NewAvailableSlotsRepository()
	slotID := "e30c4b41-816c-4ad3-adff-385ff7ddf875"
	venueID := "d1148e6e-1da2-45a8-a4ea-b4c808cb079f"
	artistID := "67789e34-401f-455c-be03-4c39d932a3ff"
	startTime := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	duration := time.Minute * 45
	showID := venueshow.NewShowID(venueID, startTime)

	p := &AvailableSlotsTests{
		ProjectionTests: projections2.NewProjectionTests(t, func() projections2.Projection {
			return application.NewAvailableSlotsProjection(r)
		}),
		repository: r,
		slotID:     slotID,
		venueID:    venueID,
		startTime:  startTime,
		duration:   duration,
		showID:     showID,
		artistID:   artistID,
	}

	t.Run("ShouldAddSlot", p.ShouldAddSlot)
	t.Run("ShouldMarkSlotAsUnavailable", p.ShouldMarkSlotAsUnavailable)

}

func (s AvailableSlotsTests) ShouldAddSlot(t *testing.T) {
	s.Given(
		events.NewShowScheduled(venueshow.NewShowID(s.venueID, s.startTime).Value, s.venueID, s.startTime),
		events.NewSlotScheduled(s.slotID, s.startTime, s.duration),
	)
	s.Then(
		[]*readmodel.AvailableSlot{readmodel.NewAvailableSlot(s.slotID, s.startTime, s.duration, true)},
		s.repository.GetSlotsAvailableOn(s.startTime))
}

func (s AvailableSlotsTests) ShouldMarkSlotAsUnavailable(t *testing.T) {
	s.Given(
		events.NewShowScheduled(venueshow.NewShowID(s.venueID, s.startTime).Value, s.venueID, s.startTime),
		events.NewSlotScheduled(s.slotID, s.startTime, s.duration),
		events.NewSlotBooked(s.slotID, s.showID.Value, s.artistID),
	)
	s.Then(
		[]*readmodel.AvailableSlot{readmodel.NewAvailableSlot(s.slotID, s.startTime, s.duration, false)},
		s.repository.GetSlotsAvailableOn(s.startTime))
}

type AvailableSlotsTests struct {
	projections2.ProjectionTests

	repository readmodel.AvailableSlotsRepository
	slotID     string
	venueID    string
	artistID   string
	showID     venueshow.ShowID
	startTime  time.Time
	duration   time.Duration
}
