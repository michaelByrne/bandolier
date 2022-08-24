package events

type ArtistPaid struct {
	AmountInCents int
	NewBalance    int
	ShowID        string
	ArtistID      string
}

func NewArtistPaid(amountInCents, newBalance int, showID string, artistID string) ArtistPaid {
	return ArtistPaid{
		AmountInCents: amountInCents,
		ShowID:        showID,
		ArtistID:      artistID,
		NewBalance:    newBalance,
	}
}
