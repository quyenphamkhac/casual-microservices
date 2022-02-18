package rabbitmq

import (
	"errors"
	"fmt"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/streadway/amqp"
)

type rabbitMQ struct {
	conn *amqp.Connection
	cfg  *config.RabbitMQ
}

func NewRabbitMQConn(cfg *config.RabbitMQ) *rabbitMQ {
	return &rabbitMQ{
		cfg: cfg,
	}
}

func (r *rabbitMQ) Conn() *amqp.Connection {
	return r.conn
}

func (r *rabbitMQ) Dial() error {
	if r.cfg == nil {
		return errors.New("rabbitmq config is nil")
	}
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		r.cfg.User,
		r.cfg.Password,
		r.cfg.Host,
		r.cfg.Port,
	)
	var err error
	r.conn, err = amqp.Dial(connStr)
	if err != nil {
		return err
	}
	return nil
}
