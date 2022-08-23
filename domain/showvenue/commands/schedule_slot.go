package commands

import "time"

type ScheduleSlot struct {
	ID       string
	Start    time.Time
	Duration time.Duration
	VenueID  string
}

func NewScheduleSlot(id string, start time.Time, duration time.Duration, venueID string) ScheduleSlot {
	return ScheduleSlot{
		ID:       id,
		Start:    start,
		Duration: duration,
		VenueID:  venueID,
	}
}
