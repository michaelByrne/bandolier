package events

import "time"

type ShowScheduled struct {
	ID      string
	VenueID string
	Date    time.Time
}

func NewShowScheduled(id string, venueID string, date time.Time) ShowScheduled {
	return ShowScheduled{
		ID:      id,
		VenueID: venueID,
		Date:    date,
	}
}
