.PHONY: build
build:
	go build -v ./cmd/rest-api

.DEFAULT_GOAL := build