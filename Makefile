PROJECT=~/go/src/gitlab.com/Tobbeman/script-runner
.PHONY: help
.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build using default tags
	@go build $(PROJECT)/cmd/script-runner/

arm: ## Cross compile for arm
	@GOARCH=arm GOOS=linux go build $(PROJECT)/cmd/script-runner/ 

test: ## Run tests
	@go test -cover -v ./...
