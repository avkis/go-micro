.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

## start: starts the application
start:
	@echo "Starting loggin service"
	go run ./cmd/api