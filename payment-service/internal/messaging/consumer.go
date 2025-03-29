package messaging

import (
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// ReservationNotification represents the payload for a reservation update.
type ReservationNotification struct {
	ReservationID string `json:"reservation_id"`
	Status        string `json:"status"`
}

// StartReservationConsumer connects to RabbitMQ and listens on a dedicated queue.
func StartReservationConsumer(rabbitMQURL, queueName string) {
	for {
		conn, err := amqp.Dial(rabbitMQURL)
		if err != nil {
			log.Printf("Failed to connect to RabbitMQ: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Failed to open a channel: %v", err)
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// Declare the queue
		q, err := ch.QueueDeclare(
			queueName, // name of the queue
			true,      // durable
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Printf("Failed to declare queue: %v", err)
			ch.Close()
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		msgs, err := ch.Consume(
			q.Name,
			"",
			false, // manual ack
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Printf("Failed to register consumer: %v", err)
			ch.Close()
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("Reservation consumer started on queue: %s", q.Name)
		// Process messages in a goroutine.
		forever := make(chan bool)
		go func() {
			for d := range msgs {
				log.Printf("🔔 Raw message body: %s", d.Body) // 👈 ADD THIS
			
				var notification ReservationNotification
				if err := json.Unmarshal(d.Body, &notification); err != nil {
					log.Printf("❌ Error decoding message: %v", err)
					d.Nack(false, true)
					continue
				}
				log.Printf("✅ Received reservation notification: %+v", notification)
				d.Ack(false)
			}
		}()
		<-forever // Block until something goes wrong
		// If the consumer loop ends, close resources and try reconnecting.
		ch.Close()
		conn.Close()
		log.Println("Consumer disconnected; reconnecting in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
}
