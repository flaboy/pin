package usererrors

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func New(message string) error {
	return &Error{message}
}
