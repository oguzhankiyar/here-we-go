package errors

type ConfigValidateError struct {
	message string
}

func NewConfigValidateError(message string) ConfigValidateError {
	return ConfigValidateError{
		message: "config could not be validated: " + message,
	}
}

func (e ConfigValidateError) Error() string {
	return e.message
}
