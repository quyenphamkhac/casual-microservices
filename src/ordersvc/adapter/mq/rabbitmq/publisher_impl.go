package rabbitmq

import (
	"errors"
	"log"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type publisherImpl struct {
	cfg  *config.RabbitMQ
	ch   *amqp.Channel
	conn *amqp.Connection
}

func NewPublisher(cfg *config.RabbitMQ) (*publisherImpl, error) {
	conn, ch, err := rabbitmq.NewRabbitMQSession(cfg)
	if err != nil {
		log.Fatalf("Failed to create new publisher: %s", err)
		return nil, err
	}
	return &publisherImpl{
		cfg:  cfg,
		ch:   ch,
		conn: conn,
	}, nil
}

func (p *publisherImpl) Publish(data interface{}, pattern interface{}, options interface{}) error {
	body, ok := data.([]byte)
	if !ok {
		return errors.New("message data is wrong")
	}
	routingKeys, ok := pattern.([]string)
	if !ok {
		return errors.New("routing key must be an array")
	}
	opts, ok := options.(rabbitmq.PublishingOptions)
	if !ok {
		return errors.New("publishing options is wrong")
	}
	for _, routingKey := range routingKeys {
		err := p.ch.Publish(opts.Exchange, routingKey, opts.Mandatory, opts.Immediate, amqp.Publishing{
			ContentType:     opts.ContentType,
			ContentEncoding: opts.ContentEncoding,
			DeliveryMode:    opts.DeliveryMode,
			Priority:        opts.Priority,
			CorrelationId:   opts.CorrelationId,
			ReplyTo:         opts.ReplyTo,
			Expiration:      opts.Expiration,
			MessageId:       opts.MessageId,
			Timestamp:       opts.Timestamp,
			Type:            opts.Type,
			UserId:          opts.UserId,
			AppId:           opts.AppId,
			Body:            body,
		})

		if err != nil {
			log.Fatalf("Failed to publish message: %s", err)
			return err
		}
	}
	return nil
}
