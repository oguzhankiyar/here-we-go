package errors

type BadRequestError struct {
	Code    string
	Message string
}

func NewBadRequestError(code, message string) *BadRequestError {
	return &BadRequestError{
		Code:    code,
		Message: message,
	}
}

func (e BadRequestError) Error() string {
	return e.Message
}
