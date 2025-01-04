# Makefile
GOEXEC ?= go

.PHONY: test
test:
	@echo "Running tests..."
	${GOEXEC} test -v ./...