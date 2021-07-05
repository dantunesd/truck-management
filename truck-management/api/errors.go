package api

import "net/http"

type ClientErrors struct {
	ErrorMessage string
	Code         int
}

func NewError(errorMessage string, code int) *ClientErrors {
	return &ClientErrors{
		ErrorMessage: errorMessage,
		Code:         code,
	}
}

func NewBadRequest(errorMessage string) *ClientErrors {
	return NewError(errorMessage, http.StatusBadRequest)
}

func NewConflict(errorMessage string) *ClientErrors {
	return NewError(errorMessage, http.StatusConflict)
}

func NewNotFound(errorMessage string) *ClientErrors {
	return NewError(errorMessage, http.StatusNotFound)
}

func (d *ClientErrors) Error() string {
	return d.ErrorMessage
}
