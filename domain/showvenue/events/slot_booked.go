package events

type SlotBooked struct {
	ID         string
	ShowID     string
	ArtistID   string
	ArtistName string
	Headliner  bool
}

func NewSlotBooked(id, showID, artistID, artistName string, headliner bool) SlotBooked {
	return SlotBooked{
		ID:         id,
		ShowID:     showID,
		ArtistID:   artistID,
		ArtistName: artistName,
		Headliner:  headliner,
	}
}
