package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/venueshow/events"
	"bandolier/infrastructure/inmemory"
	"bandolier/infrastructure/projections"
	"errors"
)

type ShowDetailProjection struct {
	projections.ProjectionBase

	repository readmodel.ShowDetailRepository
}

func NewShowDetailProjection(r readmodel.ShowDetailRepository) *ShowDetailProjection {
	p := projections.NewProjection()

	p.When(events.ShowScheduled{}, func(e interface{}) error {
		s := e.(events.ShowScheduled)

		r.AddShow(*readmodel.NewShowDetail(s.ID, "", s.Date, s.VenueID, []readmodel.Slot{}))

		return nil
	})
	p.When(events.SlotScheduled{}, func(e interface{}) error {
		s := e.(events.SlotScheduled)

		return r.AddSlotToShow(s.ID, s.ShowID, s.StartTime, s.Duration)
	})
	p.When(events.SlotBooked{}, func(e interface{}) error {
		b := e.(events.SlotBooked)

		if b.Headliner {
			err := r.SetHeadliner(b.ShowID, b.ArtistName)
			if err != nil {
				if !errors.Is(err, inmemory.HeadlinerAlreadySetError{}) {
					return err
				}
			}
		}

		slot, err := r.GetSlot(b.ShowID, b.ID)
		if err != nil {
			return err
		}

		return r.AddSlotBookingToShow(b.ShowID, *readmodel.NewBooking(b.ID, b.ArtistID, b.ArtistName, slot.Start, slot.Duration))
	})

	return &ShowDetailProjection{p, r}
}
