package mq

type Publisher interface {
	Publish(data interface{}, pattern interface{}, options interface{}) error
}
