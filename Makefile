.PHONY: run help test
.DEFAULT_GOAL := help

help: ## List all the command helps.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install the project.
	@cp .env.dist .env
	@dep ensure

run: ## Run the server.
	@go run cmd/server/main.go

test: ## Test the server.
	@cd cmd/server && go test -race -v

test-cli: ## Test the cli.
	@cd cmd/cli && go test -race -v

bench: ## Benchmark of the project.
	@cd cmd/server && go test -bench=.

build: ## Build the docker image.
	@cd ml/fixture/ && docker build --no-cache -t api-rest-ml .

start-ml: ## Run the docker image
	@cd ml/fixture/ && docker run --rm -v $(PWD)/ml/fixture/:/src/app api-rest-ml

save: ## Save the docker image in a tar file
	@cd ml/fixture/ && docker save --output=api-rest-ml.tar api-rest-ml

protoc: ## Generate protobuffer
	@cd cmd/grpc/ && protoc *.proto --go_out=plugins=grpc:.
