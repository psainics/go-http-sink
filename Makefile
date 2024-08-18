BINARY_NAME := "httpsink"
APP_NAME := "httpsink"

all: run

build:
	@go build -o dist/$(BINARY_NAME) cmd/$(APP_NAME)/main.go

run:
	@go run cmd/$(APP_NAME)/main.go

.PHONY: all build run  
