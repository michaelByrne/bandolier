package showbank

import (
	"bandolier/domain/showbank/commands"
	"bandolier/infrastructure"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repo BankRepository) CommandHandlers {
	commandHandlers := CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.PayDoor{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.PayDoor)

		account, err := repo.Get(cmd.ShowID)
		if err != nil {
			return err
		}

		err = account.PayDoor(cmd.AmountInCents, cmd.ShowID)
		if err != nil {
			return err
		}

		repo.Save(account)
		return nil
	})

	commandHandlers.Register(commands.OpenBank{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.OpenBank)

		account, err := repo.Get(cmd.ShowID)
		if err != nil {
			return err
		}

		err = account.OpenBank(cmd.ShowID, cmd.PresaleInCents)
		if err != nil {
			return err
		}

		repo.Save(account)
		return nil
	})

	commandHandlers.Register(commands.ReceiveCovers{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.ReceiveCovers)

		account, err := repo.Get(cmd.ShowID)
		if err != nil {
			return err
		}

		err = account.ReceiveCovers(cmd.AmountInCents, cmd.ShowID)
		if err != nil {
			return err
		}

		repo.Save(account)
		return nil
	})

	commandHandlers.Register(commands.PayArtist{}, func(c infrastructure.Command, metadata infrastructure.CommandMetadata) error {
		cmd := c.(commands.PayArtist)

		account, err := repo.Get(cmd.ShowID)
		if err != nil {
			return err
		}

		err = account.PayArtist(cmd.AmountInCents, cmd.ShowID, cmd.ArtistID)
		if err != nil {
			return err
		}

		repo.Save(account)
		return nil
	})

	return commandHandlers
}
