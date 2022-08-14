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
		r.Add(*readmodel.NewAvailableSlot(s.ID, s.StartTime, s.Duration, true))

		return nil
	})
	p.When(events.SlotBooked{}, func(e interface{}) error {
		b := e.(events.SlotBooked)
		r.MarkAsUnavailable(b.ID)

		return nil
	})
	p.When(events.SlotBookingCancelled{}, func(e interface{}) error {
		c := e.(events.SlotBookingCancelled)
		r.MarkAsAvailable(c.ID)

		return nil
	})

	return &AvailableSlotsProjection{p, r}
}
