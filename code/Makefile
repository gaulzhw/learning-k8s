# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL := help

#
# Go.
#
GO_VERSION ?= 1.18-alpine
GO_CONTAINER_IMAGE ?= golang:$(GO_VERSION)

#
# Directories.
#
# Full directory of where the Makefile resides
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
BIN_DIR := bin
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/$(BIN_DIR)

#
# Binaries.
#
# Note: Need to use abspath so we can invoke these from subdirectories
GOLANGCI_LINT := $(abspath $(TOOLS_BIN_DIR)/golangci-lint)
# code gen
CONTROLLER_GEN := $(abspath $(TOOLS_BIN_DIR)/controller-gen)
CONVERSION_GEN := $(abspath $(TOOLS_BIN_DIR)/conversion-gen)
# code gen
CLIENT_GEN := $(abspath $(TOOLS_BIN_DIR)/client-gen)
CONVERSION_GEN := $(abspath $(TOOLS_BIN_DIR)/conversion-gen)
DEEPCOPY_GEN := $(abspath $(TOOLS_BIN_DIR)/deepcopy-gen)
DEFAULTER_GEN := $(abspath $(TOOLS_BIN_DIR)/defaulter-gen)
INFORMER_GEN := $(abspath $(TOOLS_BIN_DIR)/informer-gen)
LISTER_GEN := $(abspath $(TOOLS_BIN_DIR)/lister-gen)

# Define Docker related variables. Releases should modify and double check these vars.
REGISTRY ?= github.com/learning_k8s

#
# Images.
#
IMAGE_NAME ?= example
CONTROLLER_IMG ?= $(REGISTRY)/$(IMAGE_NAME)

# release
RELEASE_TAG ?= $(shell git describe --tags --abbrev=0)

help:  # Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[0-9A-Za-z_-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Generate / Manifests
## --------------------------------------

##@ generate:

.PHONY: generate
generate: ## Run all generate-xxx targets
	$(MAKE) generate-go-deepcopy generate-go-convertion generate-manifests

.PHONY: generate-go-deepcopy
generate-go-deepcopy: $(CONTROLLER_GEN) ## Generate deepcopy code
	$(CONTROLLER_GEN) \
		object:headerFile=$(BOILERPLATE_FILE) \
		paths=./...

.PHONY: generate-go-convertion
generate-go-convertion: $(CONVERSION_GEN) ## Generate convertion code
	$(CONVERSION_GEN) \
		--go-header-file=$(BOILERPLATE_FILE) \
		--input-dirs=github.com/gaulzhw/leaning_k8s/pkg/apis/apps/v1alpha1 \
	  	--output-file-base=zz_generated.conversion

.PHONY: generate-manifests
generate-manifests: $(CONTROLLER_GEN) ## Generate manifests e.g. CRD, RBAC etc
	$(CONTROLLER_GEN) \
		paths=./... \
		crd:crdVersions=v1 \
		rbac:roleName=mcp-manager \
		output:crd:dir=deploy/crd \
		output:rbac:dir=deploy/rbac

.PHONY: generate-go-client
generate-go-client: $(CLIENT_GEN) ## Generate client code
	$(CLIENT_GEN) \
		--go-header-file=hack/boilerplate.go.txt \
		--input-base="" \
		--input=github.com/gaulzhw/leaning_k8s/pkg/apis/apps/v1alpha1 \
		--clientset-name=versioned \
		--output-package=github.com/multi-cluster-platform/mcp/pkg/generated

.PHONY: generate-go-lister
generate-go-lister: $(LISTER_GEN) ## Generate lister code
	$(LISTER_GEN) \
		--go-header-file=hack/boilerplate.go.txt \
		--input-dirs=github.com/gaulzhw/leaning_k8s/pkg/apis/apps/v1alpha1 \
		--output-package=github.com/multi-cluster-platform/mcp/pkg/generated/listers

