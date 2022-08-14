package venueshow

type SlotStatus int

const (
	SlotDoesNotExist SlotStatus = iota
	SlotAvailable
	SlotBooked
)
