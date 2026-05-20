# OperaTree Makefile

# -------------------------------------------------------------------
# Variables
# -------------------------------------------------------------------

MODULE    := github.com/hanymamdouh82/operatree
BINARY    := operatree
CMD_DIR   := .
BUILD_DIR := ./build

VERSION    := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT     := $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -ldflags "\
	-X main.Version=$(VERSION) \
	-X main.Commit=$(COMMIT) \
	-X main.BuildDate=$(BUILD_DATE) \
	-s -w"

GO      := go
GOFLAGS :=

# -------------------------------------------------------------------
# Install directory — OS agnostic, no Go required
# -------------------------------------------------------------------

UNAME := $(shell uname -s 2>/dev/null || echo "Windows")

ifeq ($(UNAME), Linux)
	INSTALL_DIR := /usr/local/bin
else ifeq ($(UNAME), Darwin)
	INSTALL_DIR := /usr/local/bin
else
	# Windows — works with Git Bash, WSL, or PowerShell
	INSTALL_DIR := $(USERPROFILE)\AppData\Local\Microsoft\WindowsApps
endif

# -------------------------------------------------------------------
# Default target
# -------------------------------------------------------------------

.DEFAULT_GOAL := help

# -------------------------------------------------------------------
# Help
# -------------------------------------------------------------------

.PHONY: help
help: ## Show this help message
	@echo ""
	@echo "  OperaTree — build targets"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "  Install directory: $(INSTALL_DIR)"
	@echo ""

# -------------------------------------------------------------------
# Build
# -------------------------------------------------------------------

.PHONY: build
build: ## Build binary to ./build/
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY) $(CMD_DIR)
	chmod +x $(BUILD_DIR)/$(BINARY)
	@echo "Built: $(BUILD_DIR)/$(BINARY) ($(VERSION))"

.PHONY: build-all
build-all: ## Cross-compile for Linux, macOS, Windows
	@mkdir -p $(BUILD_DIR)
	GOOS=linux   GOARCH=amd64  $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-amd64   $(CMD_DIR)
	GOOS=linux   GOARCH=arm64  $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-linux-arm64   $(CMD_DIR)
	GOOS=darwin  GOARCH=amd64  $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-amd64  $(CMD_DIR)
	GOOS=darwin  GOARCH=arm64  $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-darwin-arm64  $(CMD_DIR)
	GOOS=windows GOARCH=amd64  $(GO) build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY)-windows-amd64.exe $(CMD_DIR)
	@echo "Cross-compiled binaries in $(BUILD_DIR)/"

# -------------------------------------------------------------------
# Install / Uninstall
# -------------------------------------------------------------------

.PHONY: install
install: build ## Install binary to system path (may require sudo on Linux/macOS)
	@echo "Installing to $(INSTALL_DIR)..."
ifeq ($(UNAME), Windows)
	cp $(BUILD_DIR)/$(BINARY).exe "$(INSTALL_DIR)/$(BINARY).exe"
else ifeq ($(shell test -w $(INSTALL_DIR) && echo yes), yes)
	cp $(BUILD_DIR)/$(BINARY) $(INSTALL_DIR)/$(BINARY)
else
	sudo cp $(BUILD_DIR)/$(BINARY) $(INSTALL_DIR)/$(BINARY)
endif
	@echo "Installed: $(INSTALL_DIR)/$(BINARY)"

.PHONY: uninstall
uninstall: ## Remove binary from system path (may require sudo on Linux/macOS)
	@echo "Removing from $(INSTALL_DIR)..."
ifeq ($(UNAME), Windows)
	rm -f "$(INSTALL_DIR)/$(BINARY).exe"
else ifeq ($(shell test -w $(INSTALL_DIR) && echo yes), yes)
	rm -f $(INSTALL_DIR)/$(BINARY)
else
	sudo rm -f $(INSTALL_DIR)/$(BINARY)
endif
	@echo "Uninstalled: $(INSTALL_DIR)/$(BINARY)"

# -------------------------------------------------------------------
# Run
# -------------------------------------------------------------------

.PHONY: run
run: build ## Build and run operatree
	$(BUILD_DIR)/$(BINARY)

# -------------------------------------------------------------------
# Test
# -------------------------------------------------------------------

.PHONY: test
test: ## Run all tests
	$(GO) test ./... -v

.PHONY: test-short
test-short: ## Run tests without long-running cases
	$(GO) test ./... -short

.PHONY: coverage
coverage: ## Run tests with coverage report
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# -------------------------------------------------------------------
# Code quality
# -------------------------------------------------------------------

.PHONY: lint
lint: ## Run golangci-lint
	@which golangci-lint > /dev/null || (echo "golangci-lint not found. Install: https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run ./...

.PHONY: fmt
fmt: ## Format all Go source files
	$(GO) fmt ./...

.PHONY: vet
vet: ## Run go vet
	$(GO) vet ./...

.PHONY: tidy
tidy: ## Tidy go.mod and go.sum
	$(GO) mod tidy

.PHONY: check
check: fmt vet lint test ## Run all checks (fmt, vet, lint, test)

# -------------------------------------------------------------------
# Demo
# -------------------------------------------------------------------

.PHONY: demo
demo: ## Record VHS demo tape
	@which vhs > /dev/null || (echo "vhs not found. Install: https://github.com/charmbracelet/vhs" && exit 1)
	vhs demo/demo.tape

# -------------------------------------------------------------------
# Clean
# -------------------------------------------------------------------

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	@echo "Cleaned"

# -------------------------------------------------------------------
# Version
# -------------------------------------------------------------------

.PHONY: version
version: ## Print current version info
	@echo "Version:    $(VERSION)"
	@echo "Commit:     $(COMMIT)"
	@echo "Build date: $(BUILD_DATE)"
