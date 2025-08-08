.PHONY: build run test clean docker-build docker-run

# Build the application
build:
	go build -o bin/main ./cmd/main.go

# Run the application
run:
	go run ./cmd/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f database.db

# Build Docker image
docker-build:
	docker build -t task-be .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Stop Docker Compose
docker-stop:
	docker-compose down

# Install Air for development
install-air:
	go install github.com/cosmtrek/air@latest

# Run with Air (auto-reload)
dev:
	air