.PHONY: generate-go-informer
generate-go-informer: $(INFORMER_GEN) ## Generate informer code
	$(INFORMER_GEN) \
		--go-header-file=hack/boilerplate.go.txt \
		--input-dirs=github.com/gaulzhw/leaning_k8s/pkg/apis/apps/v1alpha1 \
		--versioned-clientset-package=github.com/gaulzhw/leaning_k8s/pkg/generated/versioned \
		--listers-package=github.com/multi-cluster-platform/mcp/pkg/generated/listers \
		--output-package=github.com/multi-cluster-platform/mcp/pkg/generated/informers

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
	GO111MODULE=off $(GOLANGCI_LINT) run -v

## --------------------------------------
## Docker
## --------------------------------------

##@ docker:

.PHONY: docker-build
docker-build: ## Build image
	docker build --build-arg builder_image=$(GO_CONTAINER_IMAGE) --build-arg package=cmd/controller/main.go . -t $(CONTROLLER_IMG):$(RELEASE_TAG)

.PHONY: docker-push
docker-push: ## Push image
	docker push $(CONTROLLER_IMG):$(RELEASE_TAG)

.PHONY: set-manifest
set-manifest: ## Update manifest image and pull policy
	$(MAKE) set-manifest-image MANIFEST_IMG=$(CONTROLLER_IMG) MANIFEST_TAG=$(RELEASE_TAG) TARGET_RESOURCE="./deploy/base/controller.yaml"
	$(MAKE) set-manifest-pull-policy PULL_POLICY=IfNotPresent TARGET_RESOURCE="./deploy/base/controller.yaml"

.PHONY: set-manifest-pull-policy
set-manifest-pull-policy: ## Update manifest pull policy
	sed -i'' -e 's@imagePullPolicy: .*@imagePullPolicy: '"$(PULL_POLICY)"'@' $(TARGET_RESOURCE)

.PHONY: set-manifest-image
set-manifest-image: ## Update manifest image
	sed -i'' -e 's@image: .*@image: '"${MANIFEST_IMG}:$(MANIFEST_TAG)"'@' $(TARGET_RESOURCE)

## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

golangci-lint: $(GOLANGCI_LINT) ## Build a local copy of golangci-lint
controller-gen: $(CONTROLLER_GEN) ## Build a local copy of controller-gen
client-gen: $(CLIENT_GEN) ## Build a local copy of client-gen
conversion-gen: $(CONVERSION_GEN) ## Build a local copy of conversion-gen
deepcopy-gen: $(DEEPCOPY_GEN) ## Build a local copy of deepcopy-gen
defaulter-gen: $(DEFAULTER_GEN) ## Build a local copy of defaulter-gen
informer-gen: $(INFORMER_GEN) ## Build a local copy of informer-gen
lister-gen: $(LISTER_GEN) ## Build a local copy of lister-gen
kustomize: $(KUSTOMIZE) ## Build a local copy of kustomize

$(GOLANGCI_LINT): $(TOOLS_DIR)/go.mod # Build golangci-lint from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint

$(CONTROLLER_GEN): $(TOOLS_DIR)/go.mod # Build controller-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/controller-gen sigs.k8s.io/controller-tools/cmd/controller-gen

$(CLIENT_GEN): $(TOOLS_DIR)/go.mod # Build client-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/client-gen k8s.io/code-generator/cmd/client-gen

$(CONVERSION_GEN): $(TOOLS_DIR)/go.mod # Build conversion-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/conversion-gen k8s.io/code-generator/cmd/conversion-gen

$(DEEPCOPY_GEN): $(TOOLS_DIR)/go.mod # Build deepcopy-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/deepcopy-gen k8s.io/code-generator/cmd/deepcopy-gen

$(DEFAULTER_GEN): $(TOOLS_DIR)/go.mod # Build defatul-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/defaulter-gen k8s.io/code-generator/cmd/defaulter-gen

$(INFORMER_GEN): $(TOOLS_DIR)/go.mod # Build informer-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/informer-gen k8s.io/code-generator/cmd/informer-gen

$(LISTER_GEN): $(TOOLS_DIR)/go.mod # Build lister-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/lister-gen k8s.io/code-generator/cmd/lister-gen

$(KUSTOMIZE): # Download kustomize using hack script into tools folder.
	KUSTOMIZE_VERSION=4.5.1 hack/ensure-kustomize.sh
