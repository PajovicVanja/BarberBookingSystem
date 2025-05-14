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

type PaymentService interface {
	ProcessPayment(payment *models.Payment) error
	GetPaymentByID(id int64) (*models.Payment, error)
	GetPaymentsByUser(userID int64) ([]*models.Payment, error)
	GetPaymentsByBarber(barberID int64) ([]*models.Payment, error)
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

func (s *paymentService) ProcessPayment(payment *models.Payment) error {
	// 🔍 1) Fetch the reservation to ensure it’s been accepted
	url := fmt.Sprintf("http://reservation-service:8000/api/reservations/%s", payment.ReservationID)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return errors.New("could not fetch reservation")
	}
	defer resp.Body.Close()

	// The reservation endpoint returns a Reservation JSON directly
	var res struct {
		ID             string `json:"id"`
		UserID         string `json:"user_id"`
		BarberID       string `json:"barber_id"`
		AppointmentTime string `json:"appointment_time"`
		Status         string `json:"status"`
		Message        string `json:"message,omitempty"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return errors.New("error decoding reservation")
	}
	if res.Status != "accepted" {
		return errors.New("reservation not accepted")
	}

	// 🔖 2) Record which barber this payment is for
	barberID, err := strconv.ParseInt(res.BarberID, 10, 64)
	if err != nil {
		return errors.New("invalid barber ID in reservation")
	}
	payment.BarberID = barberID

	// ✔ 3) Process & save the payment
	payment.Status = "success"
	if err := s.repo.Create(payment); err != nil {
		return err
	}

	// 📣 4) Publish payment.processed event
	go func(p models.Payment) {
		_ = s.mq.Publish("payment.processed", p)
	}(*payment)

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
