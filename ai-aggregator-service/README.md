# AI Aggregator Service

A comprehensive microservices-based AI aggregator service that provides a unified API for accessing multiple AI providers (OpenAI, Anthropic, Google AI, Cohere) with features like load balancing, rate limiting, caching, and analytics.

## Architecture Overview

This service implements a microservices architecture with the following components:

- **API Gateway**: Entry point for all client requests
- **Auth Service**: Handles authentication and authorization
- **Unified API**: Core API service for AI requests
- **Provider Router**: Routes requests to appropriate AI providers
- **User Management**: User account and subscription management
- **Billing Service**: Payment processing and usage tracking
- **Analytics Service**: Usage analytics and reporting
- **Monitoring Service**: Health checks and metrics

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- PostgreSQL 14+
- Redis 6+
- NATS Server

### Environment Setup

1. **Clone and setup environment:**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

2. **Start development environment:**
   ```bash
   make dev-up
   ```

3. **Run individual services:**
   ```bash
   make run-gateway      # API Gateway
   make run-auth         # Auth Service
   make run-unified-api  # Unified API
   make run-provider-router  # Provider Router
   ```

### Environment Variables

All configuration is handled through environment variables. Copy `.env.example` to `.env` and configure:

#### Server Configuration
- `AGG_SERVER_HOST`: Server host (default: 0.0.0.0)
- `AGG_SERVER_PORT`: Server port (default: 8080)

#### Database Configuration
- `AGG_DATABASE_HOST`: PostgreSQL host (default: localhost)
- `AGG_DATABASE_PORT`: PostgreSQL port (default: 5432)
- `AGG_DATABASE_USER`: PostgreSQL user (default: postgres)
- `AGG_DATABASE_PASSWORD`: PostgreSQL password (default: postgres)
- `AGG_DATABASE_DBNAME`: Database name (default: ai_aggregator)

#### Redis Configuration
- `AGG_REDIS_HOST`: Redis host (default: localhost)
- `AGG_REDIS_PORT`: Redis port (default: 6379)
- `AGG_REDIS_PASSWORD`: Redis password (optional)
- `AGG_REDIS_DB`: Redis database (default: 0)

#### NATS Configuration
- `AGG_NATS_URL`: NATS server URL (default: nats://localhost:4222)

#### Authentication
- `AGG_AUTH_JWT_SECRET`: JWT secret key
- `AGG_AUTH_JWT_EXPIRATION`: JWT expiration duration (default: 24h)

#### AI Provider API Keys
- `OPENAI_API_KEY`: OpenAI API key
- `ANTHROPIC_API_KEY`: Anthropic API key
- `GOOGLE_AI_API_KEY`: Google AI API key
- `COHERE_API_KEY`: Cohere API key

## Development

### Available Make Commands

```bash
make help              # Display help
make build             # Build all services
make test              # Run all tests
make dev-up            # Start development environment
make dev-down          # Stop development environment
make docker-build      # Build Docker images
make docker-run        # Run with Docker Compose
make lint              # Run linter
make fmt               # Format code
make db-migrate        # Run database migrations
make setup             # Setup development environment
```

### Project Structure

```
ai-aggregator-service/
├── cmd/                    # Service entry points
│   ├── api-gateway/
│   ├── auth-service/
│   ├── unified-api/
│   └── provider-router/
├── internal/               # Internal packages
│   ├── config/            # Configuration management
│   ├── logger/            # Logging utilities
│   ├── auth/              # Authentication
│   ├── cache/             # Caching layer
│   ├── database/          # Database operations
│   ├── models/            # Data models
│   ├── providers/         # AI provider integrations
│   ├── middleware/        # HTTP middleware
│   └── utils/             # Utility functions
├── api/                    # API definitions
├── migrations/             # Database migrations
├── scripts/                # Utility scripts
├── deployments/            # Deployment configurations
├── docker-compose.yml      # Docker Compose for production
├── docker-compose.dev.yml  # Docker Compose for development
├── Dockerfile             # Multi-stage Dockerfile
├── Makefile               # Development tasks
└── README.md              # This file
```

### API Endpoints

#### Authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout
- `POST /api/v1/auth/refresh` - Refresh token

#### AI Operations
- `POST /api/v1/chat/completions` - Chat completions
- `POST /api/v1/completions` - Text completions
- `POST /api/v1/embeddings` - Create embeddings
- `GET /api/v1/models` - List available models

#### User Management
- `GET /api/v1/users/profile` - Get user profile
- `PUT /api/v1/users/profile` - Update user profile
- `GET /api/v1/users/usage` - Get usage statistics

### Testing

```bash
# Run all tests
make test

# Run unit tests only
make test-unit

# Run integration tests
make test-integration

# Run tests with coverage
make test-coverage
```

### Database

The service uses PostgreSQL for persistent storage and Redis for caching.

#### Migrations
```bash
# Run migrations
make db-migrate

# Seed database with initial data
make db-seed
```

### Monitoring

#### Metrics
- Prometheus metrics available at `/metrics`
- Health check endpoint at `/health`

#### Logging
- Structured logging using Go's slog package
- Configurable log level and format (JSON/text)

### Deployment

#### Docker Deployment
```bash
# Build and run with Docker Compose
make docker-build
make docker-run

# Production deployment
docker-compose -f docker-compose.yml up -d
```

#### Environment-specific Configurations

- **Development**: Uses `docker-compose.dev.yml` with hot reload
- **Production**: Uses `docker-compose.yml` with optimized settings
- **Testing**: Uses separate test containers

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.