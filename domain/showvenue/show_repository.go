package showvenue

type ShowRepository interface {
	Save(show *Show)
	Get(id ShowID) (*Show, error)
}
