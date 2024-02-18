package usererrors

import "fmt"

type Error struct {
	code    string
	message string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func New(code string, message ...string) error {
	err := &Error{
		code: code,
	}
	if len(message) > 0 {
		err.message = message[0]
	}
	return err
}
