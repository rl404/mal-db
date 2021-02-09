# Base Go commands.
GO_CMD   := go
GO_FMT   := $(GO_CMD) fmt
GO_CLEAN := $(GO_CMD) clean
GO_BUILD := $(GO_CMD) build -mod vendor
GO_RUN   := $(GO_CMD) run -mod vendor

# Base swagger commands.
SWAG     := swag
SWAG_GEN := ${SWAG} init

# Project executable file, and its binary.
CMD_PATH    := ./cmd/mal
BINARY_NAME := mal

# Default makefile target.
.DEFAULT_GOAL := run

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@golint `go list ./... | grep -v /vendor/`

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

# Generate swagger docs.
.PHONY: swagger
swagger:
	@${SWAG_GEN} -g cmd/mal/main.go -o ./api

# Build the project executable binary.
.PHONY: build
build: clean fmt
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v .

# Prepare database.
.PHONY: install
install: build swagger
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) install

# Build and run the binary.
.PHONY: run
run: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) server

#
.PHONY: worker
worker: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) worker

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD   := docker-compose
COMPOSE_PATH  := deployment/docker-compose.yml
COMPOSE_PATH2 := deployment/install.yml

# Prepare database.
.PHONY: docker-install
docker-install: clean fmt
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH2) up
	@$(DOCKER_IMAGE) prune -f --filter label=stage=mal_db_builder

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build: clean fmt
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=mal_db_builder

# Start built docker containers.
.PHONY: docker-up
docker-up:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) logs --follow --tail 20

# Build and start docker container for the project.
.PHONY: docker
docker: docker-build docker-up

# Stop docker container.
.PHONY: docker-stop
docker-stop:
	@$(COMPOSE_CMD) -f $(COMPOSE_PATH) stop