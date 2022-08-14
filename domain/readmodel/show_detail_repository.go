package readmodel

import "time"

type ShowDetailRepository interface {
	AddShow(ShowDetail)
	AddSlotToShow(id, showID string, start time.Time, duration time.Duration) error
	AddSlotBookingToShow(showID string, slot Booking) error
	GetShowDetail(showID string) (*ShowDetail, error)
	SetHeadliner(showID string, headliner string) error
	GetSlot(showID string, slotID string) (*Slot, error)
	Clear()
}
