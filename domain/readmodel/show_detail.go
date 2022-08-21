package readmodel

import "time"

type Booking struct {
	SlotID     string `json:"slotId"`
	ArtistID   string `json:"artistId"`
	ArtistName string `json:"artistName"`
	Headliner  bool   `json:"headliner"`
}

type ShowDetail struct {
	ShowID    string `json:"showId"`
	Headliner string
	Date      time.Time `json:"date"`
	Venue     string    `json:"venue"`
	Slots     []Slot
}

type Slot struct {
	ID       string        `json:"id"`
	ShowID   string        `json:"showId"`
	Start    time.Time     `json:"start"`
	Duration time.Duration `json:"duration"`
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

func NewBooking(id string, artistID string, artistName string, headliner bool) *Booking {
	return &Booking{
		SlotID:     id,
		ArtistID:   artistID,
		ArtistName: artistName,
		Headliner:  headliner,
	}
}

func NewSlot(id string, start time.Time, duration time.Duration, showID string) *Slot {
	return &Slot{
		ID:       id,
		Start:    start,
		Duration: duration,
		ShowID:   showID,
	}
}
