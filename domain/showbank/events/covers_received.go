package events

type CoversReceived struct {
	ShowID        string
	AmountInCents int
	NewBalance    int
}

func NewCoversReceived(showID string, amountInCents, newBalance int) CoversReceived {
	return CoversReceived{
		ShowID:        showID,
		AmountInCents: amountInCents,
		NewBalance:    newBalance,
	}
}
