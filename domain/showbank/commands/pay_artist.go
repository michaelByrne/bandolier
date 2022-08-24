package commands

type PayArtist struct {
	AmountInCents int
	ShowID        string
	ArtistID      string
}

func NewPayArtist(amountInCents int, showID string, artistID string) PayArtist {
	return PayArtist{
		AmountInCents: amountInCents,
		ShowID:        showID,
		ArtistID:      artistID,
	}
}
