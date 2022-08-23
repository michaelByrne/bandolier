package commands

import "time"

type ArchiveShow struct {
	VenueID           string
	Date              time.Time
	DoorAmountInCents int
}

func NewArchiveShow(venueID string, date time.Time, amountInCents int) ArchiveShow {
	return ArchiveShow{
		VenueID:           venueID,
		Date:              date,
		DoorAmountInCents: amountInCents,
	}
}
