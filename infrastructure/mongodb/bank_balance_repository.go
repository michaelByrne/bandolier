package mongodb

import (
	"bandolier/domain/readmodel"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BankBalanceRepository struct {
	readmodel.BankBalanceRepository

	balanceCollection *mongo.Collection
	db                *mongo.Database
}

func NewBankBalanceRepository(db *mongo.Database) *BankBalanceRepository {
	return &BankBalanceRepository{
		db:                db,
		balanceCollection: db.Collection("bank_balance"),
	}
}

func (m *BankBalanceRepository) GetBalance(showID string) (*readmodel.BankBalance, error) {
	balance := BankBalance{}
	result := m.balanceCollection.FindOne(context.TODO(), bson.M{"id": showID})
	if result == nil {
		return nil, fmt.Errorf("balance not found")
	}

	err := result.Decode(&balance)
	if err != nil {
		return nil, err
	}

	return &readmodel.BankBalance{
		ShowID:         balance.ShowID,
		BalanceInCents: balance.BalanceInCents,
	}, nil
}

func (m *BankBalanceRepository) OpenBank(showID string, initialBalanceInCents int) error {
	_, err := m.balanceCollection.InsertOne(context.TODO(), BankBalance{
		ShowID:         showID,
		BalanceInCents: initialBalanceInCents,
	})
	return err
}

func (m *BankBalanceRepository) UpdateBalance(showID string, balanceInCents int) error {
	_, err := m.balanceCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": showID},
		bson.M{"$set": bson.M{"balanceInCents": balanceInCents}},
	)
	return err
}
