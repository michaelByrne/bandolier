package events

type SlotBooked struct {
	ID       string
	ShowID   string
	ArtistID string
}

func NewSlotBooked(id string, showID string, artistID string) SlotBooked {
	return SlotBooked{
		ID:       id,
		ShowID:   showID,
		ArtistID: artistID,
	}
}
