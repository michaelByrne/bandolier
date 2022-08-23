package readmodel

type BankBalanceRepository interface {
	GetBalance(showID string) (*BankBalance, error)
	UpdateBalance(showID string, balanceInCents int) error
	OpenBank(showID string, initialBalanceInCents int) error
}
