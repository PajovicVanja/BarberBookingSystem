package messaging

import (
	"encoding/json"
	"log"

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

// NewRabbitMQClient creates a new RabbitMQ client and declares the event queue.
func NewRabbitMQClient(url, eventQueue string) (MessageQueue, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	// Ensure the event queue exists
	_, err = ch.QueueDeclare(
		eventQueue, // queue name
		true,       // durable
		false,      // auto-delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &rabbitMQClient{conn: conn, channel: ch}, nil
}

// Publish sends a message to the specified routingKey (queue).
func (r *rabbitMQClient) Publish(routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	if err := r.channel.Publish(
		"",         // default exchange
		routingKey, // routing key == queue name
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		log.Printf("🔴 failed to publish to queue %q: %v", routingKey, err)
	}
	return err
}

// Close cleans up the RabbitMQ connection.
func (r *rabbitMQClient) Close() error {
	if err := r.channel.Close(); err != nil {
		r.conn.Close()
		return err
	}
	return r.conn.Close()
}
