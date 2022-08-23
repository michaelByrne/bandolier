package events

type ShowArchived struct {
	ShowID            string
	DoorAmountInCents int
}

func NewShowArchived(id string, door int) ShowArchived {
	return ShowArchived{
		ShowID:            id,
		DoorAmountInCents: door,
	}
}
