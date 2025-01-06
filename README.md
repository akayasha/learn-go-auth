# Learn Go Auth

This is a simple authentication system built using the Go programming language and the Gin web framework. It allows users to register, log in, verify their email via OTP (One-Time Password), and resend OTPs when needed.

## Features

- **User Registration**: Allows new users to register with a username, email, password, and role.
- **Login**: Users can log in using their email and password.
- **Email Verification**: Sends an OTP to the user's email for verification.
- **Resend OTP**: Allows users to request a new OTP if needed.

## Installation

To get started with this project, clone the repository:

```bash
git clone https://github.com/akayasha/learn-go-auth
```

## Prerequisites

Make sure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

## Setup

1. Navigate to your project directory:
   ```bash
   cd learn-go-auth
   ```

2. Install the required dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file to store any environment variables you may need (such as database credentials or email service configuration).

## Running the Application

To run the application locally:

```bash
go run main.go
```

This will start the Gin server, and you can access the application at http://localhost:8080.

## API Endpoints

### `POST /register`

Registers a new user.

**Request Body:**
```json
{
    "username": "string",
    "email": "string",
    "password": "string",
    "role": "string"
}
```

- **username**: The desired username for the user.
- **email**: The email address for the user.
- **password**: The user's password.
- **role**: The role assigned to the user (e.g., "admin", "user").

**Response:**
```json
{
    "status": 201,
    "message": "User registered successfully",
    "data": {
        "user": { ... }
    }
}
```

### `POST /login`

Logs a user in.

**Request Body:**
```json
{
    "email": "string",
    "password": "string"
}
```

- **email**: The email address of the user.
- **password**: The user's password.

**Response:**
```json
{
    "status": 200,
    "message": "Login successful",
    "data": { ... }
}
```

### `POST /verify-email`

Verifies the user's email with an OTP.

**Request Body:**
```json
{
    "email": "string",
    "otp": "string"
}
```

- **email**: The email address to verify.
- **otp**: The OTP sent to the user's email for verification.

**Response:**
```json
{
    "status": 200,
    "message": "Email verified successfully",
    "data": null
}
```

### `GET /resend-otp`

Resends the OTP to the user's email.

**Query Parameters:**
- `email`: The email address to resend the OTP to.

**Response:**
```json
{
    "status": 200,
    "message": "OTP resent successfully",
    "data": null
}
```

## Error Handling

If there are any issues with the request, the API will return an error with a status code and a message explaining the issue. Example:

```json
{
    "status": 400,
    "message": "Invalid request data",
    "data": null
}
```

Common Error Status Codes:
- **400 Bad Request**: The request is malformed or missing required data.
- **401 Unauthorized**: The provided credentials are invalid.
- **404 Not Found**: The requested resource was not found.
- **500 Internal Server Error**: An unexpected error occurred on the server.

## Folder Structure

- `controllers/`: Contains the handlers for different routes (e.g., register, login, verify email).
- `services/`: Contains the business logic for interacting with the database or external services.
- `utils/`: Contains utility functions, such as response helpers.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - A fast HTTP web framework for Go.
- [dotenv](https://github.com/joho/godotenv) - Loads environment variables from a `.env` file.
- Any other libraries for services like sending OTPs, email verifications, etc.

## License

This project is open-source and available under the MIT License.
