package job

type Job interface {
	Id() string
	Run(map[string]interface{}) error
}
