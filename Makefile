# Project name
APP_NAME = ainit

# Entry point
MAIN = ./main.go

# Output binary (optional)
BIN = bin/$(APP_NAME)

# Default target
.DEFAULT_GOAL := help

.PHONY: build run test tidy clean help

## Build the Go project
build:
	go build -o $(BIN) $(MAIN)

## Run the project with example arguments
run:
	go run $(MAIN) "create a RESTful API server" myapp

## Run tests
test:
	go test -v ./...

## Run go mod tidy
tidy:
	go mod tidy

## Clean build files
clean:
	rm -rf bin/

## Show help
help:
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo ""
	@echo "Targets:"
	@grep -E '^##' Makefile | sed -E 's/^## //;s/:.*//'
