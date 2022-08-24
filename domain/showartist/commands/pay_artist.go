package commands

type PayArtist struct {
	ArtistID      string
	AmountInCents int
	ShowID        string
}

func NewPayArtist(artistID string, amountInCents int, showID string) PayArtist {
	return PayArtist{
		ArtistID:      artistID,
		ShowID:        showID,
		AmountInCents: amountInCents,
	}
}
