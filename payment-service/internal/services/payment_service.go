package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"paymentservice/internal/messaging"
	"paymentservice/internal/models"
	"paymentservice/internal/repository"
	"paymentservice/internal/utils"

	gobreaker "github.com/sony/gobreaker/v2"
)

// PaymentService defines our service interface.
type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
	GetPaymentByID(id int64) (*models.Payment, error)
	GetPaymentsByUser(userID int64) ([]*models.Payment, error)
	GetPaymentsByBarber(barberID int64) ([]*models.Payment, error)
	HandleWebhook(data map[string]interface{}) error
	DeletePayment(id int64) error
}

type paymentService struct {
	repo                  repository.PaymentRepository
	mq                    messaging.MessageQueue
	eventQueue            string
	reservationServiceURL string
}

// NewPaymentService takes repo, mq, the domain‐event queue name, and the reservation‐service base URL.
// We RESET the circuit breaker here so that any “open” state from a previous run is cleared.
func NewPaymentService(
	repo repository.PaymentRepository,
	mq messaging.MessageQueue,
	eventQueue string,
	reservationServiceURL string,
) PaymentService {
   if eventQueue != "" {
       utils.ResetReservationCB()
   }

	return &paymentService{
		repo:                  repo,
		mq:                    mq,
		eventQueue:            eventQueue,
		reservationServiceURL: reservationServiceURL,
	}
}

// ProcessPayment calls the reservation‐service via a circuit‐breaker, saves the payment,
// and publishes a PaymentProcessed event.
func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	url := fmt.Sprintf("%s/api/reservations/%s", s.reservationServiceURL, payment.ReservationID)

	resIfc, err := utils.ReservationCB.Execute(func() (*http.Response, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("reservation service returned %d", resp.StatusCode)
		}
		return resp, nil
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return utils.ErrCircuitOpen
		}
		return fmt.Errorf("could not fetch reservation: %w", err)
	}
	resp := resIfc
	defer resp.Body.Close()

	var res struct {
		ID       string `json:"id"`
		UserID   string `json:"user_id"`
		BarberID string `json:"barber_id"`
		Status   string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return errors.New("error decoding reservation")
	}
	if res.Status != "accepted" {
		return errors.New("reservation not accepted")
	}

	barberID, err := strconv.ParseInt(res.BarberID, 10, 64)
	if err != nil {
		return errors.New("invalid barber ID in reservation")
	}
	payment.BarberID = barberID

	payment.Status = "success"
	if err := s.repo.Create(payment); err != nil {
		return err
	}

	event := map[string]interface{}{
		"type": "PaymentProcessed",
		"data": *payment,
	}
	if err := s.mq.Publish(s.eventQueue, event); err != nil {
		log.Printf("🔴 failed to publish PaymentProcessed event: %v", err)
	}

	return nil
}

func (s *paymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	return s.repo.GetByID(id)
}

func (s *paymentService) GetPaymentsByUser(userID int64) ([]*models.Payment, error) {
	return s.repo.GetByUserID(userID)
}

func (s *paymentService) GetPaymentsByBarber(barberID int64) ([]*models.Payment, error) {
	return s.repo.GetByBarberID(barberID)
}

func (s *paymentService) HandleWebhook(data map[string]interface{}) error {
	// no-op or your custom logic
	return nil
}

func (s *paymentService) DeletePayment(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	event := map[string]interface{}{
		"type": "PaymentDeleted",
		"data": map[string]interface{}{"id": id},
	}
	if err := s.mq.Publish(s.eventQueue, event); err != nil {
		log.Printf("🔴 failed to publish PaymentDeleted event: %v", err)
	}
	return nil
}
