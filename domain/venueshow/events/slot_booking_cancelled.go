package events

type SlotBookingCancelled struct {
	ID string
}

func NewSlotBookingCancelled(id string) SlotBookingCancelled {
	return SlotBookingCancelled{
		ID: id,
	}
}
