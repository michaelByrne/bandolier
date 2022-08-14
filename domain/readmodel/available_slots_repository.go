package readmodel

import "time"

type AvailableSlotsRepository interface {
	Add(AvailableSlot)
	Delete(string)
	MarkAsUnavailable(string)
	MarkAsAvailable(string)
	GetSlotsAvailableOn(time time.Time) []*AvailableSlot
	Clear()
}
