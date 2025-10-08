package usererrors

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	httpStatus int
	code       string
	message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.message)
}

func (e *Error) MarshalJSON() ([]byte, error) {
	mapping := map[string]string{
		"code":    e.code,
		"message": e.message,
	}
	return json.Marshal(mapping)
}

func (e *Error) Code() string {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func New(code string, message ...string) *Error {
	err := &Error{
		code:       code,
		httpStatus: 200,
	}
	if len(message) > 0 {
		err.message = message[0]
	}
	return err
}

func (e *Error) SetHttpStatus(status int) *Error {
	e.httpStatus = status
	return e
}

func (e *Error) HttpStatus() int {
	if e.httpStatus == 0 {
		return 200
	}
	return e.httpStatus
}
