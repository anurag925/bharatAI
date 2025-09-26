.PHONY: help build run test clean docker-build docker-run deps lint fmt vet tidy docs

# Default target
help: ## Display this help message
	@echo "AI Aggregator Service"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Build targets
build: ## Build all services
	@echo "Building all services..."
	@for dir in cmd/*/; do \
		service=$$(basename $$dir); \
		echo "Building $$service..."; \
		go build -o bin/$$service ./cmd/$$service; \
	done

build-api-gateway: ## Build API Gateway service
	go build -o bin/api-gateway ./cmd/api-gateway

build-auth-service: ## Build Auth service
	go build -o bin/auth-service ./cmd/auth-service

build-unified-api: ## Build Unified API service
	go build -o bin/unified-api ./cmd/unified-api

build-provider-router: ## Build Provider Router service
	go build -o bin/provider-router ./cmd/provider-router

# Development targets
run: ## Run the API Gateway service
	go run ./cmd/api-gateway

run-api-gateway: ## Run API Gateway service
	go run ./cmd/api-gateway

run-auth-service: ## Run Auth service
	go run ./cmd/auth-service

run-unified-api: ## Run Unified API service
	go run ./cmd/unified-api

run-provider-router: ## Run Provider Router service
	go run ./cmd/provider-router

# Testing targets
test: ## Run all tests
	go test -v ./...

test-unit: ## Run unit tests
	go test -v ./tests/unit/...

test-integration: ## Run integration tests
	go test -v ./tests/integration/...

test-e2e: ## Run end-to-end tests
	go test -v ./tests/e2e/...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Database targets
db-migrate: ## Run database migrations
	go run ./tools/migrations/migrate.go

db-seed: ## Seed database with initial data
	go run ./tools/seeder/seeder.go

# Docker targets
docker-build: ## Build Docker images for all services
	docker compose build

docker-run: ## Run services with Docker Compose
	docker compose up

docker-down: ## Stop Docker services
	docker compose down

docker-logs: ## Show Docker logs
	docker compose logs -f

# Development environment
dev-up: ## Start development environment
	docker compose -f docker-compose.yml up -d

dev-down: ## Stop development environment
	docker compose -f docker-compose.yml down

dev-logs: ## Show development environment logs
	docker compose -f docker-compose.yml logs -f

# Code quality targets
lint: ## Run linter
	golangci-lint run

fmt: ## Format code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

tidy: ## Tidy go modules
	go mod tidy
	go mod verify

# Utility targets
clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf coverage.out coverage.html

deps: ## Install dependencies
	go mod download
	go mod tidy

install-tools: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# Documentation
docs: ## Generate API documentation
	swag init -g cmd/api/main.go
	swag fmt

mocks:
	mockery --all --with-expecter=true

# Monitoring
prometheus: ## Start Prometheus for monitoring
	docker compose -f docker compose.monitoring.yml up -d prometheus grafana

# All-in-one development setup
setup: ## Setup development environment
	@echo "Setting up development environment..."
	@make deps
	@make install-tools
	@cp .env.example .env
	@echo "Development environment setup complete!"
	@echo "Please edit .env file with your configuration"

# Health check
health: ## Check service health
	curl -f http://localhost:8080/health || exit 1