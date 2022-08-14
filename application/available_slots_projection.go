package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/infrastructure/projections"
)

type AvailableSlotsProjection struct {
	projections.ProjectionBase

	repository readmodel.AvailableSlotsRepository
}

func NewAvailableSlotsProjection(r readmodel.AvailableSlotsRepository) *AvailableSlotsProjection {
	p := projections.NewProjection()
	p.When(events.SlotScheduled{}, func(e interface{}) {
		s := e.(events.SlotScheduled)
		r.Add(*readmodel.NewAvailableSlot(s.ID, s.StartTime, s.Duration, true))
	})
	p.When(events.SlotBooked{}, func(e interface{}) {
		b := e.(events.SlotBooked)
		r.MarkAsUnavailable(b.ID)
	})

	return &AvailableSlotsProjection{p, r}
}
