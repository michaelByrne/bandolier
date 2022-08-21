package mongodb

import "time"

type Show struct {
	ShowID    string
	Headliner string
	Date      time.Time
	Venue     string
}

type Slot struct {
	SlotID    string    `bson:"id"`
	StartTime time.Time `bson:"start"`
	Duration  time.Duration
	ShowID    string `bson:"showid"`
}

type Booking struct {
	SlotID     string
	ArtistName string
	Headliner  bool
	ArtistID   string
}
