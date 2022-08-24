package events

type Paid struct {
	AmountInCents int
	ShowID        string
	ArtistID      string
}

func NewPaid(amountInCents int, showID string, artistID string) Paid {
	return Paid{
		AmountInCents: amountInCents,
		ShowID:        showID,
		ArtistID:      artistID,
	}
}
