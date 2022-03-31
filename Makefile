.PHONY: run
ifeq ($(MAKECMDGOALS),run)
include .env
export
endif
run:
	go run cmd/main.go

build	:
	go build -o coinconv cmd/main.go
