# Notification Engine

Notification Engine is a lightweight, asynchronous, and scalable microservice designed to handle real-time alerts and notifications for the Quant Trading Engine. It is built using **Go** and follows the **Hexagonal Architecture** (Ports and Adapters) pattern.

Currently, it supports delivering notifications via **Telegram Bot API**.

## 🏗 Architecture

The project is structured around the Hexagonal Architecture:
- `internal/core/domain`: Contains core business logic and models (e.g., `Notification` entity).
- `internal/core/ports`: Contains interfaces that define how the core interacts with the outside world.
- `internal/core/services`: Contains the application logic that orchestrates domain objects and ports.
- `internal/adapters/handlers`: Contains HTTP delivery adapters (Echo framework) to expose the core via REST API.
- `internal/adapters/messengers`: Contains infrastructure adapters for sending messages (e.g., Telegram).
- `cmd/api`: The entry point of the application, responsible for wiring up dependencies and starting the server.

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) 1.22 or higher
- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- [Make](https://www.gnu.org/software/make/)

### Configuration

Create a `.env` file in the root directory and populate it with your Telegram credentials and a secure API key:

```env
PORT=8080
API_KEY=your_secret_api_key_here
TELEGRAM_BOT_TOKEN=your_telegram_bot_token
TELEGRAM_CHAT_ID=your_telegram_chat_id
```

## 🛠 Commands

This project uses a `Makefile` to simplify common operations:

- **`make start`**: Run the application locally without compiling a binary.
- **`make build`**: Compile the application into a binary file (`bin/notification-engine`).
- **`make dev`**: Run the application using `air` for hot-reloading during development.
- **`make swagger`**: Generate/update the Swagger API documentation.
- **`make docker-build`**: Build the Docker image (`notification-engine:latest`).
- **`make docker-run`**: Run the application inside a Docker container.

## 📚 API Documentation

Once the server is running, you can access the Swagger UI documentation at:
👉 **`http://localhost:8080/swagger/index.html`**

### Send Notification
**`POST /api/v1/notify`**

**Headers:**
- `X-API-Key: {your_secret_api_key_here}`
- `Content-Type: application/json`

**Body:**
```json
{
  "level": "INFO",
  "source": "Quant Engine",
  "message": "Your Markdown Formatted Message Here"
}
```

## 🔐 Security
The API is protected by a simple middleware that checks the `X-API-Key` header against the `API_KEY` defined in the `.env` file. Unauthorized requests will be rejected with a `401 Unauthorized` status code.
