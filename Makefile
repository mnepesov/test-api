NAME ?= astronomy
GO_PACKAGE = back

# ==============================================================================
# Build

BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ${GO_PACKAGE}/cmd/${NAME}

.PHONY: build
build: clean
	@echo "> Build ${NAME}"
	@${BUILD_CMD}

.PHONY: clean
clean:
	@echo "> Cleaning binaries for ${NAME}"
	@-rm -rf bin/${NAME}

# ==============================================================================
# tests

.PHONY: tests
tests:
	@go test -v -race ./...

# ==============================================================================
# app

.PHONY: docker_env_start_binding
docker_env_start_binding: docker_env_stop
	@echo "> Run environment"
	@read -p "Pull [y/n]: " ans; \
	if [ $$ans = y ]; then \
		docker-compose pull; \
	fi
	@docker-compose up -d postgresql migrate

.PHONY: docker_env_start_full
docker_env_start_full: docker_env_stop
	@echo "> Run environment"
	@read -p "Pull [y/n]: " ans; \
	if [ $$ans = y ]; then \
		docker-compose pull; \
	fi
	@read -p "Build service [y/n]: " ans; \
	if [ $$ans = y ]; then \
		docker build -t astronomy .; \
	fi
	@docker-compose up -d

.PHONY: docker_env_stop
docker_env_stop:
	@echo "> Stop docker environment"
	@docker-compose down

.PHONY: docker_build
docker_build:
	docker build -t astronomy .

# ==============================================================================
# Linters

run-linter:
	echo "Starting linters"
	golangci-lint run ./...

# ==============================================================================
# Swagger

swagger:
	echo "Starting swagger generating"
	swag init -g **/**/*.go