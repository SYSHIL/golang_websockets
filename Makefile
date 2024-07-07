# Variables
BINARY_NAME = talk_server
CMD_DIR = ./cmd/server

# Default target
all: build

# Build the binary
build:
	@echo "Building the server..."
	go build -o $(BINARY_NAME) $(CMD_DIR)/main.go

# Run the server (build and run the binary)
run: build
	@echo "Running the server..."
	./$(BINARY_NAME)

# Run the server without building the binary
run-direct:
	@echo "Running the server directly..."
	go run $(CMD_DIR)/main.go

# Clean up binaries
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Format the code
fmt:
	@echo "Formatting the code..."
	go fmt ./...

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# PHONY targets to avoid conflicts with files of the same name
.PHONY: all build run run-direct clean fmt test deps
