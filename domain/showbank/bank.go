package showbank

import (
	"bandolier/domain/showbank/events"
	"bandolier/eventsourcing"
)

type Bank struct {
	eventsourcing.AggregateRootBase

	doorPaid       bool
	artistsPaid    bool
	isOpened       bool
	coversReceived bool
	balanceInCents int
}

func NewBank() *Bank {
	a := &Bank{
		AggregateRootBase: eventsourcing.NewAggregateRoot(),
	}

	a.Register(events.DoorPaid{}, func(e interface{}) { a.DoorPaid(e.(events.DoorPaid)) })
	a.Register(events.BankOpened{}, func(e interface{}) { a.BankOpened(e.(events.BankOpened)) })
	a.Register(events.CoversReceived{}, func(e interface{}) { a.CoversReceived(e.(events.CoversReceived)) })

	return a
}

// COMMANDS

func (b *Bank) OpenBank(showID string, presaleInCents int) error {
	if b.isOpened {
		return BankAlreadyOpenedError{}
	}

	newBalance := presaleInCents
	b.Raise(events.NewBankOpened(showID, newBalance))
	return nil
}

func (b *Bank) PayDoor(amount int, showID string) error {
	if !b.isOpened {
		return BankNotOpenedError{}
	}

	err := b.isDoorPaid()
	if err != nil {
		return err
	}

	newBalance := b.balanceInCents - amount
	b.Raise(events.NewDoorPaid(showID, amount, newBalance))
	return nil
}

func (b *Bank) ReceiveCovers(amount int, showID string) error {
	if !b.isOpened {
		return BankNotOpenedError{}
	}

	newBalance := b.balanceInCents + amount
	b.Raise(events.NewCoversReceived(showID, amount, newBalance))
	return nil
}

// EVENTS

func (b *Bank) DoorPaid(e events.DoorPaid) {
	b.doorPaid = true
	b.balanceInCents = e.BalanceInCents
}

func (b *Bank) BankOpened(e events.BankOpened) {
	b.Id = e.ShowID
	b.balanceInCents = e.PresaleInCents
	b.isOpened = true
}

func (b *Bank) CoversReceived(e events.CoversReceived) {
	b.coversReceived = true
	b.balanceInCents = e.NewBalance
}

// HELPERS

func (b *Bank) isDoorPaid() error {
	if b.doorPaid {
		return &DoorAlreadyPaidError{}
	}

	return nil
}
