package commands

type PayDoor struct {
	StaffID       string
	ShowID        string
	AmountInCents int
}

func NewPayDoor(showID string, amountInCents int) PayDoor {
	return PayDoor{
		ShowID:        showID,
		AmountInCents: amountInCents,
	}
}
