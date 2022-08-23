package showartist

type ArtistAlreadyCreatedError struct{}

func (e ArtistAlreadyCreatedError) Error() string {
	return "Artist already created"
}
