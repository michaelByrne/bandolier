package commands

type BookSlot struct {
	ID         string
	ArtistID   string
	ArtistName string
	Headliner  bool
	ShowID     string
}

func NewBookSlot(id, artistID, showID string, artistName string, headliner bool) BookSlot {
	return BookSlot{
		ID:         id,
		ArtistID:   artistID,
		ArtistName: artistName,
		Headliner:  headliner,
		ShowID:     showID,
	}
}
