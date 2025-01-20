# Makefile
GOEXEC = go
COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out
COMPOSE_FILE = ./docker-compose.yml
DOCKER_COMPOSE = docker compose -f "${COMPOSE_FILE}"

.PHONY: test
test:
	@echo "Running tests..."
	${GOEXEC} test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	${GOEXEC} test -bench=. -json -coverprofile="${COVERAGE_REPORT}" ./... > "${TEST_REPORT}"

.PHONY: test-integration
test-integration:
	@echo "Running integration tests..."
	make infra-up
	${GOEXEC} test -v ./...
	make infra-down

.PHONY: serve
serve:
	@echo "Starting server..."
	${GOEXEC} run cmd/server/main.go

.PHONY: infra-up
infra-up:
	@echo "Starting infrastructure..."
	${DOCKER_COMPOSE} up -d

.PHONY: infra-down
infra-down:
	@echo "Stopping infrastructure..."
	${DOCKER_COMPOSE} down