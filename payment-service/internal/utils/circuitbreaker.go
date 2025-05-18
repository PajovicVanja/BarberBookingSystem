package utils

import (
    "errors"
    "log"
    "net/http"
    "time"

    gobreaker "github.com/sony/gobreaker/v2"
)

// ReservationCB is the (global) circuit breaker used by PaymentService.
var ReservationCB *gobreaker.CircuitBreaker[*http.Response]

// ErrCircuitOpen is returned when the circuit is open.
var ErrCircuitOpen = errors.New("reservation service unavailable (circuit open)")

// newDefaultReservationCB creates a fresh breaker with your default settings.
func newDefaultReservationCB() *gobreaker.CircuitBreaker[*http.Response] {
    settings := gobreaker.Settings{
        Name:        "ReservationServiceCB",
        MaxRequests: 5,
        Interval:    30 * time.Second,
        Timeout:     60 * time.Second,
        ReadyToTrip: func(counts gobreaker.Counts) bool {
            if counts.Requests < 10 {
                return false
            }
            return float64(counts.TotalFailures)/float64(counts.Requests) > 0.5
        },
        OnStateChange: func(name string, from, to gobreaker.State) {
            log.Printf("circuit breaker %q: %s → %s", name, from.String(), to.String())
        },
    }
    return gobreaker.NewCircuitBreaker[*http.Response](settings)
}

// ResetReservationCB wipes out the old breaker and replaces it with a fresh one.
func ResetReservationCB() {
    ReservationCB = newDefaultReservationCB()
}

func init() {
    // initialize global breaker on package load
    ResetReservationCB()
}
