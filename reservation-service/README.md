**Barber Booking System - Reservation Service**
====================================================

The **Reservation Service** is a microservice responsible for handling appointment scheduling in the Barber Booking System. It allows customers to book, modify, or cancel appointments while enabling barbers to manage their availability. This service ensures that appointments are properly validated and stored in the database.

It is built using **Python (FastAPI)** and communicates with other microservices via **REST API**. Additionally, it integrates with the **Payment Service** using **RabbitMQ** for asynchronous processing of payment confirmations.

**Features**
------------

* **Appointment Booking** – Customers can reserve available time slots.
* **Manage Availability** – Barbers can set and update their schedules.
* **Cancel & Modify Appointments** – Users can reschedule or cancel bookings.
* **Integration with Payment Service** – Ensures payments are processed before confirming appointments.
* **Real-time Updates** – Sends notifications (via RabbitMQ) about appointment status changes.

**Tech Stack**
--------------

*   **Python (FastAPI)** – For handling API requests
    
*   **MongoDB** – Document database for flexible appointment storage
    
*   **RabbitMQ** – For asynchronous messaging between services
    
*   **Pydantic** – Data validation and serialization
    

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

**Setup Instructions (Coming Soon)**
------------------------------------

Instructions for setting up and running the reservation service will be added after backend development is complete.