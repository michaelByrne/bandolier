package venueshow

import (
	"bandolier/domain/venueshow/commands"
	"bandolier/infrastructure"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repo ShowRepository) CommandHandlers {
	commandHandlers := CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.ScheduleShow{}, func(c infrastructure.Command) error {
		cmd := c.(commands.ScheduleShow)

		id := NewShowID(cmd.VenueID, cmd.Date)
		show, err := repo.Get(id)
		if err != nil {
			return err
		}

		err = show.ScheduleShow(Venue{ID: cmd.VenueID, Name: cmd.VenueName}, cmd.Date, cmd.Slots)
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	commandHandlers.Register(commands.ScheduleSlot{}, func(c infrastructure.Command) error {
		cmd := c.(commands.ScheduleSlot)
		show, err := repo.Get(NewShowID(cmd.VenueID, cmd.Start))
		if err != nil {
			return err
		}

		err = show.ScheduleSlot(cmd.ID, cmd.Start, cmd.Duration)
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	commandHandlers.Register(commands.BookSlot{}, func(c infrastructure.Command) error {
		cmd := c.(commands.BookSlot)
		show, err := repo.Get(NewShowID(cmd.VenueID, cmd.Start))
		if err != nil {
			return err
		}

		err = show.BookSlot(cmd.ID, cmd.ArtistID, cmd.ArtistName, cmd.Headliner)
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	commandHandlers.Register(commands.CancelSlotBooking{}, func(c infrastructure.Command) error {
		cmd := c.(commands.CancelSlotBooking)
		show, err := repo.Get(NewShowID(cmd.VenueID, cmd.Date))
		if err != nil {
			return err
		}

		err = show.CancelSlotBooking(cmd.ID)
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	commandHandlers.Register(commands.Cancel{}, func(c infrastructure.Command) error {
		cmd := c.(commands.Cancel)
		show, err := repo.Get(NewShowID(cmd.VenueID, cmd.Date))
		if err != nil {
			return err
		}

		err = show.Cancel()
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	commandHandlers.Register(commands.ArchiveShow{}, func(c infrastructure.Command) error {
		cmd := c.(commands.ArchiveShow)
		show, err := repo.Get(NewShowID(cmd.VenueID, cmd.Date))
		if err != nil {
			return err
		}

		err = show.Archive(cmd.DoorAmountInCents)
		if err != nil {
			return err
		}

		repo.Save(show)
		return nil
	})

	return commandHandlers
}
