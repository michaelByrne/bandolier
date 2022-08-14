package commands

import "time"

type Cancel struct {
	VenueID string
	Date    time.Time
}

func NewCancel(venueID string, date time.Time) Cancel {
	return Cancel{
		VenueID: venueID,
		Date:    date,
	}
}
