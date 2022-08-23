package showvenue

type Artist struct {
	Name      string
	ID        string
	Headliner bool
}

func NewArtist(name string, id string, headliner bool) Artist {
	return Artist{
		Name:      name,
		ID:        id,
		Headliner: headliner,
	}
}
