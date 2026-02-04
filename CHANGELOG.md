# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Planned
- Wallet Service implementation
- Transaction Service implementation
- P2P Service implementation
- Exchange Service implementation
- Redis integration for caching and sessions
- Kafka integration for event-driven architecture
- WebSocket support for real-time updates

## [0.1.0] - 05 February of 2026

### Added

#### Infrastructure
- Docker Compose configuration with all required services:
  - PostgreSQL (TimescaleDB) for data storage
  - Redis for caching and sessions
  - Kafka + Zookeeper for event streaming
  - Prometheus for metrics collection
  - Grafana for metrics visualization
- Health check endpoints (`/health`) for API Gateway
- Prometheus metrics endpoint (`/metrics`) for all services
- Docker network configuration for service discovery

#### API Gateway
- HTTP REST API server using Gin framework
- Versioned API routes (`/api/v1/...`)
- Structured logging middleware with zerolog
- Prometheus metrics middleware for HTTP requests
- gRPC client infrastructure for inter-service communication
- Graceful shutdown support

#### Auth Service
- gRPC server implementation for authentication
- User registration endpoint (`POST /api/v1/auth/register`)
- User login endpoint (`POST /api/v1/auth/login`)
- JWT token generation (access and refresh tokens)
- Token validation endpoint (`POST /api/v1/auth/validate`)
- Token refresh endpoint (`POST /api/v1/auth/refresh`)
- Password hashing using bcrypt
- User repository with PostgreSQL implementation
- Clean Architecture structure (domain, infrastructure, delivery layers)

#### Database
- Database migrations using golang-migrate
- Users table with email, username, password_hash
- Indexes for performance optimization
- Migration rollback support

#### Observability
- Structured logging with zerolog (JSON format)
- Prometheus metrics collection:
  - HTTP request metrics (total, duration, status codes)
  - Custom business metrics support
- Grafana provisioning:
  - Prometheus datasource configuration
  - HTTP metrics dashboard
- Prometheus configuration with service discovery

#### Development Tools
- Makefile for common development tasks
- Environment variable configuration support
- Docker service templates
- Project documentation structure

### Changed
- Initial project structure following Clean Architecture principles
- Separation of concerns: domain, infrastructure, delivery layers

### Security
- Password hashing with bcrypt
- JWT token-based authentication
- Secure token storage and validation

---

## [Unreleased] - Future Releases

### Planned Features

#### Wallet Service
- Multi-blockchain wallet support (Bitcoin, Ethereum, Solana, BSC)
- Wallet creation and management
- Address generation
- Balance tracking
- Redis caching for balances

#### Transaction Service
- Transaction creation and processing
- Transaction status tracking
- Transaction history with filtering
- Event sourcing with Kafka
- WebSocket real-time updates

#### P2P Service
- P2P order creation and matching
- Escrow system for secure transactions
- User ratings and reviews
- Chat functionality for deals

#### Exchange Service
- Order book management
- Matching engine
- Trading pairs support
- Trade history

#### Notification Service
- Email notifications
- Push notifications
- WebSocket notifications
- Notification templates

---

## Version History

- **0.1.0** - Initial release with Auth Service and basic infrastructure
