package commands

type ReceiveCovers struct {
	ShowID        string
	AmountInCents int
}

func NewReceiveCovers(showID string, amountInCents int) ReceiveCovers {
	return ReceiveCovers{
		ShowID:        showID,
		AmountInCents: amountInCents,
	}
}
