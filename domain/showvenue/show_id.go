package showvenue

import "time"

type ShowID struct {
	Value string
}

func NewShowID(venueID string, date time.Time) ShowID {
	return ShowID{
		Value: venueID + "_" + date.Format("2006-01-02"),
	}
}

func NewShowIDFrom(showID string) ShowID {
	return ShowID{
		Value: showID,
	}
}
