package commands

import "time"

type BookSlot struct {
	ID       string
	ArtistID string
	VenueID  string
	Start    time.Time
}

func NewBookSlot(id string, artistID string, venueID string, start time.Time) BookSlot {
	return BookSlot{
		ID:       id,
		ArtistID: artistID,
		VenueID:  venueID,
		Start:    start,
	}
}
