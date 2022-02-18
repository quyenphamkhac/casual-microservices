package rabbitmq

import (
	"fmt"
	"log"
	"sync"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/streadway/amqp"
)

var initRMQConnOnce sync.Once

func NewRabbitMQConn(cfg *config.RabbitMQ) (*amqp.Connection, error) {
	var conn *amqp.Connection
	var err error
	initRMQConnOnce.Do(func() {
		connStr := fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			cfg.User,
			cfg.Password,
			cfg.Host,
			cfg.Port,
		)
		conn, err = amqp.Dial(connStr)
	})
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
		return nil, err
	}
	return conn, err
}
