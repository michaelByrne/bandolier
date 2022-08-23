package showartist

import "time"

type ArtistID struct {
	Value string
}

func NewArtistID(name string, date time.Time) ArtistID {
	return ArtistID{
		Value: name + "_" + date.Format("2006-01-02"),
	}
}
