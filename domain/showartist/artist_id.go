package showartist

import (
	"github.com/iancoleman/strcase"
	"time"
)

type ArtistID struct {
	Value string
}

func NewArtistID(name string, date time.Time) ArtistID {
	return ArtistID{
		Value: strcase.ToSnake(name) + "_" + date.Format("2006-01-02"),
	}
}
