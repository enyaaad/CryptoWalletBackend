# CryptoWallet Backend

Cryptocurrency wallet with account, wallet, and transaction management functionality.

📋 [Changelog](./CHANGELOG.md) - Project change history

## What it does

MVP Plan:

- User registration and authentication with JWT tokens
- Creation and management of cryptocurrency wallets
- Sending and receiving cryptocurrencies
- Transaction history
- P2P exchange between users
- Blockchain integration (Bitcoin, Ethereum, Solana)
- Real-time monitoring and metrics

## Architecture

The project is built on a microservices architecture using Clean Architecture:

```
cmd/
├── api-gateway/          # HTTP REST API for frontend (Gin)
├── auth-service/         # Authentication and authorization (gRPC)
├── wallet-service/       # Wallet management (gRPC)
├── transaction-service/  # Transaction processing (gRPC)
├── p2p-service/         # P2P exchange (gRPC)
└── exchange-service/    # Cryptocurrency exchange (gRPC)

internal/
├── api-gateway/
│   ├── delivery/http/   # HTTP handlers (Gin) for frontend
│   ├── infrastructure/grpc/  # gRPC clients for calling microservices
│   └── middleware/      # Logging, metrics
├── auth/                 # Auth microservice
│   ├── domain/          # Business logic (entities, interfaces)
│   ├── infrastructure/  # External systems (PostgreSQL, JWT, Password)
│   └── delivery/grpc/   # gRPC server
├── wallet/              # Wallet microservice
└── ...

api/
└── proto/               # Proto files for gRPC
    └── auth.proto

pkg/                     # Shared components
├── logger/              # Structured logging
└── metrics/             # Prometheus metrics
```

### Service Communication

- **Frontend ↔ API Gateway**: HTTP REST API (Gin)
- **API Gateway ↔ Microservices**: gRPC
- **Microservices ↔ Microservices**: gRPC

## Technologies

### Core Stack
- **Go 1.23** - primary programming language
- **Gin** - HTTP framework for REST API (API Gateway)
- **gRPC** - inter-service communication between microservices
- **Protocol Buffers** - contract definition for gRPC
- **PostgreSQL** - primary database
- **Redis** - caching and sessions
- **Apache Kafka** - event-driven communication between services

### Monitoring and Observability
- **Prometheus** - metrics collection
- **Grafana** - metrics visualization and dashboards
- **Zerolog** - structured logging

### Infrastructure
- **Docker Compose** - local development
- **golang-migrate** - database migrations
- **JWT** - authentication and authorization

## Getting Started

1. Install dependencies:
```bash
go mod download
```

2. Start infrastructure (PostgreSQL, Redis, Prometheus, Grafana):
```bash
make up
```

3. Apply migrations:

###TODO

4. Start Auth Service (gRPC server on port 50051):

###TODO

5. Start API Gateway (HTTP REST API on port 8080):
```bash
go run cmd/api-gateway/main.go
```

## API Endpoints

### TODO

## Development

### Project Structure
The project follows Clean Architecture principles:
- **Domain layer** contains business logic and interfaces (does not depend on external systems)
- **Infrastructure layer** implements external dependencies (DB, gRPC clients/servers)
- **Delivery layer** handles requests (HTTP handlers for API Gateway, gRPC handlers for microservices)

### Microservices

Each microservice:
- Has its own entry point in `cmd/<service-name>/main.go`
- Contains full implementation in `internal/<service-name>/`
- Provides gRPC API for other services
- Has its own database (or schema in shared DB)

### Testing

### TODO

### Migrations
```bash
# Create a new migration
migrate create -ext sql -dir migrations -seq add_users_table

# Apply migrations
migrate -path migrations -database "postgres://cryptowallet:cryptowallet123@localhost:5432/cryptowallet?sslmode=disable" up
```

### gRPC Code Generation
```bash
# After modifying proto files
protoc --go_out=. --go-grpc_out=. api/proto/*.proto
```

## Service Ports

- **API Gateway**: 8080 (HTTP REST API)
- **Auth Service**: 50051 (gRPC)
- **Wallet Service**: 50052 (gRPC)
- **Transaction Service**: 50053 (gRPC)
- **PostgreSQL**: 5432
- **Redis**: 6379
- **Prometheus**: 9090
- **Grafana**: 3000

## Project Goals

- Learning modern backend development approaches
- Practicing microservices architecture
- Mastering gRPC for inter-service communication
- Mastering event-driven approaches
- Implementing monitoring
- Applying Go development best practices
