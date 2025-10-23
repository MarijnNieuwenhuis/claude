# Bootstrap Go Service

A bootstrap/skeleton project for Go microservices at BTCDirect, featuring standard infrastructure patterns and integrations.

## Features

- **Standard Application Setup**: Pre-configured with BTCDirect go-modules for app lifecycle, HTTP, logging, messaging, and SQL
- **Database Support**: MySQL/MariaDB with Cloud SQL connector and migration framework
- **HTTP Server**: Gorilla Mux router with health/readiness endpoints
- **Pub/Sub Messaging**: Google Cloud Pub/Sub integration with local emulator support
- **Environment Configuration**: Support for dev, stage, acc, sandbox, and prod environments
- **Structured Logging**: Zap-based logging throughout
- **Sentry Integration**: Error tracking and monitoring
- **Docker**: Multi-stage build with distroless final image
- **CI/CD**: GitLab CI/CD integration with BTCDirect templates

## Getting Started

### Prerequisites

- Go 1.25+
- MySQL/MariaDB database
- (Optional) Google Cloud Pub/Sub emulator for local development

### Installation

1. Clone this repository to create your new service:
```bash
git clone <this-repo> my-new-service
cd my-new-service
```

2. Rename the module in `go.mod`:
```bash
# Replace "gitlab.com/btcdirect-api/bootstrap-go-service" with your service name
# e.g., "gitlab.com/btcdirect-api/my-service"
```

3. Update all import statements:
```bash
find . -type f -name "*.go" -exec sed -i '' 's/gitlab.com\/btcdirect-api\/bootstrap-go-service/gitlab.com\/btcdirect-api\/my-service/g' {} +
```

4. Rename the cmd directory:
```bash
mv cmd/bootstrap-go-service cmd/my-service
```

5. Update `Dockerfile` and `Makefile` to reference your new service name

6. Configure your `.env` file with your database and service settings

### Running Locally

```bash
# Run the service
make run

# Run database migrations
make migrate

# Run tests
make test
```

## Project Structure

```
.
├── cmd/bootstrap-go-service/    # Application entry point
├── internal/
│   ├── app/                     # Application initialization and config
│   ├── db/                      # Database connection and migrations
│   ├── http/
│   │   ├── handler/            # HTTP handlers
│   │   └── server/             # Server setup and routing
│   └── messenger/
│       ├── inbound/            # Message consumers (webhook pattern)
│       └── outbound/           # Message publishers (event pattern)
├── vendor/                     # BTCDirect go-modules
│   └── gitlab.com/btcdirect-api/go-modules/
│       ├── app/               # Application lifecycle
│       ├── http/              # HTTP utilities
│       ├── logger/            # Logging
│       ├── messenger/         # Pub/Sub messaging
│       └── sql/               # Database utilities
├── Dockerfile                 # Multi-stage Docker build
├── Makefile                   # Build automation
└── .gitlab-ci.yml            # CI/CD pipeline

```

## Adding Your Business Logic

### 1. Database Migrations

Add SQL migration files to `internal/db/migrations/`:
```sql
-- 001_create_users_table.sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. HTTP Routes

Add your routes in `internal/http/server/routes.go`:
```go
r.HandleFunc("/users", handler.ListUsers(app)).Methods("GET")
r.HandleFunc("/users/{id}", handler.GetUser(app)).Methods("GET")
```

### 3. Message Handlers

Implement message handlers in `internal/messenger/inbound/` and register them in `internal/app/app.go`

### 4. Business Services

Add your business logic in new packages under `internal/`

## Configuration

Environment variables (configure in `.env`):

- `APP_ENV`: Environment (dev, stage, acc, sandbox, prod)
- `HTTP_PORT`: HTTP server port (default: 8080)
- `LOG_LEVEL`: Logging level (debug, info, warn, error)
- `DATABASE_URL`: MySQL connection string
- `SENTRY_DSN`: Sentry error tracking DSN
- `PUBSUB_EMULATOR`: Pub/Sub emulator host (for local dev)
- `PUBSUB_PROJECT`: Google Cloud project ID

## Building

### Local Build
```bash
go build -o app ./cmd/bootstrap-go-service
```

### Docker Build
```bash
docker build -t my-service .
```

## Testing

```bash
go test ./internal/... ./pkg/...
```

## Deployment

The service includes a `.gitlab-ci.yml` file configured for BTCDirect's CI/CD pipeline. Push to your GitLab repository to trigger automated builds and deployments.

## License

Proprietary - BTCDirect