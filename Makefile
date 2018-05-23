.PHONY: run help test
.DEFAULT_GOAL := help

help: ## List all the command helps.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install the project.
	@cp .env.dist .env
	@dep ensure

run: ## Run the server.
	@go run cmd/server/main.go

test: ## Test the project
	@cd cmd/server && go test -race -v

bench: ## Benchmark of the project
	@cd cmd/server && go test -bench=.


