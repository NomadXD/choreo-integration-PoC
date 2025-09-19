# CRM Backend Go Service

A simple Go REST API service that accepts ticket submissions via POST requests and publishes them to a RabbitMQ message queue.

## Features

- REST API endpoint: `POST /ticket`
- Accepts JSON payload with ticket information
- Publishes tickets to RabbitMQ queue
- Health check endpoint: `GET /health`
- CORS support
- Environment-based configuration

## API Endpoints

### POST /ticket

Accepts a ticket submission and publishes it to RabbitMQ.

**Request Body:**

```json
{
  "ticketId": "INC-33",
  "customerName": "Jane Doe",
  "customerEmail": "lahiru97udayanga@gmail.com",
  "issue": "Cannot reset password."
}
```

**Response:**

```json
{
  "success": true,
  "message": "Ticket submitted successfully",
  "ticketId": "INC-33"
}
```

### GET /health

Health check endpoint.

**Response:**

```json
{
  "status": "ok",
  "message": "CRM Backend Service is running"
}
```

## Prerequisites

- Go 1.21 or higher
- RabbitMQ server running on default port (5672)

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

The service uses environment variables for configuration. Copy `.env.example` to `.env` and modify as needed:

- `PORT`: Server port (default: 8080)
- `RABBITMQ_URL`: RabbitMQ connection URL (default: amqp://guest:guest@localhost:5672/)
- `QUEUE_NAME`: RabbitMQ queue name (default: tickets)

## Running the Service

1. Start RabbitMQ server:

   ```bash
   # Using Docker
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

   # Or using Homebrew on macOS
   brew services start rabbitmq
   ```

2. Run the service:
   ```bash
   go run main.go
   ```

The service will start on port 8080 (or the port specified in the PORT environment variable).

## Testing

You can test the API using curl:

```bash
curl -X POST http://localhost:8080/ticket \
  -H "Content-Type: application/json" \
  -d '{
    "ticketId": "INC-33",
    "customerName": "Jane Doe",
    "customerEmail": "lahiru97udayanga@gmail.com",
    "issue": "Cannot reset password."
  }'
```

## Project Structure

```
.
├── main.go                    # Application entry point
├── go.mod                     # Go module file
├── .env.example              # Environment variables example
├── README.md                 # This file
└── internal/
    ├── models/
    │   └── ticket.go         # Ticket data models
    ├── handlers/
    │   └── ticket.go         # HTTP handlers
    └── rabbitmq/
        └── client.go         # RabbitMQ client implementation
```

## Dependencies

- [Gin](https://github.com/gin-gonic/gin) - HTTP web framework
- [AMQP](https://github.com/streadway/amqp) - RabbitMQ client library
