package errors

type NotFoundError struct {
	Code    string
	Message string
}

func NewNotFoundError(code, message string) *NotFoundError {
	return &NotFoundError{
		Code:    code,
		Message: message,
	}
}

func (e NotFoundError) Error() string {
	return e.Message
}
