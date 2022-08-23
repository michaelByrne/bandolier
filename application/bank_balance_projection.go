package application

import (
	"bandolier/domain/readmodel"
	"bandolier/domain/showbank/events"
	"bandolier/infrastructure/mongodb"
	"bandolier/infrastructure/projections"
)

type BankBalanceProjection struct {
	projections.ProjectionBase

	readmodel.BankBalanceRepository
}

func NewBankBalanceProjection(r *mongodb.BankBalanceRepository) *BankBalanceProjection {
	p := projections.NewProjection()
	p.When(events.BankOpened{}, func(e interface{}) error {
		b := e.(events.BankOpened)
		err := r.OpenBank(b.ShowID, b.PresaleInCents)
		if err != nil {
			return err
		}

		return nil
	})
	p.When(events.CoversReceived{}, func(e interface{}) error {
		b := e.(events.CoversReceived)
		err := r.UpdateBalance(b.ShowID, b.NewBalance)
		if err != nil {
			return err
		}

		return nil
	})
	p.When(events.DoorPaid{}, func(e interface{}) error {
		b := e.(events.DoorPaid)
		err := r.UpdateBalance(b.ShowID, b.BalanceInCents)
		if err != nil {
			return err
		}

		return nil
	})

	return &BankBalanceProjection{p, r}
}
