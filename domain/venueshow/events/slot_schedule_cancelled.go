package events

type SlotScheduleCancelled struct {
	ID string
}

func NewSlotScheduleCancelled(id string) SlotScheduleCancelled {
	return SlotScheduleCancelled{
		ID: id,
	}
}
