package readmodel

type ShowDetailRepository interface {
	AddShow(s ShowDetail) error
	AddSlot(s Slot) error
	AddBooking(s Booking) error
	GetShowDetail(showID string) (ShowDetail, error)
	GetShowIDForSlot(slotID string) (string, error)
}
