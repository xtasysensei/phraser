##
# Project Title
#
# @file
# @version 0.1
BINARY=./bin/phraser

build:
	@echo "Building Phraser..."
	@go build -o $(BINARY)
run:
	@echo "Running Phraser..."
	@go run main.go
test:
	@go test -v ./...
fmt:
	@go fmt ./...
install:
	@go install
clean:
	@echo "Removing Phraser binary..."
	@rm -rf $(BINARY)

all: fmt test build

.PHONY: build run test fmt install clean all

# end
