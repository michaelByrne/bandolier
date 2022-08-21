package mongodb

import (
	"bandolier/domain/readmodel"
	"time"
)

type AvailableSlot struct {
	ID        string
	ShowID    string
	Date      string
	StartTime string
	Duration  time.Duration
	IsBooked  bool
}

func (a *AvailableSlot) ToAvailableSlot() *readmodel.AvailableSlot {
	return &readmodel.AvailableSlot{
		SlotID:    a.ID,
		ShowID:    a.ShowID,
		Start:     a.StartTime,
		Date:      a.Date,
		Duration:  a.Duration,
		Available: !a.IsBooked,
	}
}

func FromAvailableSlot(a readmodel.AvailableSlot) *AvailableSlot {
	return &AvailableSlot{
		ID:        a.SlotID,
		ShowID:    a.ShowID,
		StartTime: a.Start,
		Date:      a.Date,
		Duration:  a.Duration,
		IsBooked:  !a.Available,
	}
}
