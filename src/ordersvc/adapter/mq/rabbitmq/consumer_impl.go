package rabbitmq

import (
	"log"
	"sync"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type consumerImpl struct {
	cfg  *config.RabbitMQ
	conn *amqp.Connection
	ch   *amqp.Channel
	mu   *sync.RWMutex
}

type Action int

type Delivery struct {
	amqp.Delivery
}

type Handler func(d Delivery) (action Action)

const (
	Ack Action = iota
	NackDiscard
	NackRequeue
)

type ConsumerOptions struct {
	// Queue Declare Options
	QueueDeclare    bool
	QueueName       string
	QueueDurable    bool
	QueueAutoDelete bool
	QueueExclusive  bool
	QueueNoWait     bool
	QueueArgs       rabbitmq.Table
	// Consume Options
	ConsumerName string
	AutoAck      bool
	Exclusive    bool
	NoWait       bool
	NoLocal      bool
	Args         rabbitmq.Table
}

func NewConsumer(cfg *config.RabbitMQ) (*consumerImpl, error) {
	conn, ch, err := rabbitmq.NewRabbitMQSession(cfg)
	if err != nil {
		log.Fatalf("Failed to create new publisher: %s", err)
		return nil, err
	}
	return &consumerImpl{
		cfg:  cfg,
		ch:   ch,
		conn: conn,
		mu:   &sync.RWMutex{},
	}, err
}

func (c *consumerImpl) StartConsuming(handler Handler, options ConsumerOptions) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if options.QueueDeclare {
		_, err := c.ch.QueueDeclare(
			options.QueueName,
			options.QueueDurable,
			options.QueueAutoDelete,
			options.QueueExclusive,
			options.QueueNoWait,
			amqp.Table(options.QueueArgs),
		)
		if err != nil {
			log.Fatalf("Unable to declare queue: %v", err)
		}
	}

	msgs, err := c.ch.Consume(
		options.QueueName,
		options.ConsumerName,
		options.AutoAck,
		options.Exclusive,
		options.NoLocal,
		options.NoWait,
		amqp.Table(options.Args),
	)
	if err != nil {
		log.Fatalf("Unable to start consumer: %s", err)
	}

	go func(d <-chan amqp.Delivery, opts ConsumerOptions) {
		for msg := range msgs {
			if opts.AutoAck {
				handler(Delivery{msg})
				continue
			}
			switch handler(Delivery{msg}) {
			case Ack:
				err := msg.Ack(false)
				if err != nil {
					log.Fatalf("Can't ack message: %v", err)
				}
			case NackDiscard:
				err := msg.Nack(false, false)
				if err != nil {
					log.Fatalf("Can't nack message: %v", err)
				}
			case NackRequeue:
				err := msg.Nack(false, true)
				if err != nil {
					log.Fatalf("Can't nack message: %v", err)
				}
			}
			handler(Delivery{msg})
		}
	}(msgs, options)

	return nil
}
