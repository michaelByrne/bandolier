package showvenue

type Venue struct {
	ID   string
	Name string
}

func NewVenue(id string, name string) Venue {
	return Venue{
		ID:   id,
		Name: name,
	}
}
