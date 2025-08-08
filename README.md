# Task Management API

A RESTful API for managing tasks built with Go, Echo framework, GORM, and PostgreSQL.

## Features

- CRUD operations for tasks
- Basic authentication for write operations
- Pagination and filtering
- Clean Architecture with DDD principles
- Auto-reload with Air
- Docker support
- Graceful shutdown with context cancellation
- Structured logging with slog
- Context-aware database operations

## Tech Stack

- **Framework**: Echo v4
- **Database**: PostgreSQL with GORM
- **Auto-reload**: Air
- **Authentication**: Basic Auth
- **Architecture**: Clean Architecture with DDD
- **Logging**: Structured logging with slog
- **Context**: Context cancellation pattern
- **Configuration**: envconfig for structured config management

## Project Structure

```
task-be/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── domain/                 # Domain layer (entities, interfaces)
│   │   ├── task.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── application/            # Application layer (use cases)
│   │   └── service/
│   │       └── task_service.go
│   │       └── task_service_test.go
│   ├── infrastructure/         # Infrastructure layer (external concerns)
│   │   ├── database/
│   │   │   └── database.go
│   │   ├── config/
│   │   │   └── config.go
│   │   ├── logger/
│   │   │   └── logger.go
│   │   ├── repository/
│   │   │   └── task_repository.go
│   │   └── middleware/
│   │       └── auth.go
│   └── interfaces/             # Interface layer (HTTP handlers)
│       └── http/
│           ├── dto/
│           │   └── task_dto.go
│           ├── handler/
│           │   └── task_handler.go
│           └── router/
│               └── router.go
├── Dockerfile
├── docker-compose.yml
├── .air.toml
├── go.mod
└── env.example
└── README.md
```

## API Endpoints

### Public Endpoints (No Authentication Required)
- `GET /tasks` - Get all tasks with pagination and filtering
- `GET /tasks/:id` - Get a specific task by ID

### Protected Endpoints (Basic Authentication Required)
- `POST /tasks` - Create a new task
- `PATCH /tasks/:id` - Update an existing task
- `DELETE /tasks/:id` - Delete a task

### Authentication
Use Basic Authentication with the following credentials:
- Username: `admin`
- Password: `password123`

## Setup Instructions

### Prerequisites
- Go 1.21 or higher
- PostgreSQL (for local development) or Docker (for containerized setup)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/usernamesalah/task-be.git
   cd task-be
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables**
   ```bash
   cp env.example .env
   # Edit .env file with your configuration
   ```

4. **Run with Air (auto-reload)**
   ```bash
   # Install Air if not already installed
   go install github.com/cosmtrek/air@latest
   
   # Run the application
   air
   ```

5. **Run tests**
   ```bash
   go test ./...
   ```

### Docker Setup

1. **Build and run with Docker Compose (includes PostgreSQL)**
   ```bash
   docker-compose up --build
   ```

2. **Run in background**
   ```bash
   docker-compose up -d
   ```

3. **Stop the application**
   ```bash
   docker-compose down
   ```

4. **Stop and remove volumes**
   ```bash
   docker-compose down -v
   ```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `READ_TIMEOUT` | Server read timeout | `30s` |
| `WRITE_TIMEOUT` | Server write timeout | `30s` |
| `IDLE_TIMEOUT` | Server idle timeout | `60s` |
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | PostgreSQL username | `postgres` |
| `DB_PASSWORD` | PostgreSQL password | `password` |
| `DB_NAME` | PostgreSQL database name | `taskdb` |
| `DB_SSLMODE` | PostgreSQL SSL mode | `disable` |
| `BASIC_AUTH_USERNAME` | Basic auth username | `admin` |
| `BASIC_AUTH_PASSWORD` | Basic auth password | `password123` |

## API Examples

### Create a Task
```bash
curl -X POST http://localhost:3000/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=" \
  -d '{
    "title": "Complete project",
    "description": "Finish the task management API"
  }'
```

### Get All Tasks
```bash
curl http://localhost:3000/tasks
```

### Get Tasks with Pagination
```bash
curl "http://localhost:3000/tasks?page=1&limit=10"
```

### Filter Tasks by Status
```bash
curl "http://localhost:3000/tasks?status=TO_DO"
```

### Update a Task
```bash
curl -X PATCH http://localhost:3000/tasks/1 \
  -H "Content-Type: application/json" \
  -H "Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM=" \
  -d '{
    "status": "IN_PROGRESS"
  }'
```

### Delete a Task
```bash
curl -X DELETE http://localhost:3000/tasks/1 \
  -H "Authorization: Basic YWRtaW46cGFzc3dvcmQxMjM="
```

## Design Decisions

1. **Clean Architecture**: Separated concerns into domain, application, infrastructure, and interface layers
2. **DDD Principles**: Used domain entities and value objects for business logic
3. **PostgreSQL**: Production-ready database with ACID compliance
4. **Basic Auth**: Echo's built-in basic authentication middleware
5. **Echo Framework**: Lightweight and fast HTTP framework
6. **GORM**: ORM for database operations with auto-migration
7. **Context Pattern**: Used throughout the application for cancellation and timeouts
8. **Structured Logging**: JSON-formatted logs with structured fields for better observability
9. **Graceful Shutdown**: Proper cleanup of resources on application termination
10. **Configuration Management**: Structured configuration using envconfig

## Advanced Features

### Graceful Shutdown
The application implements graceful shutdown with a 30-second timeout:
- Handles SIGINT and SIGTERM signals
- Properly closes HTTP server
- Logs shutdown process with structured logging

### Context Cancellation Pattern
- All database operations use context for cancellation
- HTTP handlers pass request context to business logic
- Supports timeout and cancellation throughout the call chain

### Structured Logging
- JSON-formatted logs using Go's `slog` package
- Structured fields for better observability
- Log levels: Info, Warn, Error
- Includes request context (IP, path, user) in authentication logs

### Context-Aware Operations
- Database queries respect context cancellation
- Repository layer propagates context to GORM
- Service layer logs operations with context information

### Configuration Management
- Structured configuration using envconfig
- Type-safe configuration with defaults
- Environment variable validation
- Separate configuration for server, database, and auth

## Testing

The project includes unit tests for the service layer. Run tests with:
```bash
go test ./...
```
