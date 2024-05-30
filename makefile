# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=myapp
BINARY_UNIX=$(BINARY_NAME)_unix

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/main.go

# Run tests
test:
	$(GOTEST) -v ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Run the application
run: 
	$(GOCMD) run cmd/main.go

# Run the application with Docker Compose
run-docker:
	docker-compose up -d

# Stop Docker Compose
stop-docker:
	docker-compose down

.PHONY: build test clean run run-docker stop-docker
