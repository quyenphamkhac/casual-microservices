package mq

type Publisher interface {
	Connect() error
	Close() error
	Publish(data interface{}, options interface{}) (interface{}, error)
}
