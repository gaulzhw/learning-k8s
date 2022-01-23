# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL:=help

GOPATH  := $(shell go env GOPATH)
GOARCH  := $(shell go env GOARCH)
GOOS    := $(shell go env GOOS)
GOPROXY := $(shell go env GOPROXY)
ifeq ($(GOPROXY),)
GOPROXY := https://proxy.golang.org
endif
export GOPROXY

export GO111MODULE=on

# Directories.
ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/bin
TEMPLATES_DIR := templates

BIN_DIR := bin
GOLANGCI_LINT := $(TOOLS_BIN_DIR)/golangci-lint
CONTROLLER_GEN := $(TOOLS_BIN_DIR)/controller-gen
CONVERSION_GEN := $(TOOLS_BIN_DIR)/conversion-gen
GO_BINDATA := $(TOOLS_BIN_DIR)/go-bindata
GINKGO := $(TOOLS_BIN_DIR)/ginkgo
APISERVER_BOOT := $(TOOLS_BIN_DIR)/apiserver-boot
EXP_DIR := exp

# Allow overriding manifest generation destination directory
MANIFEST_ROOT ?= deploy
CONTROLLER_DEPLOYMENT_FILE ?= $(MANIFEST_ROOT)/base/deployment.yaml
CONTROLLER_RBAC_ROOT ?= $(MANIFEST_ROOT)/base

# Define Docker related variables.
HARBOR = harbor.test.shopeemobile.com
ifeq ($(ENV),live)
HARBOR := harbor.shopeemobile.com
endif

# core
CONTROLLER_IMAGE_NAME ?= devops-sz/haproxy-controller
CONTROLLER_IMG ?= $(HARBOR)/$(CONTROLLER_IMAGE_NAME)
HAPROXY_IMAGE_NAME ?= devops-sz/kube-haproxy
HAPROXY_IMG ?= $(HARBOR)/$(HAPROXY_IMAGE_NAME)

TAG ?= $(shell git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --tags || git  rev-parse HEAD)

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9A-Za-z_-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Binaries
## --------------------------------------

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod # Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(GO_BINDATA): $(TOOLS_DIR)/go.mod
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/go-bindata github.com/go-bindata/go-bindata/go-bindata

$(GINKGO): $(TOOLS_DIR)/go.mod
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/ginkgo github.com/onsi/ginkgo/ginkgo

$(APISERVER_BOOT): $(TOOLS_DIR)/go.mod # Build apiserver-boot from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(APISERVER_BOOT)/apiserver-boot sigs.k8s.io/apiserver-builder-alpha/cmd/apiserver-boot

apiserver: $(APISERVER_BOOT)

## --------------------------------------
## Linting
## --------------------------------------

test: $(CONTROLLER_GEN)

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint codebase
	$(GOLANGCI_LINT) run -v

## --------------------------------------
## Testing
## --------------------------------------

.PHONY: test-e2e
test-e2e: $(GINKGO) ## Run the e2e tests
	$(GINKGO) -v -tags=e2e -r ./...

## --------------------------------------
## Modules
## --------------------------------------

.PHONY: modules
modules: ## Runs go mod to ensure modules are up to date
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

## --------------------------------------
## Docker
## --------------------------------------

.PHONY: docker-build
docker-build: ## Build the docker image
	docker build -f build/controller/Dockerfile -t $(CONTROLLER_IMG):$(TAG) .

.PHONY: docker-push
docker-push: ## Push the docker image
	docker push $(CONTROLLER_IMG):$(TAG)

## --------------------------------------
## GENERATE/MANIFEST
## --------------------------------------

.PHONY: update-deployment
update-deployment: ## update controller image version in deployment
	sed -i'' -e 's@image: ${CONTROLLER_IMG}:.*@image: '"${CONTROLLER_IMG}:$(TAG)"'@' $(CONTROLLER_DEPLOYMENT_FILE)
	rm $(CONTROLLER_DEPLOYMENT_FILE)-e

.PHONY: generate-rbac
generate-rbac: $(CONTROLLER_GEN) ## Generate controller manifests e.g. RBAC.
	$(CONTROLLER_GEN) paths=./pkg/controllers/... output:rbac:dir=$(CONTROLLER_RBAC_ROOT) rbac:roleName=haproxy-controller

## --------------------------------------
## RELEASE
## --------------------------------------

.PHONY: release
release: docker-build docker-push generate-rbac update-deployment ## release for controller
