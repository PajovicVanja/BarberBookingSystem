package messaging

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

// MessageQueue is an interface for message publishing.
type MessageQueue interface {
	Publish(routingKey string, message interface{}) error
	Close() error
}

type rabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQClient creates a new RabbitMQ client.
func NewRabbitMQClient(url string) (MessageQueue, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &rabbitMQClient{
		conn:    conn,
		channel: ch,
	}, nil
}

// Publish sends a message to the specified routing key.
func (r *rabbitMQClient) Publish(routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	return r.channel.Publish(
		"", // default exchange
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Close cleans up the RabbitMQ connection.
func (r *rabbitMQClient) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	return r.conn.Close()
}
