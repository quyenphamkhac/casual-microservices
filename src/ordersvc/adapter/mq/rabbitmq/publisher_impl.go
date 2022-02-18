package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

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

func (p *publisherImpl) Publish(data interface{}) (interface{}, error) {
	fmt.Println("Hello")
	eventDataBytes, _ := json.Marshal(data)

	q, err := p.ch.QueueDeclare("producs_queue1", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(q.Name)
	err = p.ch.Publish("", q.Name, false, true, amqp.Publishing{
		ContentType: "application/json",
		Body:        eventDataBytes,
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
