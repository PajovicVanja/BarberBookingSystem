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

### 1\. Prerequisites

Make sure you have the following installed:

*   [Go](https://go.dev/dl/) (version 1.21 or higher)
    
*   [MySQL](https://dev.mysql.com/downloads/mysql/) (or an alternative like XAMPP)
    
*   [RabbitMQ](https://www.rabbitmq.com/download.html)
    


### 2\. Set Up the MySQL Database

1.  Start your MySQL server.
    
2.  Create a database named paymentdb and a user with access:
    

`   CREATE DATABASE paymentdb;  CREATE USER 'root'@'localhost' IDENTIFIED BY 'root';  GRANT ALL PRIVILEGES ON paymentdb.* TO 'root'@'localhost';  FLUSH PRIVILEGES;   `

> You can customize credentials, but make sure they match the DATABASE\_DSN value in the next step.

### 3\. Set Up Environment Variables

Create a .env file or export environment variables directly:

`   export SERVER_PORT=8080  export DATABASE_DSN=root:root@tcp(localhost:3306)/paymentdb  export RABBITMQ_URL=amqp://guest:guest@localhost:5672/   `

### 4\. Download Dependencies

`   go mod download   `

### 5\. Build and Run the Service

`   go build -o paymentservice ./cmd/payment-service  ./paymentservice   `

The service should be running at [http://localhost:8080](http://localhost:8080)

### 6\. RabbitMQ Setup

If not already running:

*   Start RabbitMQ locally.
    
*   Access the management UI at [http://localhost:15672](http://localhost:15672)
    
*   Default credentials: guest / guest