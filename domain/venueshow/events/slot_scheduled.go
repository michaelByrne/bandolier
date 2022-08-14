package events

import "time"

type SlotScheduled struct {
	ID        string
	StartTime time.Time
	Duration  time.Duration
}

func NewSlotScheduled(id string, startTime time.Time, duration time.Duration) SlotScheduled {
	return SlotScheduled{
		ID:        id,
		StartTime: startTime,
		Duration:  duration,
	}
}
