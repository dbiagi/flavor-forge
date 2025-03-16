# Makefile
GOEXEC = go
COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out
TEST_FILES = ./internal/...
TEST_FILES_INTEGRATION = ./tests/...
COMPOSE_FILE = ./docker-compose.yml
DOCKER_COMPOSE = docker compose -f "${COMPOSE_FILE}"

.PHONY: tests
tests:
	make test-unit
	make test-integration

.PHONY: test-unit
test-unit:
	@echo "Running tests..."
	${GOEXEC} test -v ${TEST_FILES}

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	${GOEXEC} test -bench=. -json -coverprofile="${COVERAGE_REPORT}" ${TEST_FILES} > "${TEST_REPORT}"

.PHONY: test-integration
test-integration:
	@echo "Running integration tests..."
	make infra-up
	${GOEXEC} test -v ${TEST_FILES_INTEGRATION}
	make infra-down

.PHONY: serve-dev
serve-dev:
	@echo "Starting server..."
	${GOEXEC} run cmd/main.go serve --env=dev

.PHONY: infra-up
infra-up:
	@echo "Starting infrastructure..."
	${DOCKER_COMPOSE} up -d

.PHONY: infra-down
infra-down:
	@echo "Stopping infrastructure..."
	${DOCKER_COMPOSE} down

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	${GOEXEC} mod vendor

.PHONY: build
build:
	@echo "Building..."
	${GOEXEC} build -o bin/app cmd/server/main.go
