package config

import "os"

type Config struct {
	ServerPort             string
	DatabaseDSN            string
	RabbitMQURL            string
	RabbitMQEventQueue     string
	ReservationServiceURL  string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	cfg := &Config{
		ServerPort:            getEnv("SERVER_PORT", "8080"),
		DatabaseDSN:           getEnv("DATABASE_DSN", "root:root@tcp(localhost:3306)/paymentdb?parseTime=true"),
		RabbitMQURL:           getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		RabbitMQEventQueue:    getEnv("RABBITMQ_EVENT_QUEUE", "payment_events"),
		ReservationServiceURL: getEnv("RESERVATION_SERVICE_URL", "http://reservation-service:8000"),
	}
	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
