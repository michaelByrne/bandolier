package application

import (
	"bandolier/domain/showbank/commands"
	"bandolier/domain/venueshow/events"
	"bandolier/eventsourcing"
	"bandolier/infrastructure"
	"time"
)

type ShowArchiverProcessManager struct {
	infrastructure.EventHandlerBase
}

func NewShowArchiverProcessManager(
	s eventsourcing.ColdStorage,
	c infrastructure.CommandStore,
	es infrastructure.EventStore,
	archiveThreshold time.Duration,
) *ShowArchiverProcessManager {
	h := infrastructure.NewEventHandler()

	h.When(events.ShowArchived{}, func(e interface{}, _ infrastructure.EventMetadata) error {
		d := e.(events.ShowArchived)

		err := c.Send(commands.NewReceiveCovers(d.ShowID, d.DoorAmountInCents), infrastructure.CommandMetadata{})
		if err != nil {
			return err
		}

		return nil
	})

	return &ShowArchiverProcessManager{h}
}
