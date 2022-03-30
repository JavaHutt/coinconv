.PHONY: run
ifeq ($(MAKECMDGOALS),run)
include .env
export
endif
run:
	go run main.go


build	:
	go build -o coinconv
