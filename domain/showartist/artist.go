package showartist

import (
	"bandolier/domain/showartist/events"
	"bandolier/eventsourcing"
)

type Artist struct {
	eventsourcing.AggregateRootBase

	isCreated bool
	name      string
	payments  []*Payment
}

func NewArtist() *Artist {
	a := &Artist{
		AggregateRootBase: eventsourcing.NewAggregateRoot(),
	}

	a.Register(events.ArtistCreated{}, func(e interface{}) { a.ArtistCreated(e.(events.ArtistCreated)) })
	a.Register(events.Paid{}, func(e interface{}) { a.Paid(e.(events.Paid)) })

	return a
}

// COMMANDS

func (a *Artist) CreateArtist(name string, id string) error {
	if a.isCreated {
		return ArtistAlreadyCreatedError{}
	}

	a.Raise(events.NewArtistCreated(name, id))
	return nil
}

func (a *Artist) PayArtist(artistID, showID string, amountInCents int) error {
	if !a.isCreated {
		return ArtistNotCreatedError{}
	}

	if a.paymentExists(showID, artistID) {
		return ArtistAlreadyPaidForShowError{}
	}

	a.Raise(events.NewPaid(amountInCents, showID, artistID))
	return nil
}

// EVENTS

func (a *Artist) ArtistCreated(e events.ArtistCreated) {
	a.isCreated = true
	a.name = e.Name
	a.Id = e.ID
}

func (a *Artist) Paid(e events.Paid) {
	a.payments = append(a.payments, &Payment{
		AmountInCents: e.AmountInCents,
		ShowID:        e.ShowID,
		ArtistID:      e.ArtistID,
	})
}

// HELPERS

func (a *Artist) paymentExists(showID, ArtistID string) bool {
	for _, p := range a.payments {
		if p.ShowID == showID && p.ArtistID == ArtistID {
			return true
		}
	}

	return false
}
