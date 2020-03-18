BINARY=goddd

DOCKER_IMAGE_NAME=franklincarrero/bookstore_users-api

.DEFAULT_GOAL := help

check: test lint vet ## Runs all tests

build:
	go build -v ./...
test: ## Run the unit tests
	go test -cover -race -v ./...

lint: ## Lint all files
	go list ./...

vet: ## Run the vet tool
	go vet $(shell go list ./... | grep -v /vendor/)

clean: ## Clean up build artifacts
	go clean -cache -modcache -i -r

docker-build: ## Build Docker image
	docker build --no-cache -t ${DOCKER_IMAGE_NAME} .

docker-build-no-cache: ## Build Docker image
	docker build  -t ${DOCKER_IMAGE_NAME} .

docker-push: ## Push Docker image to registry
	docker push ${DOCKER_IMAGE_NAME}

help: ## Display this help message
	@cat $(MAKEFILE_LIST) | grep -e "^[a-zA-Z_\-]*: *.*## *" | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.SILENT: build test lint vet clean docker-build docker-push help

.PHONY: all clean test