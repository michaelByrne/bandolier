package application

import (
	"bandolier/domain/showartist/events"
	"bandolier/domain/showbank/commands"
	"bandolier/infrastructure"
)

type PayArtistProcessManager struct {
	infrastructure.EventHandlerBase
}

func NewPayArtistProcessManager(
	c infrastructure.CommandStore,
) *PayArtistProcessManager {
	h := infrastructure.NewEventHandler()

	h.When(events.Paid{}, func(e interface{}, _ infrastructure.EventMetadata) error {
		d := e.(events.Paid)

		err := c.Send(commands.NewPayArtist(d.AmountInCents, d.ShowID, d.ArtistID), infrastructure.CommandMetadata{})
		if err != nil {
			return err
		}

		return nil
	})

	return &PayArtistProcessManager{h}
}
