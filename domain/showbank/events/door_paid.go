package events

type DoorPaid struct {
	ShowID         string
	AmountInCents  int
	BalanceInCents int
}

func NewDoorPaid(showID string, amountInCents, balanceInCents int) DoorPaid {
	return DoorPaid{
		ShowID:         showID,
		AmountInCents:  amountInCents,
		BalanceInCents: balanceInCents,
	}
}
