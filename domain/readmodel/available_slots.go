package readmodel

import "time"

type AvailableSlot struct {
	SlotID    string
	Start     time.Time
	Duration  time.Duration
	Available bool
}

type AvailableSlots []*AvailableSlot

func NewAvailableSlot(id string, s time.Time, d time.Duration, available bool) *AvailableSlot {
	return &AvailableSlot{
		SlotID:    id,
		Start:     s,
		Duration:  d,
		Available: available,
	}
}
