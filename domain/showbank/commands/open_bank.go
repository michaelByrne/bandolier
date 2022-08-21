package commands

type OpenBank struct {
	ShowID         string
	PresaleInCents int
}

func NewOpenBank(showID string, presaleInCents int) OpenBank {
	return OpenBank{
		ShowID:         showID,
		PresaleInCents: presaleInCents,
	}
}
