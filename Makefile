.PHONY: build
build:
	go build -v ./

.PHONY: dev
dev:
	go run main.go

.DEFAULT_GOAL := build
