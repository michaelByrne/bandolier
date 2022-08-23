package events

type ArtistCreated struct {
	Name string
	ID   string
}

func NewArtistCreated(name string, id string) ArtistCreated {
	return ArtistCreated{
		Name: name,
		ID:   id,
	}
}
