package commands

import (
	"time"
)

type ScheduleShow struct {
	VenueID   string
	VenueName string
	Date      time.Time
	Slots     []ScheduledSlot
}

type ScheduledSlot struct {
	StartTime time.Time
	Duration  time.Duration
}

func NewScheduleShow(venueID, venueName string, date time.Time, slots []ScheduledSlot) ScheduleShow {
	return ScheduleShow{
		VenueID:   venueID,
		Date:      date,
		Slots:     slots,
		VenueName: venueName,
	}
}

func NewScheduledSlot(start time.Time, duration time.Duration) ScheduledSlot {
	return ScheduledSlot{
		StartTime: start,
		Duration:  duration,
	}
}
