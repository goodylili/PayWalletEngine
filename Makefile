# Makefile for Go project using Docker and Docker Compose

# Project name
PROJECT_NAME := mygoproject

# Docker Compose file
DOCKER_COMPOSE_FILE := docker-compose.yml

# Docker Compose command
DOCKER_COMPOSE := docker-compose -f $(DOCKER_COMPOSE_FILE)

# Go source files
GO_SRC := $(shell find . -type f -name '*.go')

# Default target
.DEFAULT_GOAL := help

# Targets
.PHONY: build run test lint clean help

# Build Docker image
build:
	@echo "Building Docker image..."
	$(DOCKER_COMPOSE) build

# Run Docker container
run:
	@echo "Starting Docker container..."
	$(DOCKER_COMPOSE) up -d

# Run tests
test: $(GO_SRC)
	@echo "Running tests..."
	go test ./...

# Run lint
lint: $(GO_SRC)
	@echo "Running lint..."
	golangci-lint run ./...

# Stop and remove Docker containers
clean:
	@echo "Stopping and removing Docker containers..."
	$(DOCKER_COMPOSE) down

# Display help
help:
	@echo "Usage: make [TARGET]"
	@echo "Targets:"
	@echo "  build   Build Docker image"
	@echo "  run     Start Docker container"
	@echo "  test    Run tests"
	@echo "  lint    Run lint"
	@echo "  clean   Stop and remove Docker containers"

