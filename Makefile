SHELL := /bin/bash

PHONY: deps
deps: #Sync all dependencies
	go mod tidy
	go mod vendor

PHONY: build
build: deps
	go build -o bin/main

PHONY: serve
serve: build
	bin/main

PHONY:lint
lint:
	golangci-lint run ./...