package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/projections"
)

type AvailableSlotsProjection struct {
	projections.ProjectionBase

	repository readmodel.AvailableSlotsRepository
}

func NewAvailableSlotsProjection(r readmodel.AvailableSlotsRepository) *AvailableSlotsProjection {
	p := projections.NewProjection()
	p.When(events.SlotScheduled{}, func(e interface{}) error {
		s := e.(events.SlotScheduled)
		startTime := s.StartTime.Format("15:04")
		date := s.StartTime.Format("2006-01-02")
		err := r.Add(*readmodel.NewAvailableSlot(s.ID, startTime, date, s.Duration, true, s.ShowID))
		if err != nil {
			return err
		}

		return nil
	})
	p.When(events.SlotBooked{}, func(e interface{}) error {
		b := e.(events.SlotBooked)
		err := r.MarkAsUnavailable(b.ID)
		if err != nil {
			return err
		}

		return nil
	})
	p.When(events.SlotBookingCancelled{}, func(e interface{}) error {
		c := e.(events.SlotBookingCancelled)
		err := r.MarkAsAvailable(c.ID)
		if err != nil {
			return err
		}

		return nil
	})

	return &AvailableSlotsProjection{p, r}
}
