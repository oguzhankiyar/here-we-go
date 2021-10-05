package interfaces

type Logger interface {
	Init() error
	Debug(msg string, args ...map[string]interface{})
	Info(msg string, args ...map[string]interface{})
	Warn(msg string, args ...map[string]interface{})
	Error(msg string, err error, args ...map[string]interface{})
	Fatal(msg string, err error, args ...map[string]interface{})
}
