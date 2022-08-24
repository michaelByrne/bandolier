package showartist

type Payment struct {
	ArtistID      string
	AmountInCents int
	ShowID        string
}

func NewPayment(artistID string, amountInCents int, showID string) Payment {
	return Payment{
		ArtistID:      artistID,
		ShowID:        showID,
		AmountInCents: amountInCents,
	}
}
