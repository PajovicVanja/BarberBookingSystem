package repository

import (
	"database/sql"
	"paymentservice/internal/models"
	"time"
	"log"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id int64) (*models.Payment, error)
	GetByUserID(userID int64) ([]*models.Payment, error)
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// Create inserts a new payment record into the database.
func (r *paymentRepository) Create(payment *models.Payment) error {
	query := "INSERT INTO payments (user_id, reservation_id, amount, status, payment_method, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, payment.UserID, payment.ReservationID, payment.Amount, payment.Status, payment.PaymentMethod, time.Now())
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	payment.ID = id
	return nil
}

func (r *paymentRepository) GetByID(id int64) (*models.Payment, error) {
    query := "SELECT id, user_id, reservation_id, amount, status, payment_method, created_at FROM payments WHERE id = ?"
    row := r.db.QueryRow(query, id)
    var payment models.Payment
    err := row.Scan(&payment.ID, &payment.UserID, &payment.ReservationID, &payment.Amount, &payment.Status, &payment.PaymentMethod, &payment.CreatedAt)
    if err != nil {
        // Log the error to see what exactly is happening.
        log.Printf("Error retrieving payment with id=%d: %v", id, err)
        return nil, err
    }
    return &payment, nil
}

// GetByUserID retrieves all payment records for a given user.
func (r *paymentRepository) GetByUserID(userID int64) ([]*models.Payment, error) {
	query := "SELECT id, user_id, reservation_id, amount, status, payment_method, created_at FROM payments WHERE user_id = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.UserID, &payment.ReservationID, &payment.Amount, &payment.Status, &payment.PaymentMethod, &payment.CreatedAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}
	return payments, nil
}

// NewMySQLDB creates and pings a MySQL database connection.
func NewMySQLDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
