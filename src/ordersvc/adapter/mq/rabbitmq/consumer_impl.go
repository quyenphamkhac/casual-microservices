package rabbitmq

import (
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/internal/logger"
	"github.com/streadway/amqp"
)

type consumerImpl struct {
	cfg    *config.Config
	conn   *amqp.Connection
	logger logger.Logger
}

type ConsumerOptions struct {
}

func NewConsumer(cfg *config.Config, conn *amqp.Connection, logger logger.Logger) *consumerImpl {
	return &consumerImpl{
		cfg:    cfg,
		conn:   conn,
		logger: logger,
	}
}

func (p *consumerImpl) Connect() error {
	return nil
}

func (p *consumerImpl) Close() error {
	return nil
}

func (p *consumerImpl) Consume(pattern interface{}, options interface{}) (interface{}, error) {
	return nil, nil
}
