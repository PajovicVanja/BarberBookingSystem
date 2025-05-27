package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"strings"

	"paymentservice/internal/config"
	"paymentservice/internal/handlers"
	"paymentservice/internal/messaging"
	"paymentservice/internal/middleware"
	"paymentservice/internal/repository"
	"paymentservice/internal/services"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed creatingdb.sql
var initSQL []byte

func applyInitScript(rawDSN string) error {
	if !strings.Contains(rawDSN, "multiStatements=true") {
		sep := "?"
		if strings.Contains(rawDSN, "?") {
			sep = "&"
		}
		rawDSN += sep + "multiStatements=true"
	}
	db, err := sql.Open("mysql", rawDSN)
	if err != nil {
		return fmt.Errorf("opening admin DB: %w", err)
	}
	defer db.Close()

	if _, err := db.Exec(string(initSQL)); err != nil {
		return fmt.Errorf("executing init SQL: %w", err)
	}
	return nil
}

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Apply DB init
	if err := applyInitScript(cfg.DatabaseDSN); err != nil {
		log.Fatalf("failed to apply DB init script: %v", err)
	}

	// Open paymentdb
	db, err := repository.NewMySQLDB(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Create RabbitMQ client *and declare* the payment_events queue
	mq, err := messaging.NewRabbitMQClient(cfg.RabbitMQURL, cfg.RabbitMQEventQueue)
	if err != nil {
		log.Fatalf("could not connect to RabbitMQ: %v", err)
	}
	defer mq.Close()

	// Start reservation-notifications consumer
	go messaging.StartReservationConsumer(cfg.RabbitMQURL, "reservation_notifications")

	// Wire up service
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := services.NewPaymentService(
		paymentRepo,
		mq,
		cfg.RabbitMQEventQueue,
		cfg.ReservationServiceURL,
	)

	// Setup HTTP server
	r := gin.Default()
	r.Use(middleware.Logger())
	handlers.RegisterPaymentRoutes(r, paymentService)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Starting server on port %s", cfg.ServerPort)
	r.Run(":" + cfg.ServerPort)
}
