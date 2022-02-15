# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL := help

#
# Directories.
#
# Full directory of where the Makefile resides
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR := bin
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/$(BIN_DIR)
GEN_OUTPUT_BASE := --output-base=$(ROOT_DIR)

#
# Binaries.
#
# Note: Need to use abspath so we can invoke these from subdirectories
KUSTOMIZE := $(abspath $(TOOLS_BIN_DIR)/kustomize)
CONTROLLER_GEN := $(abspath $(TOOLS_BIN_DIR)/controller-gen)
GOLANGCI_LINT := $(abspath $(TOOLS_BIN_DIR)/golangci-lint)
CONVERSION_GEN := $(abspath $(TOOLS_BIN_DIR)/conversion-gen)
CONVERSION_VERIFIER := $(abspath $(TOOLS_BIN_DIR)/conversion-verifier)

#
# Container.
#
GO_CONTAINER_VERSION ?= 1.17-alpine
GO_CONTAINER_IMAGE ?= docker.io/library/golang:$(GO_CONTAINER_VERSION)

#
# Image.
#
# Define Docker related variables. Releases should modify and double check these vars.
REGISTRY ?= harbor.test.sample.io
IMAGE_NAME ?= sample/sample-operator
CONTROLLER_IMG ?= $(REGISTRY)/$(IMAGE_NAME)

# Set image tag from branch or tag
TAG ?= $(shell git rev-parse --abbrev-ref HEAD | grep -v HEAD || git describe --tags || git  rev-parse HEAD)

# Allow overriding the imagePullPolicy
PULL_POLICY ?= Always

help: # Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[0-9A-Za-z_-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ generate:

.PHONY: generate
generate: ## Run all generate-manifests, generate-go-deepcopy and generate-go-conversions targets
	$(MAKE) generate-go-deepcopy generate-manifests

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc. for core
	$(CONTROLLER_GEN) \
		paths=./pkg/apis/... \
		crd:crdVersions=v1 \
		output:crd:dir=./config/crd/bases

.PHONY: generate-go-deepcopy
generate-go-deepcopy: $(DEEPCOPY_GEN) ## Generate deepcopy code
	$(DEEPCOPY_GEN) \
		--input-dirs=./pkg/apis \
		--output-dirs=./pkg/apis \
		--output-file-base=zz_generated.deepcopy $(GEN_OUTPUT_BASE)

## --------------------------------------
## Lint / Verify
## --------------------------------------

##@ lint and verify:

.PHONY: modules
modules: ## Run go mod tidy to ensure modules are up to date
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

.PHONY: lint
lint: $(GOLANGCI_LINT) ## Lint the codebase
	$(GOLANGCI_LINT) run -v

.PHONY: lint-fix
lint-fix: $(GOLANGCI_LINT) ## Lint the codebase and run auto-fixers if supported by the linter
	GOLANGCI_LINT_EXTRA_ARGS=--fix $(MAKE) lint

.PHONY: verify-modules
verify-modules: modules  ## Verify go modules are up to date
	@if (find . -name 'go.mod' | xargs -n1 grep -q -i 'k8s.io/client-go.*+incompatible'); then \
		find . -name "go.mod" -exec grep -i 'k8s.io/client-go.*+incompatible' {} \; -print; \
		echo "go module contains an incompatible client-go version"; exit 1; \
	fi

.PHONY: verify-gen
verify-gen: generate  ## Verfiy go generated files are up to date
	@if !(git diff --quiet HEAD); then \
		git diff; \
		echo "generated files are out of date, run make generate"; exit 1; \
	fi

## --------------------------------------
## Binaries
## --------------------------------------

##@ build:

.PHONY: docker-build
docker-build: ## Build the docker image for core controller manager
	docker build --build-arg builder_image=$(GO_CONTAINER_IMAGE) . -t $(CONTROLLER_IMG):$(TAG)
	$(MAKE) set-manifest-image MANIFEST_IMG=$(CONTROLLER_IMG) MANIFEST_TAG=$(TAG) TARGET_RESOURCE="./config/default/manager_image_patch.yaml"
	$(MAKE) set-manifest-pull-policy TARGET_RESOURCE="./config/default/manager_pull_policy.yaml"

.PHONY: docker-push
docker-push: ## Push the docker images
	docker push $(CONTROLLER_IMG):$(TAG)

## --------------------------------------
## Release
## --------------------------------------

##@ release:

.PHONY: release
release: ## Build and push container images using the latest git tag for the commit
	# Build the manifests
	$(MAKE) release-manifests

.PHONY: release-manifests
release-manifests: $(KUSTOMIZE) ## Build the manifests to publish with a release
	$(KUSTOMIZE) build config/default > deploy/components.yaml

## --------------------------------------
## Manifest
## --------------------------------------

##@ manifest:

.PHONY: set-manifest-pull-policy
set-manifest-pull-policy: ## Setup manifest pull policy
	$(info Updating kustomize pull policy file for manager resources)
	sed -i'' -e 's@imagePullPolicy: .*@imagePullPolicy: '"$(PULL_POLICY)"'@' $(TARGET_RESOURCE)

.PHONY: set-manifest-image
set-manifest-image: ## Setup manifest image
	$(info Updating kustomize image patch file for manager resource)
	sed -i'' -e 's@image: .*@image: '"${MANIFEST_IMG}:$(MANIFEST_TAG)"'@' $(TARGET_RESOURCE)

## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

controller-gen: $(CONTROLLER_GEN) ## Build a local copy of controller-gen
conversion-gen: $(CONVERSION_GEN) ## Build a local copy of conversion-gen
kpromo: $(KPROMO)  ## Build a local copy of kpromo
golangci-lint: $(GOLANGCI_LINT) ## Build a local copy of golangci-lint
kustomize: $(KUSTOMIZE) ## Build a local copy of kustomize

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(CONVERSION_GEN): $(TOOLS_DIR)/go.mod # Build conversion-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/conversion-gen k8s.io/code-generator/cmd/conversion-gen

$(KPROMO): $(TOOLS_DIR)/go.mod # Build kpromo from tools folder.
	cd $(TOOLS_DIR) && go build -tags=tools -o $(BIN_DIR)/kpromo sigs.k8s.io/promo-tools/v3/cmd/kpromo

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod # Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(KUSTOMIZE): # Download kustomize using hack script into tools folder.
	KUSTOMIZE_VERSION=4.5.1 hack/ensure-kustomize.sh
