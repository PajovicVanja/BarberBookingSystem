package tests

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sony/gobreaker/v2"
	"paymentservice/internal/models"
	"paymentservice/internal/services"
	"paymentservice/internal/utils"
)

// dummyRepo satisfies repository.PaymentRepository but does nothing.
type dummyRepo struct{}

func (d *dummyRepo) Create(p *models.Payment) error               { p.ID = 1; return nil }
func (d *dummyRepo) GetByID(id int64) (*models.Payment, error)   { return nil, nil }
func (d *dummyRepo) GetByUserID(id int64) ([]*models.Payment, error) { return nil, nil }
func (d *dummyRepo) GetByBarberID(id int64) ([]*models.Payment, error) { return nil, nil }
func (d *dummyRepo) Delete(id int64) error                       { return nil }

// newTestBreaker constructs a circuit‐breaker with low thresholds for testing.
func newTestBreaker() *gobreaker.CircuitBreaker[*http.Response] {
	settings := gobreaker.Settings{
		Name:        "TestCB",
		MaxRequests: 1,
		Interval:    time.Second,
		Timeout:     time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.TotalFailures >= 3
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			fmt.Printf("CB state change: %s → %s\n", from.String(), to.String())
		},
	}
	return gobreaker.NewCircuitBreaker[*http.Response](settings)
}

func TestCircuitBreakerTrips(t *testing.T) {
	// 1) a handler that always fails
	failSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer failSrv.Close()

	// 2) reset the global breaker to our test instance
	utils.ReservationCB = newTestBreaker()

	// 3) construct our service, pointing at the failing server
	svc := services.NewPaymentService(&dummyRepo{}, nil, "", failSrv.URL)
	p := &models.Payment{ReservationID: "foo"}

	// 4) hammer it until the breaker opens
	var lastErr error
	for i := 0; i < 10; i++ {
		lastErr = svc.ProcessPayment(p)
		if errors.Is(lastErr, utils.ErrCircuitOpen) {
			t.Logf("breaker opened after %d attempts", i+1)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	if !errors.Is(lastErr, utils.ErrCircuitOpen) {
		t.Fatal("expected breaker to open, but it did not")
	}

	// 5) subsequent calls must fail fast
	start := time.Now()
	err := svc.ProcessPayment(p)
	elapsed := time.Since(start)

	if !errors.Is(err, utils.ErrCircuitOpen) {
		t.Fatalf("expected circuit-open error, got %v", err)
	}
	if elapsed > 5*time.Millisecond {
		t.Fatalf("expected fast-fail but took %v", elapsed)
	}
}
