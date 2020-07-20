.PHONY: help test build arm format init_hooks
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build using default tags
	@go build -o script-runner cmd/script-runner/main.go

arm: ## Cross compile for arm
	@GOARCH=arm GOOS=linux go build -o script-runner.arm cmd/script-runner/main.go

test: ## Run tests
	@go test -race $$(go list ./... | grep -v script-runner/test)

format: ## Run go fmt
	@go fmt ./...

init_hooks: ## Will setup githooks for this git repository
	@ln -sf $$(pwd)/githooks/* $$(pwd)/.git/hooks
