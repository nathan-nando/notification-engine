# Notification Engine

Notification Engine is a lightweight, asynchronous, and scalable microservice designed to handle real-time alerts and **2-way communication** for the Quant Trading Engine. It is built using **Go** and follows the **Hexagonal Architecture** (Ports and Adapters) pattern.

Currently, it supports:
1. **Push Notifications:** Delivering automated alerts via **Telegram Bot API** (e.g., trade executions, errors, snapshot reports).
2. **Control Plane:** A bi-directional Telegram chatbot interface that allows the user to query engine data and execute operational actions (e.g., stopping the engine, changing accounts, checking PnL) from Telegram.

## 🏗 Architecture

The project is structured around the Hexagonal Architecture:
- `internal/core/domain`: Contains core business logic and models (e.g., `Notification` entity).
- `internal/core/ports`: Contains interfaces that define how the core interacts with the outside world.
- `internal/core/services`: Contains the application logic that orchestrates domain objects and ports.
- `internal/adapters/handlers`: Contains HTTP delivery adapters (Echo framework) to expose the core via REST API.
- `internal/adapters/messengers`: Contains infrastructure adapters for sending messages and receiving chat commands via Telegram.
- `internal/adapters/broker`: Contains the Redis Pub/Sub adapter to route chat commands to the Python engine.
- `cmd/api`: The entry point of the application, responsible for wiring up dependencies and starting the server.

### Control Plane Flow (Redis Pub/Sub)
1. The user types a command in Telegram (e.g. `/porto`).
2. The Go `notification-engine` receives the webhook/polling event, parses it, and publishes it to the Redis channel `engine:control:requests`.
3. The Python `quant-engine-v1` listens to this channel, processes the logic, queries the database, and publishes the formatted response to `engine:control:responses`.
4. The Go `notification-engine` reads the response and forwards it back to the user's Telegram chat.

## 🚀 Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) 1.22 or higher
- [Docker](https://docs.docker.com/get-docker/) & Docker Compose
- [Make](https://www.gnu.org/software/make/)
- Redis Server (for Control Plane Pub/Sub)

### Configuration

Create a `.env` file in the root directory and populate it with your Telegram credentials and a secure API key:

```env
PORT=8080
API_KEY=your_secret_api_key_here
TELEGRAM_BOT_NOTIF_TOKEN=your_telegram_bot_token
TELEGRAM_BOT_CONTROL_TOKEN=your_control_plane_bot_token
TELEGRAM_CHAT_ID=your_telegram_chat_id
REDIS_URL=redis://localhost:6379
```

## 🤖 Control Plane Commands
The Control Plane bot supports the following commands:

**ℹ️ Informational**
- `/status` - Lihat kondisi engine, koneksi, & cron interval
- `/porto` - Lihat saldo, equity, dan margin live MT5
- `/accounts` - Lihat daftar akun trading (MT5)
- `/signals` - Lihat 5 sinyal terakhir
- `/signal <id>` - Lihat detail sinyal by ID
- `/positions` - Lihat 5 posisi terbuka terakhir
- `/position <id>` - Lihat detail posisi/trade by ID
- `/orders` - Lihat 5 order historis terakhir
- `/risk` atau `/threshold` - Lihat parameter risiko
- `/errors` atau `/logs` - Lihat log error terakhir

**⚡ Action / Control**
- `/change_account <id>` - Ganti koneksi akun MT5 live/demo/paper
- `/start_all` - Nyalakan engine cron & eksekusi auto
- `/stop_all` - Matikan engine cron & eksekusi auto
- `/start_trade` - Nyalakan eksekusi trade otomatis
- `/stop_trade` - Matikan eksekusi trade otomatis
- `/start_cron` - Nyalakan engine (cron) saja
- `/stop_cron` - Matikan engine (cron) saja
- `/set_cron <mins>` - Ubah interval waktu eksekusi cron
- `/flatten_all` - Tutup semua posisi terbuka (panic button)

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
