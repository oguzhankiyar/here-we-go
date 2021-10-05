package errors

type ConfigParseError struct {
	message string
}

func NewConfigParseError(err error) ConfigParseError {
	return ConfigParseError{
		message: "config could not be parsed: " + err.Error(),
	}
}

func (e ConfigParseError) Error() string {
	return e.message
}
