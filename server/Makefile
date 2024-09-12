NAME = tos
MODULE_NAME = github.com/MR5356/tos

BINARY_DIR = bin

BUILD_TIME=$(shell date +"%Y%m%d")
VERSION=$(shell cat ../VERSION 2>/dev/null || echo "unknown version")
COMMIT ?= $(shell git rev-parse --short HEAD)

GO_BUILD = CGO_ENABLED=0 go build -trimpath -ldflags '-X "$(MODULE_NAME)/constant.Version=$(VERSION)" -X "$(MODULE_NAME)/constant.Commit=$(COMMIT)" -X "$(MODULE_NAME)/constant.BuildTime=$(BUILD_TIME)" -w -s'

.DEFAULT_GOAL := help

.PHONY: help
help:  ## Show this help
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'

.PHONY: deps
deps:  ## Install dependencies
	@go mod download

.PHONY: test
test: deps  ## Test
	go test $(shell go list ./... | grep -v /docs) -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: version
version:  ## Show version
	@echo $(VERSION)

.PHONY: build
build:  ## Build
	@$(GO_BUILD) -o $(BINARY_DIR)/$(NAME)

.PHONY: linux-amd64
linux-amd64:  ## Build linux-amd64
	GOOS=linux GOARCH=amd64 $(GO_BUILD) -o $(BINARY_DIR)/$(NAME)-$@
	mv $(BINARY_DIR)/$(NAME)-$@ $(NAME)
	tar -zcvf $(NAME)-$@.tar.gz $(NAME)
	rm $(NAME) && mv $(NAME)-$@.tar.gz $(BINARY_DIR)

.PHONY: linux-arm64
linux-arm64:  ## Build linux-arm64
	GOOS=linux GOARCH=arm64 $(GO_BUILD) -o $(BINARY_DIR)/$(NAME)-$@
	mv $(BINARY_DIR)/$(NAME)-$@ $(NAME)
	tar -zcvf $(NAME)-$@.tar.gz $(NAME)
	rm $(NAME) && mv $(NAME)-$@.tar.gz $(BINARY_DIR)

.PHONY: windows-amd64
windows-amd64:  ## Build windows-amd64
	GOOS=windows GOARCH=amd64 $(GO_BUILD) -o $(BINARY_DIR)/$(NAME)-$@.exe
	zip -m -j $(NAME)-$@.zip $(BINARY_DIR)/$(NAME)-$@.exe && mv $(NAME)-$@.zip $(BINARY_DIR)

.PHONY: darwin-amd64
darwin-amd64:  ## Build darwin-amd64
	GOOS=darwin GOARCH=amd64 $(GO_BUILD) -o $(BINARY_DIR)/$(NAME)-$@
	mv $(BINARY_DIR)/$(NAME)-$@ $(NAME)
	tar -zcvf $(NAME)-$@.tar.gz $(NAME)
	rm $(NAME) && mv $(NAME)-$@.tar.gz $(BINARY_DIR)

.PHONY: darwin-arm64
darwin-arm64:  ## Build darwin-arm64
	GOOS=darwin GOARCH=arm64 $(GO_BUILD) -o $(BINARY_DIR)/$(NAME)-$@
	mv $(BINARY_DIR)/$(NAME)-$@ $(NAME)
	tar -zcvf $(NAME)-$@.tar.gz $(NAME)
	rm $(NAME) && mv $(NAME)-$@.tar.gz $(BINARY_DIR)

.PHONY: all
all: linux-amd64 linux-arm64 windows-amd64 darwin-amd64 darwin-arm64  ## Build all

.PHONY: clean
clean:  ## Clean
	@rm -rf $(BINARY_DIR)
	@rm -rf server/static/assets
	@rm -rf server/static/index.html
	@rm -rf server/static/favicon.ico
	@rm -rf coverage.out
