package services

import (
	"errors"
	"paymentservice/internal/messaging"
	"paymentservice/internal/models"
	"paymentservice/internal/repository"
)

type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
	GetPaymentByID(id int64) (*models.Payment, error)
	GetPaymentsByUser(userID int64) ([]*models.Payment, error)
	HandleWebhook(data map[string]interface{}) error
}

type paymentService struct {
	repo repository.PaymentRepository
	mq   messaging.MessageQueue
}

func NewPaymentService(repo repository.PaymentRepository, mq messaging.MessageQueue) PaymentService {
	return &paymentService{
		repo: repo,
		mq:   mq,
	}
}

// ProcessPayment validates the reservation asynchronously,
// processes the payment, saves it to the DB, and publishes an event.
func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	// Simulate asynchronous validation with the Reservation Service.
	valid := make(chan bool)
	go func() {
		// In a real system, you would perform an HTTP call or similar.
		valid <- true
	}()
	if !<-valid {
		return errors.New("invalid reservation")
	}

	// Process the payment (simulate success).
	payment.Status = "success"

	// Save the payment record in the database.
	if err := s.repo.Create(payment); err != nil {
		return err
	}

	// Publish the processed payment event asynchronously.
	go func(p models.Payment) {
		if err := s.mq.Publish("payment.processed", p); err != nil {
			// In production, log the error appropriately.
		}
	}(*payment)

	return nil
}

func (s *paymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *paymentService) GetPaymentsByUser(userID int64) ([]*models.Payment, error) {
	return s.repo.GetByUserID(userID)
}

func (s *paymentService) HandleWebhook(data map[string]interface{}) error {
	// Process webhook data.
	// For example, update payment status based on webhook event.
	return nil
}
