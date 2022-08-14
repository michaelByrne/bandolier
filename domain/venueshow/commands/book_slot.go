package commands

import "time"

type BookSlot struct {
	ID         string
	ArtistID   string
	ArtistName string
	Headliner  bool
	VenueID    string
	Start      time.Time
}

func NewBookSlot(id, artistID, venueID string, start time.Time, artistName string, headliner bool) BookSlot {
	return BookSlot{
		ID:         id,
		ArtistID:   artistID,
		ArtistName: artistName,
		Headliner:  headliner,
		VenueID:    venueID,
		Start:      start,
	}
}
