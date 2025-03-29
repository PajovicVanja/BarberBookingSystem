package models

import "time"

// Payment represents a payment record.
type Payment struct {
	ID            int64     `json:"id" db:"id"`
	UserID        int64     `json:"user_id" db:"user_id"`
	ReservationID int64     `json:"reservation_id" db:"reservation_id"`
	Amount        float64   `json:"amount" db:"amount"`
	Status        string    `json:"status" db:"status"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
