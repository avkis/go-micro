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
	@echo "Starting mailer service"
	go run ./cmd/api

