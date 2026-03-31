BINARY     := myapp
MODULE     := github.com/xuxiaohu/myapp
CMD        := ./cmd/myapp

VERSION    ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS    := -ldflags "-X $(MODULE)/pkg/version.Version=$(VERSION)"

GO         := go
GOFLAGS    :=
GOTEST     := $(GO) test $(GOFLAGS) -race

BUILD_DIR  := bin

.PHONY: all build test lint fmt vet tidy clean help

all: lint test build

## build: compile the binary to bin/
build:
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) $(CMD)

## run: build and run the binary
run: build
	./$(BUILD_DIR)/$(BINARY)

## test: run all tests with race detector
test:
	$(GOTEST) ./...

## test-cover: run tests and show coverage
test-cover:
	$(GOTEST) -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## lint: run golangci-lint
lint:
	golangci-lint run ./...

## fmt: format all Go files
fmt:
	$(GO) fmt ./...
	goimports -w .

## vet: run go vet
vet:
	$(GO) vet ./...

## tidy: tidy and verify go.mod
tidy:
	$(GO) mod tidy
	$(GO) mod verify

## clean: remove build artifacts
clean:
	@rm -rf $(BUILD_DIR) coverage.out coverage.html

## help: print this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'
