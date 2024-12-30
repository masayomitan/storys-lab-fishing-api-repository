package response

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")

	ErrInvalidInput = errors.New("invalid input")
)

type Error struct {
	Errors     []string `json:"errors"`
	statusCode int
}

func NewError(err error, status int) *Error {
	return &Error{
		Errors:     []string{err.Error()},
		statusCode: status,
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		Errors:     messages,
		statusCode: status,
	}
}

func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
