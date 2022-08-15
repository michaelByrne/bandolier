package events

type BankOpened struct {
	ShowID         string
	PresaleInCents int
}

func NewBankOpened(showID string, presaleInCents int) BankOpened {
	return BankOpened{
		ShowID:         showID,
		PresaleInCents: presaleInCents,
	}
}
