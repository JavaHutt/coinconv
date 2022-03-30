.PHONY: run
ifeq ($(MAKECMDGOALS),run)
include .env
export
endif
run:
	go run cmd/main.go


.PHONY: build
ifeq ($(MAKECMDGOALS),build)
include .env
export
endif
build	:
	go build -o coinconv cmd/main.go
