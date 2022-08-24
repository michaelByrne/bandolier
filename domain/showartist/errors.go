package showartist

type ArtistAlreadyCreatedError struct{}

func (e ArtistAlreadyCreatedError) Error() string {
	return "Artist already created"
}

type ArtistNotCreatedError struct{}

func (e ArtistNotCreatedError) Error() string {
	return "Artist not created"
}

type ArtistAlreadyPaidForShowError struct{}

func (e ArtistAlreadyPaidForShowError) Error() string {
	return "Artist already paid for show"
}
