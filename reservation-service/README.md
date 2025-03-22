**Barber Booking System - Reservation Service**
====================================================

The **Reservation Service** is a microservice responsible for handling appointment scheduling in the Barber Booking System. It allows customers to book, modify, or cancel appointments while enabling barbers to manage their availability. This service ensures that appointments are properly validated and stored in the database.

It is built using **Python (FastAPI)** and communicates with other microservices via **REST API and gRPC**. Additionally, it integrates with the **Payment Service** using **RabbitMQ** for asynchronous processing of payment confirmations.

**Features**
------------

* **Appointment Booking** – Customers can reserve available time slots.
* **Manage Availability** – Barbers can set and update their schedules.
* **Cancel & Modify Appointments** – Users can reschedule or cancel bookings.
* **Integration with Payment Service** – Ensures payments are processed before confirming appointments.
* **Real-time Updates** – Sends notifications (via RabbitMQ) about appointment status changes.
* **gRPC Support** – High-performance RPC interface for internal microservice communication.

**Tech Stack**
--------------

*   **Python (FastAPI & gRPC)** – For handling API and gRPC requests
    
*   **MongoDB** – Document database for flexible appointment storage
    
*   **RabbitMQ** – For asynchronous messaging between services
    
*   **Pydantic** – Data validation and serialization

*   **gRPC (grpcio, grpcio-tools)** – Inter-service communication

*   **Protocol Buffers (protobuf)** – Interface definition for gRPC

**API Endpoints (Planned)**
---------------------------

| Method  | Endpoint                      | Description                          |
|---------|--------------------------------|--------------------------------------|
| POST    | `/api/reservations`           | Create a new appointment            |
| GET     | `/api/reservations/:id`       | Retrieve appointment details        |
| GET     | `/api/reservations/user/:id`  | Get all bookings for a user         |
| PATCH   | `/api/reservations/:id`       | Modify an existing appointment      |
| DELETE  | `/api/reservations/:id`       | Cancel an appointment               |
| POST    | `/api/reservations/confirm`   | Confirm payment & finalize booking  |


### gRPC Services

| RPC Method             | Description                          |
|------------------------|--------------------------------------|
| `CreateReservation`    | Create a new appointment             |
| `GetReservation`       | Retrieve details of a reservation    |
| `CancelReservation`    | Cancel an appointment                |
| `UpdateReservation`    | Modify an existing appointment       |
| `ListUserReservations` | List all reservations for a user     |
| `ConfirmPayment`       | Finalize reservation after payment   |

## Setup Instructions
1. **Install Dependencies:**  
   ```bash
   pip install -r requirements.txt
Generate gRPC Stubs:

bash
Copy
python -m grpc_tools.protoc -I./app/proto --python_out=./app --grpc_python_out=./app app/proto/reservation.proto
Configure MongoDB and RabbitMQ:
Ensure MongoDB and RabbitMQ are running locally or adjust the environment variables accordingly.

Run the FastAPI Server:

bash
Copy
uvicorn app.main:app --host 0.0.0.0 --port 8000
Run the gRPC Server:

bash
Copy
python app/grpc_server.py
Run Tests:

bash
Copy
pytest
Docker
To build and run the Docker container:

bash
Copy
docker build -t reservation-service .
docker run -p 8000:8000 -p 50051:50051 reservation-service



🐳 Docker Instructions
Build the Docker image:

bash
Copy
Edit
docker build -t reservation-service .
Run the container:

bash
Copy
Edit
docker run -p 8000:8000 -p 50051:50051 reservation-service-ita
Open a shell inside the container:

bash
Copy
Edit
docker exec -it <docker id> /bin/bash
Run tests inside the container:

bash
Copy
Edit
PYTHONPATH=/app pytest -s tests/