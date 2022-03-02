package mq

type Consumer interface {
	StartConsuming(pattern interface{}, options interface{}) (interface{}, error)
	StopConsuming(pattern interface{}) error
	Disconnect()
}
