package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/projections"
)

type ShowDetailProjection struct {
	projections.ProjectionBase

	repository readmodel.ShowDetailRepository
}

func NewShowDetailProjection(r readmodel.ShowDetailRepository) *ShowDetailProjection {
	p := projections.NewProjection()

	p.When(events.ShowScheduled{}, func(e interface{}) error {
		s := e.(events.ShowScheduled)

		err := r.AddShow(*readmodel.NewShowDetail(s.ID, "", s.Date, s.VenueID, []readmodel.Slot{}))
		if err != nil {
			return err
		}

		return nil
	})
	p.When(events.SlotScheduled{}, func(e interface{}) error {
		s := e.(events.SlotScheduled)

		return r.AddSlot(*readmodel.NewSlot(s.ID, s.StartTime, s.Duration, s.ShowID))
	})
	p.When(events.SlotBooked{}, func(e interface{}) error {
		b := e.(events.SlotBooked)

		return r.AddBooking(*readmodel.NewBooking(b.ID, b.ArtistID, b.ArtistName, b.Headliner))
	})

	return &ShowDetailProjection{p, r}
}
