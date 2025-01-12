# Makefile
GOEXEC = go
COVERAGE_REPORT = coverage.out
TEST_REPORT = report.out

.PHONY: test
test:
	@echo "Running tests..."
	${GOEXEC} test -v ./...

.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	${GOEXEC} test -json -coverprofile="${COVERAGE_REPORT}" ./... > "${TEST_REPORT}"