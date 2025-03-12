**Barber Booking System - User Service**
=============================================

**Project Description**
-----------------------

The **User Service** is a microservice responsible for managing users (customers and barbers) within the Barber Booking System. It handles **authentication, user profiles, and role-based access control**. This service ensures secure user data management and provides authentication tokens for accessing other services.

It is built using **Node.js (Express.js)** and communicates with other microservices via **REST API**.

**Features**
------------

* **User Registration & Login** – Secure authentication for customers and barbers.
* **JWT-Based Authentication** – Issues and validates JWT tokens for session management.
* **Role Management** – Differentiates between customers and barbers.
* **Profile Management** – Users can update personal details and passwords.
* **API Security** – Protects routes using authentication middleware.

**Tech Stack**
--------------

*   **Node.js (Express.js)** – REST API framework
    
*   **PostgreSQL** – Relational database for user storage
    
*   **bcrypt.js** – Password hashing for security
    
*   **jsonwebtoken (JWT)** – Authentication token handling
    
*   **Passport.js** – Middleware for authentication strategies
    



**API Endpoints (Planned)**
---------------------------

| Method  | Endpoint                | Description                        |
|---------|-------------------------|------------------------------------|
| POST    | `/api/users/register`   | Register a new user               |
| POST    | `/api/users/login`      | Authenticate user & return JWT    |
| GET     | `/api/users/profile`    | Retrieve logged-in user profile   |
| PATCH   | `/api/users/profile`    | Update user profile details       |
| DELETE  | `/api/users/:id`        | Delete user account (admin only)  |

**Setup Instructions (Coming Soon)**
------------------------------------

Instructions for setting up and running the User Service will be provided after backend development is completed.