package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"paymentservice/internal/handlers"
	"paymentservice/internal/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// mockPaymentService implements the PaymentService interface for testing.
type mockPaymentService struct{}

func (m *mockPaymentService) ProcessPayment(payment *models.Payment) error {
	payment.ID = 1
	return nil
}

func (m *mockPaymentService) GetPaymentByID(id int64) (*models.Payment, error) {
	return &models.Payment{
		ID:            id,
		UserID:        1,
		ReservationID: 1,
		Amount:        100.0,
		Status:        "success",
		PaymentMethod: "credit_card",
	}, nil
}

func (m *mockPaymentService) GetPaymentsByUser(userID int64) ([]*models.Payment, error) {
	return []*models.Payment{
		{
			ID:            1,
			UserID:        userID,
			ReservationID: 1,
			Amount:        100.0,
			Status:        "success",
			PaymentMethod: "credit_card",
		},
	}, nil
}

func (m *mockPaymentService) HandleWebhook(data map[string]interface{}) error {
	return nil
}

func TestProcessPayment(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	service := &mockPaymentService{}
	handlers.RegisterPaymentRoutes(r, service)

	payment := models.Payment{
		UserID:        1,
		ReservationID: 1,
		Amount:        100.0,
		PaymentMethod: "credit_card",
	}
	jsonValue, _ := json.Marshal(payment)
	req, _ := http.NewRequest("POST", "/api/payments", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	var respPayment models.Payment
	json.Unmarshal(resp.Body.Bytes(), &respPayment)
	assert.Equal(t, int64(1), respPayment.ID)
}

func TestGetPaymentByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	service := &mockPaymentService{}
	handlers.RegisterPaymentRoutes(r, service)

	req, _ := http.NewRequest("GET", "/api/payments/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	var payment models.Payment
	json.Unmarshal(resp.Body.Bytes(), &payment)
	assert.Equal(t, int64(1), payment.ID)
}

func TestGetPaymentsByUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	service := &mockPaymentService{}
	handlers.RegisterPaymentRoutes(r, service)

	req, _ := http.NewRequest("GET", "/api/payments/user/1", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
	var payments []models.Payment
	json.Unmarshal(resp.Body.Bytes(), &payments)
	assert.Equal(t, 1, len(payments))
}

func TestHandleWebhook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	service := &mockPaymentService{}
	handlers.RegisterPaymentRoutes(r, service)

	payload := map[string]interface{}{
		"event": "payment.updated",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/api/payments/webhook", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, 200, resp.Code)
}
