// payment-service/tests/payment_service_test.go
package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"paymentservice/internal/models"
	"paymentservice/internal/services"

	"github.com/stretchr/testify/assert"
)

// fakeRepo implements repository.PaymentRepository
type fakeRepo struct{}

func (f *fakeRepo) Create(p *models.Payment) error {
	p.ID = 77
	p.CreatedAt = time.Now()
	return nil
}
func (f *fakeRepo) GetByID(id int64) (*models.Payment, error)            { return nil, nil }
func (f *fakeRepo) GetByUserID(userID int64) ([]*models.Payment, error) { return nil, nil }
func (f *fakeRepo) GetByBarberID(barberID int64) ([]*models.Payment, error) {
	return nil, nil
}
func (f *fakeRepo) Delete(id int64) error { return nil }

// fakeMQ captures publishes
type fakeMQ struct {
	lastRoutingKey string
	lastMessage    interface{}
}

func (f *fakeMQ) Publish(routingKey string, message interface{}) error {
	f.lastRoutingKey = routingKey
	f.lastMessage = message
	return nil
}
func (f *fakeMQ) Close() error { return nil }

func TestProcessPayment_PublishesEvent(t *testing.T) {
	// 1) HTTP test server that returns an "accepted" reservation
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]interface{}{
			"id":               "res-123",
			"user_id":          "1",
			"barber_id":        "2",
			"appointment_time": "2025-05-18T10:00:00Z",
			"status":           "accepted",
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer ts.Close()

	// 2) Inject fakes & point the service at ts.URL
	repo := &fakeRepo{}
	mq := &fakeMQ{}
	svc := services.NewPaymentService(repo, mq, "payment_events", ts.URL)

	// 3) Call ProcessPayment
	p := &models.Payment{
		UserID:        1,
		ReservationID: "res-123",
		Amount:        50,
		PaymentMethod: "card",
	}
	err := svc.ProcessPayment(p)
	assert.NoError(t, err)
	assert.Equal(t, int64(77), p.ID)

	// 4) Inspect the event map
	evt, ok := mq.lastMessage.(map[string]interface{})
	assert.True(t, ok, "event must be a map[string]interface{}")
	assert.Equal(t, "payment_events", mq.lastRoutingKey)
	assert.Equal(t, "PaymentProcessed", evt["type"])

	// 5) The "data" field is your Payment struct
	published, ok := evt["data"].(models.Payment)
	assert.True(t, ok, `"data" should be a models.Payment`)
	assert.Equal(t, p.ID, published.ID)
}

func TestDeletePayment_PublishesEvent(t *testing.T) {
	repo := &fakeRepo{}
	mq := &fakeMQ{}
	svc := services.NewPaymentService(repo, mq, "payment_events", "")

	err := svc.DeletePayment(42)
	assert.NoError(t, err)

	evt, ok := mq.lastMessage.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "PaymentDeleted", evt["type"])

	dataMap, ok := evt["data"].(map[string]interface{})
	assert.True(t, ok)
	// since we built it with an int64, it stays int64
	assert.Equal(t, int64(42), dataMap["id"])
}
