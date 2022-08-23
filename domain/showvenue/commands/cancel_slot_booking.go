package commands

import "time"

type CancelSlotBooking struct {
	ID      string
	VenueID string
	Date    time.Time
}

func NewCancelSlotBooking(id string, venueID string, date time.Time) CancelSlotBooking {
	return CancelSlotBooking{
		ID:      id,
		VenueID: venueID,
		Date:    date,
	}
}
