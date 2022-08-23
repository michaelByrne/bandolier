package mongodb

import (
	"bandolier/domain/readmodel"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ShowDetailRepository struct {
	readmodel.ShowDetailRepository

	slotCollection    *mongo.Collection
	bookingCollection *mongo.Collection
	showCollection    *mongo.Collection
}

func NewShowDetailRepository(db *mongo.Database) *ShowDetailRepository {
	return &ShowDetailRepository{
		showCollection:    db.Collection("shows"),
		slotCollection:    db.Collection("slots"),
		bookingCollection: db.Collection("bookings"),
	}
}

func (m *ShowDetailRepository) AddShow(s readmodel.ShowDetail) error {
	_, err := m.showCollection.InsertOne(context.TODO(), s)
	return err
}

func (m *ShowDetailRepository) AddSlot(s readmodel.Slot) error {
	_, err := m.slotCollection.InsertOne(context.TODO(), s)
	return err
}

func (m *ShowDetailRepository) AddBooking(s readmodel.Booking) error {
	_, err := m.bookingCollection.InsertOne(context.TODO(), s)
	return err
}

func (m *ShowDetailRepository) GetShowIDForSlot(slotID string) (string, error) {
	var slot Slot
	result := m.slotCollection.FindOne(context.TODO(), bson.M{"id": slotID})
	if result == nil {
		return "", fmt.Errorf("slot not found")
	}

	err := result.Decode(&slot)
	if err != nil {
		return "", err
	}

	return slot.ShowID, nil
}

func (m *ShowDetailRepository) GetShowDetail(showID string) (readmodel.ShowDetail, error) {
	show := readmodel.ShowDetail{}
	result := m.showCollection.FindOne(context.TODO(), bson.M{"showid": showID})
	if result == nil {
		return show, fmt.Errorf("show not found")
	}

	err := result.Decode(&show)
	if err != nil {
		return show, err
	}

	slots, err := m.getSlots(showID)
	if err != nil {
		return show, err
	}

	for dex, slot := range slots {
		booking, err := m.getBooking(slot.ID)
		if err != nil {
			continue
		}

		if booking != nil {
			slots[dex].Booking = booking

			if booking.Headliner {
				show.Headliner = booking.ArtistName
			}
		}
	}

	show.Slots = slots

	return show, nil
}

func (m *ShowDetailRepository) getSlots(id string) ([]readmodel.Slot, error) {
	var slots []Slot
	cursor, err := m.slotCollection.Find(context.TODO(), bson.M{"showid": id})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &slots)
	if err != nil {
		return nil, err
	}

	slotsOut := make([]readmodel.Slot, len(slots))
	for dex, slot := range slots {
		slotsOut[dex] = readmodel.Slot{
			ID:       slot.SlotID,
			Start:    slot.StartTime,
			Duration: slot.Duration,
			ShowID:   slot.ShowID,
		}
	}

	return slotsOut, nil
}

func (m *ShowDetailRepository) getBooking(id string) (*readmodel.Booking, error) {
	var booking Booking
	result := m.bookingCollection.FindOne(context.TODO(), bson.M{"slotid": id})
	if result == nil {
		return nil, nil
	}

	err := result.Decode(&booking)
	if err != nil {
		return nil, err
	}

	return &readmodel.Booking{
		SlotID:     booking.SlotID,
		ArtistID:   booking.ArtistID,
		ArtistName: booking.ArtistName,
		Headliner:  booking.Headliner,
	}, nil
}
