package errors

type AlreadyExistError struct {
	Code    string
	Message string
}

func NewAlreadyExistError(code, message string) *AlreadyExistError {
	return &AlreadyExistError{
		Code:    code,
		Message: message,
	}
}

func (e AlreadyExistError) Error() string {
	return e.Message
}
