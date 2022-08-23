package events

type ShowScheduleCancelled struct {
	ShowID string
}

func NewShowScheduleCancelled(showID string) ShowScheduleCancelled {
	return ShowScheduleCancelled{
		ShowID: showID,
	}
}
