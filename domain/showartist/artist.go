package showartist

import (
	"bandolier/domain/showartist/events"
	"bandolier/eventsourcing"
)

type Artist struct {
	eventsourcing.AggregateRootBase

	isCreated bool
	name      string
}

func NewArtist() *Artist {
	a := &Artist{
		AggregateRootBase: eventsourcing.NewAggregateRoot(),
	}

	a.Register(events.ArtistCreated{}, func(e interface{}) { a.ArtistCreated(e.(events.ArtistCreated)) })

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

// EVENTS

func (a *Artist) ArtistCreated(e events.ArtistCreated) {
	a.isCreated = true
	a.name = e.Name
	a.Id = e.ID
}
