all: run

run:
	@go run cmd/httpsink/main.go

.PHONY: all run  
