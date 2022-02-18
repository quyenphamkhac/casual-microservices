package rabbitmq

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/internal/logger"
	"github.com/streadway/amqp"
)

type publisherImpl struct {
	cfg    *config.Config
	conn   *amqp.Connection
	logger logger.Logger
}

type PublisherOptions struct {
}

func NewPublisher(cfg *config.Config, conn *amqp.Connection, logger logger.Logger) *publisherImpl {
	return &publisherImpl{
		cfg:    cfg,
		conn:   conn,
		logger: logger,
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
