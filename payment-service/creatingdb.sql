CREATE DATABASE IF NOT EXISTS paymentdb;

USE paymentdb;

CREATE TABLE IF NOT EXISTS payments (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    barber_id BIGINT NOT NULL,
    reservation_id VARCHAR(50) NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    status VARCHAR(50) NOT NULL,
    payment_method VARCHAR(100) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO payments (user_id, reservation_id, barber_id, amount, status, payment_method)
VALUES
(1, '101', 10, 49.99, 'success', 'credit_card'),
(2, '102', 12, 75.00, 'pending', 'paypal');
