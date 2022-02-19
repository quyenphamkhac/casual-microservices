package rabbitmq

import (
	"fmt"
	"log"
	"time"

	"github.com/quyenphamkhac/casual-microservices/src/ordersvc/config"
	"github.com/streadway/amqp"
)

type Table map[string]interface{}

type PublishingOptions struct {
	// routing
	Exchange  string
	Mandatory bool
	Immediate bool

	// properties
	ContentType     string    // MIME content type
	ContentEncoding string    // MIME content encoding
	DeliveryMode    uint8     // Transient (0 or 1) or Persistent (2)
	Priority        uint8     // 0 to 9
	CorrelationId   string    // correlation identifier
	ReplyTo         string    // address to to reply to (ex: RPC)
	Expiration      string    // message expiration spec
	MessageId       string    // message identifier
	Timestamp       time.Time // message timestamp
	Type            string    // message type name
	UserId          string    // creating user id - ex: "guest"
	AppId           string
}

func NewRabbitMQSession(cfg *config.RabbitMQ) (*amqp.Connection, *amqp.Channel, error) {
	var conn *amqp.Connection
	var ch *amqp.Channel
	var err error
	connStr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)
	conn, err = amqp.Dial(connStr)
	if err != nil {
		log.Fatalf("Failed to establish rabbitmq connection: %s", err)
		return nil, nil, err
	}
	ch, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to establish rabbitmq channel: %s", err)
		return nil, nil, err
	}
	return conn, ch, err
}
