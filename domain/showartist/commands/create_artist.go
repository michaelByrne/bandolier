package commands

import "time"

type CreateArtist struct {
	Name string
	Date time.Time
}

func NewCreateArtist(name string, date time.Time) CreateArtist {
	return CreateArtist{
		Name: name,
		Date: date,
	}
}
