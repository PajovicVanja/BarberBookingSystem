// internal/repository/payment_repository.go
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"paymentservice/internal/models"
	"time"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id int64) (*models.Payment, error)
	GetByUserID(userID int64) ([]*models.Payment, error)
	GetByBarberID(barberID int64) ([]*models.Payment, error)
	Delete(id int64) error
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

// Create inserts a new payment record into the database.
func (r *paymentRepository) Create(payment *models.Payment) error {
	query := `
	INSERT INTO payments
	  (user_id, reservation_id, barber_id, amount, status, payment_method, created_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	result, err := r.db.Exec(
		query,
		payment.UserID,
		payment.ReservationID,
		payment.BarberID,
		payment.Amount,
		payment.Status,
		payment.PaymentMethod,
		time.Now(),
	)
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
	query := `
	SELECT id, user_id, reservation_id, barber_id, amount, status, payment_method, created_at
	FROM payments
	WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var p models.Payment
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.ReservationID,
		&p.BarberID,
		&p.Amount,
		&p.Status,
		&p.PaymentMethod,
		&p.CreatedAt,
	)
	if err != nil {
		log.Printf("Error retrieving payment id=%d: %v", id, err)
		return nil, err
	}
	return &p, nil
}

func (r *paymentRepository) GetByUserID(userID int64) ([]*models.Payment, error) {
	query := `
	SELECT id, user_id, reservation_id, barber_id, amount, status, payment_method, created_at
	FROM payments
	WHERE user_id = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		var p models.Payment
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.ReservationID,
			&p.BarberID,
			&p.Amount,
			&p.Status,
			&p.PaymentMethod,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, nil
}

func (r *paymentRepository) GetByBarberID(barberID int64) ([]*models.Payment, error) {
	query := `
	SELECT id, user_id, reservation_id, barber_id, amount, status, payment_method, created_at
	FROM payments
	WHERE barber_id = ?`
	rows, err := r.db.Query(query, barberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*models.Payment
	for rows.Next() {
		var p models.Payment
		if err := rows.Scan(
			&p.ID,
			&p.UserID,
			&p.ReservationID,
			&p.BarberID,
			&p.Amount,
			&p.Status,
			&p.PaymentMethod,
			&p.CreatedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, &p)
	}
	return payments, nil
}

// Delete removes a payment record by ID.
func (r *paymentRepository) Delete(id int64) error {
	query := `DELETE FROM payments WHERE id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("payment id %d not found", id)
	}
	return nil
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
