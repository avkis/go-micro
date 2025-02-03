.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/api

## start: starts the application
start:
	@echo "Starting broker service"
	go run ./cmd/api

