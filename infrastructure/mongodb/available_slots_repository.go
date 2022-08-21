package mongodb

import (
	"context"
	"fmt"
	"time"

	"bandolier/domain/readmodel"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvailableSlotsRepository struct {
	readmodel.AvailableSlotsRepository

	collection *mongo.Collection
}

func NewAvailableSlotsRepository(db *mongo.Database) *AvailableSlotsRepository {
	return &AvailableSlotsRepository{
		collection: db.Collection("available_slots_v2"),
	}
}

func (m *AvailableSlotsRepository) Add(s readmodel.AvailableSlot) error {
	mongoSlot := FromAvailableSlot(s)
	_, err := m.collection.InsertOne(context.TODO(), mongoSlot)
	return err
}

func (m *AvailableSlotsRepository) MarkAsUnavailable(slotId string) error {
	result, err := m.collection.UpdateOne(
		context.TODO(),
		bson.M{"id": slotId},
		bson.D{{"$set", bson.D{{"isbooked", true}}}})

	if result.ModifiedCount == 0 {
		return fmt.Errorf("failed to hide slot")
	}

	return err
}

func (m *AvailableSlotsRepository) MarkAsAvailable(slotId string) error {
	result, err := m.collection.UpdateOne(
		context.TODO(),
		bson.M{"id": slotId},
		bson.D{{"$set", bson.D{{"isbooked", false}}}})

	if result.ModifiedCount == 0 {
		return fmt.Errorf("failed to show slot")
	}

	return err
}

func (m *AvailableSlotsRepository) DeleteSlot(slotId uuid.UUID) error {
	result, err := m.collection.DeleteOne(
		context.TODO(),
		bson.M{"id": slotId.String()})

	if result.DeletedCount == 0 {
		return fmt.Errorf("failed to delete slot")
	}

	return err
}

func (m *AvailableSlotsRepository) GetSlotsAvailableOn(date time.Time) ([]*readmodel.AvailableSlot, error) {
	slots := make([]*readmodel.AvailableSlot, 0)
	cur, err := m.collection.Find(context.TODO(), bson.D{{"date", date.Format("2006-01-02")}, {"isbooked", false}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var result AvailableSlot
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}

		slots = append(slots, result.ToAvailableSlot())
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return slots, nil
}
