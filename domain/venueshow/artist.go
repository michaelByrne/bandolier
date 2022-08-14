package venueshow

type Artist struct {
	Name string
	ID   string
}

func NewArtist(name string, id string) Artist {
	return Artist{
		Name: name,
		ID:   id,
	}
}
