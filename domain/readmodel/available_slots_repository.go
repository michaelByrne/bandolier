package readmodel

import "time"

type AvailableSlotsRepository interface {
	Add(AvailableSlot) error
	Delete(string)
	MarkAsUnavailable(string) error
	MarkAsAvailable(string) error
	GetSlotsAvailableOn(time time.Time) ([]*AvailableSlot, error)
	Clear()
}
