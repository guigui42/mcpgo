.PHONY: all build clean run test test-coverage docker-build docker-run help fmt lint

# Binary name
BINARY_NAME=mcpgo
# Build directory
BUILD_DIR=./bin

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOLINT=golangci-lint

# Docker parameters
DOCKER_IMAGE_NAME=mcpgo
DOCKER_TAG=latest

all: test build

# Build the application
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

# Clean build files
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Run the application
run:
	$(GOCMD) run ./cmd/$(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -cover -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	$(GOFMT) ./...

# Lint code
lint:
	$(GOLINT) run

# Update dependencies
deps:
	$(GOMOD) tidy

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .

# Run Docker container
docker-run:
	docker run -p 8080:8080 $(DOCKER_IMAGE_NAME):$(DOCKER_TAG)

# Display help
help:
	@echo "Available targets:"
	@echo "  all          - Run tests and build"
	@echo "  build        - Build the application binary"
	@echo "  clean        - Clean build files"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  deps         - Update dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  help         - Display this help message"

# Default target
.DEFAULT_GOAL := help