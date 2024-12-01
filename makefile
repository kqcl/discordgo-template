# Makefile for building the Go project

BINARY_NAME := discord-bot 
BUILD_DIR := ./bin

build_and_run: build run
	@echo "Build and run completed!"

build:
	@echo "Building the project..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BUILD_DIR)/$(BINARY_NAME)

run:
	@echo "Running the program..."
	./bin/$(BINARY_NAME)

setup:
	@echo "Installing dependencies..."
	go mod tidy
