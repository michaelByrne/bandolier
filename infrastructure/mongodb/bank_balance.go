package mongodb

type BankBalance struct {
	ShowID         string `bson:"id"`
	BalanceInCents int    `bson:"balanceInCents"`
}
