**Barber Booking System - Payment Service**
================================================

The **Payment Service** is a microservice responsible for processing payments for booked barber appointments. It ensures secure transactions and records payment details. This service interacts with the **Reservation Service** to validate appointments before processing payments.

It is built using **Go (Gin framework)** and communicates with other microservices via **REST API** and  **RabbitMQ for asynchronous messaging**.

**Features**
------------

* **Secure Payment Processing** – Handles transactions for appointment bookings.
* **Integration with Reservation Service** – Ensures only valid reservations are processed.
* **Transaction Logging** – Stores successful and failed payment attempts.
* **Support for Multiple Payment Methods** – Can be extended to support credit cards, PayPal, etc.
* **Asynchronous Communication** – Uses  **RabbitMQ** to notify services of payment status.

**Tech Stack**
--------------

*   **Go (Gin framework)** – For building REST API
    
*   **MySQL** – Database for storing payment records
    
*   **RabbitMQ** – For event-driven communication
    
    



**API Endpoints (Planned)**
---------------------------
| Method | Endpoint                 | Description                       |
|--------|--------------------------|-----------------------------------|
| POST   | `/api/payments`          | Process a new payment            |
| GET    | `/api/payments/:id`      | Retrieve payment details by ID   |
| GET    | `/api/payments/user/:id` | Get payment history for a user   |
| POST   | `/api/payments/webhook`  | Handle payment gateway webhooks  |

**Setup Instructions (Coming Soon)**
------------------------------------

Instructions for setting up and running the payment service will be provided after the backend integration.