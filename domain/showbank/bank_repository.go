package showbank

type BankRepository interface {
	Save(bank *Bank)
	Get(showID string) (*Bank, error)
}
