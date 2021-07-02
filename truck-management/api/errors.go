package api

type ClientErrors struct {
	ErrorMessage string
	Code         int
}

func NewError(errorMessage string, Code int) *ClientErrors {
	return &ClientErrors{
		ErrorMessage: errorMessage,
		Code:         Code,
	}
}

func NewBadRequest(errorMessage string) *ClientErrors {
	return NewError(errorMessage, 400)
}

func (d *ClientErrors) Error() string {
	return d.ErrorMessage
}
