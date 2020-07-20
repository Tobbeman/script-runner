.PHONY: help test build format init_hooks
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build using current system flags
	@go build -o script-runner cmd/script-runner/main.go

build_armv7l: ## Cross compile for arm
	@GOARCH=arm GOOS=linux GOARM=7 go build -o script-runner.armv7l cmd/script-runner/main.go

build_x86_64: ## Cross compile for amd64
	@GOARCH=amd64 GOOS=linux go build -o script-runner.x86_64 cmd/script-runner/main.go

build_all: build build_armv7l build_x86_64 ## Build all

test: ## Run tests
	@go test -race $$(go list ./... | grep -v script-runner/test)

format: ## Run go fmt
	@go fmt ./...

init_hooks: ## Will setup githooks for this git repository
	@ln -sf $$(pwd)/githooks/* $$(pwd)/.git/hooks
