package readmodel

import (
	"time"
)

type AvailableSlot struct {
	SlotID    string
	ShowID    string
	Date      string
	Start     string
	Duration  time.Duration
	Available bool
}

type AvailableSlots []*AvailableSlot

func NewAvailableSlot(id string, start string, date string, d time.Duration, available bool, showID string) *AvailableSlot {
	return &AvailableSlot{
		SlotID:    id,
		Start:     start,
		Date:      date,
		Duration:  d,
		Available: available,
		ShowID:    showID,
	}
}
