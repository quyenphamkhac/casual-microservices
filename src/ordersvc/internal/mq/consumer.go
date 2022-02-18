package mq

type Consumer interface {
	Connect() error
	Close() error
	Consume(pattern interface{}, options interface{}) (interface{}, error)
}
