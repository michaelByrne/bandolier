package inmemory

type ShowNotFoundError struct {
	ShowID string
}

func (e ShowNotFoundError) Error() string {
	return "Show not found: " + e.ShowID
}

type SlotNotFoundError struct {
	SlotID string
	ShowID string
}

func (e SlotNotFoundError) Error() string {
	return "Slot not found: " + e.SlotID + " in show: " + e.ShowID
}

type HeadlinerAlreadySetError struct {
	ShowID string
}

func (e HeadlinerAlreadySetError) Error() string {
	return "Headliner already set for show: " + e.ShowID
}
