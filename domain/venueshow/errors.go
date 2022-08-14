package venueshow

type ShowScheduleAlreadyArchivedError struct{}

func (e ShowScheduleAlreadyArchivedError) Error() string {
	return "Show schedule already archived"
}

type ShowScheduleAlreadyCancelledError struct{}

func (e ShowScheduleAlreadyCancelledError) Error() string {
	return "Show schedule already cancelled"
}

type ShowAlreadyScheduledError struct{}

func (e ShowAlreadyScheduledError) Error() string {
	return "Show already scheduled"
}

type ShowNotScheduledError struct{}

func (e ShowNotScheduledError) Error() string {
	return "Show not scheduled"
}

type SlotOverlapsError struct{}

func (e SlotOverlapsError) Error() string {
	return "Slot overlaps with existing slot"
}

type SlotAlreadyBookedError struct{}

func (e SlotAlreadyBookedError) Error() string {
	return "Slot already booked"
}

type SlotDoesNotExistError struct{}

func (e SlotDoesNotExistError) Error() string {
	return "Slot does not exist"
}

type InvalidSlotStatusError struct{}

func (e InvalidSlotStatusError) Error() string {
	return "Invalid slot status"
}

type SlotNotBookedError struct{}

func (e SlotNotBookedError) Error() string {
	return "Slot not booked"
}

type ShowAlreadyArchivedError struct{}

func (e ShowAlreadyArchivedError) Error() string {
	return "Show already archived"
}
