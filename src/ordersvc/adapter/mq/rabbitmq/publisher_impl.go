package rabbitmq

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/streadway/amqp"
)

type publisherImpl struct {
	cfg *config.Config
	ch  *amqp.Channel
}

type PublisherOptions struct {
}

func NewPublisher(cfg *config.Config, ch *amqp.Channel) *publisherImpl {
	return &publisherImpl{
		cfg: cfg,
		ch:  ch,
	}
}

func (p *publisherImpl) Connect() error {
	return nil
}

func (p *publisherImpl) Close() error {
	return nil
}

func (p *publisherImpl) Publish(data interface{}, options interface{}) (interface{}, error) {
	return nil, nil
}
