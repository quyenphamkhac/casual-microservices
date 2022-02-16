package rabbitmq

import (
	"fmt"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/streadway/amqp"
)

func NewRabbitMQConn(cfg *config.Config) (*amqp.Connection, error) {
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	return amqp.Dial(connStr)
}
