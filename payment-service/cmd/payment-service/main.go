package main

import (
	"log"
	"paymentservice/internal/config"
	"paymentservice/internal/handlers"
	"paymentservice/internal/messaging"
	"paymentservice/internal/middleware"
	"paymentservice/internal/repository"
	"paymentservice/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "paymentservice/docs"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize MySQL database connection
	db, err := repository.NewMySQLDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Initialize RabbitMQ connection (for publishing)
	mq, err := messaging.NewRabbitMQClient(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer mq.Close()

	// (Optional) still consume reservation notifications
	go messaging.StartReservationConsumer(cfg.RabbitMQURL, "reservation_notifications")

	// Create repository and service layer instances,
	// now passing in the reservation‐service base URL from config:
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(
		paymentRepo,
		mq,
		cfg.RabbitMQEventQueue,
		cfg.ReservationServiceURL,
	)

	// Set up Gin engine with logging middleware
	r := gin.Default()
	r.Use(middleware.Logger())

	// Register payment routes
	handlers.RegisterPaymentRoutes(r, paymentService)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Starting server on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
