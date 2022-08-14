package events

import "time"

type SlotScheduled struct {
	ID        string
	ShowID    string
	StartTime time.Time
	Duration  time.Duration
}

func NewSlotScheduled(id, showID string, startTime time.Time, duration time.Duration) SlotScheduled {
	return SlotScheduled{
		ID:        id,
		StartTime: startTime,
		Duration:  duration,
		ShowID:    showID,
	}
}
