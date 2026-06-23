.PHONY: build start dev migration-up migration-down swagger

# Build the application binary
build:
	@echo "Building notification-engine..."
	go build -o bin/notification-engine cmd/api/main.go

# Start the built application
start: build
	@echo "Starting notification-engine..."
	./bin/notification-engine

# Run the application in development mode
dev:
	@echo "Running in development mode..."
	go run cmd/api/main.go

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g cmd/api/main.go

# Database Migration placeholders (for future use if DB is added)
MIGRATION_DIR=migrations
DB_URL="postgresql://user:password@localhost:5432/dbname?sslmode=disable"

migration-up:
	@echo "Running database migrations (UP)..."
	# Example using golang-migrate: 
	# migrate -path $(MIGRATION_DIR) -database $(DB_URL) up
	@echo "No migrations configured yet."

migration-down:
	@echo "Running database migrations (DOWN)..."
	# Example using golang-migrate: 
	# migrate -path $(MIGRATION_DIR) -database $(DB_URL) down
	@echo "No migrations configured yet."

docker-build:
	@echo "Building Docker image..."
	docker build -t notification-engine:latest .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env notification-engine:latest
