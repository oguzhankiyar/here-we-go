package interfaces

type Producer interface {
	Open() error
	Close() error
	Produce(data []byte) error
}
