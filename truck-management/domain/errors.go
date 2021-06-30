package domain

type DomainErrors struct {
	ErrorMessage string
	Code         int
}

func NewError(errorMessage string, Code int) *DomainErrors {
	return &DomainErrors{
		ErrorMessage: errorMessage,
		Code:         Code,
	}
}

func NewConflict(errorMessage string) *DomainErrors {
	return NewError(errorMessage, 409)
}

func (d *DomainErrors) Error() string {
	return d.ErrorMessage
}
