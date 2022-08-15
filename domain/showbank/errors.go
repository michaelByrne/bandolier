package showbank

type DoorAlreadyPaidError struct{}

func (e DoorAlreadyPaidError) Error() string {
	return "Door already paid"
}

type BankAlreadyOpenedError struct{}

func (e BankAlreadyOpenedError) Error() string {
	return "Bank already opened for show"
}

type BankNotOpenedError struct{}

func (e BankNotOpenedError) Error() string {
	return "Bank has not been opened yet"
}
