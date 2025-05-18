package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"paymentservice/internal/messaging"
	"paymentservice/internal/models"
	"paymentservice/internal/repository"
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

// NewPaymentService takes repo, mq, the domain-event queue name, and the reservation-service base URL.
func NewPaymentService(
	repo repository.PaymentRepository,
	mq messaging.MessageQueue,
	eventQueue string,
	reservationServiceURL string,
) PaymentService {
	return &paymentService{
		repo:                  repo,
		mq:                    mq,
		eventQueue:            eventQueue,
		reservationServiceURL: reservationServiceURL,
	}
}

func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	// 1) Fetch reservation to confirm it’s accepted
	url := fmt.Sprintf("%s/api/reservations/%s", s.reservationServiceURL, payment.ReservationID)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return errors.New("could not fetch reservation")
	}
	defer resp.Body.Close()

	var res struct {
		ID              string `json:"id"`
		UserID          string `json:"user_id"`
		BarberID        string `json:"barber_id"`
		AppointmentTime string `json:"appointment_time"`
		Status          string `json:"status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return errors.New("error decoding reservation")
	}
	if res.Status != "accepted" {
		return errors.New("reservation not accepted")
	}

	// 2) Record barber ID
	barberID, err := strconv.ParseInt(res.BarberID, 10, 64)
	if err != nil {
		return errors.New("invalid barber ID in reservation")
	}
	payment.BarberID = barberID

	// 3) Save payment
	payment.Status = "success"
	if err := s.repo.Create(payment); err != nil {
		return err
	}

	// 4) **Synchronously** publish a domain event: PaymentProcessed
	event := map[string]interface{}{
		"type": "PaymentProcessed",
		"data": *payment,
	}
	_ = s.mq.Publish(s.eventQueue, event)

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
	// implement as needed
	return nil
}

func (s *paymentService) DeletePayment(id int64) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	// **Synchronously** publish a domain event: PaymentDeleted
	event := map[string]interface{}{
		"type": "PaymentDeleted",
		"data": map[string]interface{}{"id": id},
	}
	_ = s.mq.Publish(s.eventQueue, event)

	return nil
}
