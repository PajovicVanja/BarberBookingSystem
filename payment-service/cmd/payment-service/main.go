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
	"github.com/swaggo/gin-swagger" // Swagger handler
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

	// Initialize RabbitMQ connection
	mq, err := messaging.NewRabbitMQClient(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer mq.Close()

	// Create repository and service layer instances
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(paymentRepo, mq)

	// Set up Gin engine with logging middleware
	r := gin.Default()
	r.Use(middleware.Logger())

	// Register payment routes with Swagger documentation available at /swagger/index.html
	handlers.RegisterPaymentRoutes(r, paymentService)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Starting server on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
