package readmodel

type BankBalance struct {
	ShowID         string
	BalanceInCents int
}

func NewBankBalance(showID string, balanceInCents int) *BankBalance {
	return &BankBalance{
		ShowID:         showID,
		BalanceInCents: balanceInCents,
	}
}
