package showartist

import (
	"bandolier/domain/showartist/commands"
	"bandolier/infrastructure"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repo ArtistRepository) CommandHandlers {
	commandHandlers := CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.CreateArtist{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.CreateArtist)

		id := NewArtistID(cmd.Name, cmd.Date)
		artist, err := repo.Get(id.Value)
		if err != nil {
			return err
		}

		err = artist.CreateArtist(cmd.Name, id.Value)
		if err != nil {
			return err
		}

		repo.Save(artist)
		return nil
	})

	commandHandlers.Register(commands.PayArtist{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.PayArtist)

		artist, err := repo.Get(cmd.ArtistID)
		if err != nil {
			return err
		}

		err = artist.PayArtist(cmd.ArtistID, cmd.ShowID, cmd.AmountInCents)
		if err != nil {
			return err
		}

		repo.Save(artist)
		return nil
	})

	return commandHandlers
}
