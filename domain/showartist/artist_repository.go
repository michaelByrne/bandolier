package showartist

type ArtistRepository interface {
	Save(bank *Artist)
	Get(artistID string) (*Artist, error)
}
