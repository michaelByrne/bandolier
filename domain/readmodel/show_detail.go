package readmodel

import "time"

type Booking struct {
	SlotID     string
	ArtistID   string
	ArtistName string
	Start      time.Time
	Duration   time.Duration
}

type ShowDetail struct {
	ShowID    string
	Headliner string
	Date      time.Time
	Venue     string
	Slots     []Slot
}

type Slot struct {
	ID       string
	Start    time.Time
	Duration time.Duration
	Booking  *Booking
}

func NewShowDetail(showID string, headliner string, date time.Time, venue string, slots []Slot) *ShowDetail {
	return &ShowDetail{
		ShowID:    showID,
		Headliner: headliner,
		Date:      date,
		Venue:     venue,
		Slots:     slots,
	}
}

func NewBooking(id string, artistID string, artistName string, start time.Time, duration time.Duration) *Booking {
	return &Booking{
		SlotID:     id,
		ArtistID:   artistID,
		ArtistName: artistName,
		Start:      start,
		Duration:   duration,
	}
}

func NewSlot(id string, start time.Time, duration time.Duration) *Slot {
	return &Slot{
		ID:       id,
		Start:    start,
		Duration: duration,
	}
}
