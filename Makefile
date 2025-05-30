BINARY_NAME = go-ssr

CLEAN_TARGETS :=
CLEAN_TARGETS += '$(BINARY_NAME)'

.PHONY: clean
clean: ## Clean up all build artifacts
	rm -rf $(CLEAN_TARGETS)

.PHONY: build
build: clean tidy format ## Build the project
	go build -o $(BINARY_NAME) ./cmd

.PHONY: format
format: ## Format the code
	go fmt ./...

.PHONY: tidy
tidy: ## Tidy up the go modules
	go mod tidy
