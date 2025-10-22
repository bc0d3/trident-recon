.PHONY: build install clean test run fmt lint help

BINARY_NAME=trident-recon
INSTALL_PATH=$(shell go env GOPATH)/bin
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-s -w -X main.Version=$(VERSION) -X main.Commit=$(COMMIT) -X main.Date=$(DATE)"

build:
	@echo "Building $(BINARY_NAME) $(VERSION)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) .
	@echo "Build complete: ./$(BINARY_NAME)"

install:
	@echo "Installing $(BINARY_NAME)..."
	go install $(LDFLAGS)
	@echo "Installed to $(INSTALL_PATH)/$(BINARY_NAME)"

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	go clean
	@echo "Clean complete"

test:
	@echo "Running tests..."
	go test -v ./...

run:
	@echo "Running $(BINARY_NAME)..."
	go run $(LDFLAGS) . $(ARGS)

fmt:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Linting code..."
	golangci-lint run

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

help:
	@echo "Available targets:"
	@echo "  build    - Build the binary"
	@echo "  install  - Install the binary to GOPATH/bin"
	@echo "  clean    - Remove built binary"
	@echo "  test     - Run tests"
	@echo "  run      - Run the application (use ARGS='--help' for options)"
	@echo "  fmt      - Format code"
	@echo "  lint     - Lint code"
	@echo "  deps     - Download and tidy dependencies"
	@echo ""
	@echo "Example:"
	@echo "  make build"
	@echo "  make run ARGS='init'"
	@echo "  make run ARGS='run -u http://example.com'"
