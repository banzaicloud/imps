# Image URL to use all building/pushing image targets

RACE_DETECTOR ?= 0

CHART_NAME = imagepullsecrets

# Image URL to use all building/pushing image targets
IMG ?= imagepullsecrects:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"
# TODO: Use this when allowDangerousTypes feature is released to support floats
# CRD_OPTIONS ?= "crd:trivialVersions=true,allowDangerousTypes=true"
LICENSEI_VERSION = 0.3.1
GOLANGCI_VERSION = 1.26.0

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

SERVICE_NAME=$(shell basename ${CURDIR} )
REPO_ROOT=$(shell git rev-parse --show-toplevel)
MAIN_PACKAGE ?= main.go

COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE ?= $(shell date +%FT%T%z)
VERSION ?= 0.1.1
LDFLAGS += -X github.com/banzaicloud/imps/internal/version.commitHash=${COMMIT_HASH}
LDFLAGS += -X github.com/banzaicloud/imps/internal/version.buildDate=${BUILD_DATE}
LDFLAGS += -X github.com/banzaicloud/imps/internal/version.version=${VERSION}

ifeq (${RACE_DETECTOR}, 1)
    GOARGS += -race
    CGO_ENABLED = 1
endif
export CGO_ENABLED ?= 0

ifeq (${REMOTE_DEBUGGING}, 1)
	# disables inlining and optimisations
	GOARGS += -gcflags="all=-N -l"
endif

.PHONY: all
all: build

.PHONY: test
test: ensure-tools generate fmt vet manifests 	## Run tests
	KUBEBUILDER_ASSETS="${REPO_ROOT}/bin/kubebuilder-2.3.1/bin/" go test  ${GOARGS} ./... -coverprofile cover.out

bin/golangci-lint: bin/golangci-lint-${GOLANGCI_VERSION}
	@ln -sf golangci-lint-${GOLANGCI_VERSION} bin/golangci-lint
bin/golangci-lint-${GOLANGCI_VERSION}:
	@mkdir -p bin
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | bash -s -- -b ./bin/ v${GOLANGCI_VERSION}
	@mv bin/golangci-lint $@

DISABLED_LINTERS ?= --disable=gci --disable=goimports --disable=gofumpt
.PHONY: lint
lint: bin/golangci-lint ## Run linter
# "unused" linter is a memory hog, but running it separately keeps it contained (probably because of caching)
	bin/golangci-lint run --disable=unused -c .golangci.yml --timeout 2m
	bin/golangci-lint run -c .golangci.yml --timeout 2m

.PHONY: lint-fix
lint-fix: bin/golangci-lint ## Run linter & fix
# "unused" linter is a memory hog, but running it separately keeps it contained (probably because of caching)
	bin/golangci-lint run --disable=unused -c .golangci.yml --fix
	bin/golangci-lint run -c .golangci.yml --fix

.PHONY: build
build: generate fmt vet 	## Build the binary
	go build  ${GOARGS} -o bin/${SERVICE_NAME} -ldflags "${LDFLAGS}" ${MAIN_PACKAGE}

.PHONY: binary
binary:					## Build the binary without executing any code generators
	go build  ${GOARGS} -o bin/${SERVICE_NAME} -ldflags "${LDFLAGS}" ${MAIN_PACKAGE}

.PHONY: run
run: generate fmt vet manifests		## Run against the configured Kubernetes cluster in ~/.kube/config
	go run  ${GOARGS} ${MAIN_PACKAGE}

.PHONY: ensure-tools
ensure-tools:
	@scripts/download-deps.sh
	@scripts/install_kustomize.sh

.PHONY: install
install: ensure-tools manifests  		## Install CRDs into a cluster
	${REPO_ROOT}/bin/kustomize build config/crd | kubectl apply -f -

.PHONY: uninstall
uninstall: ensure-tools manifests  	## Uninstall CRDs from a cluster
	${REPO_ROOT}/bin/kustomize build config/crd | kubectl delete -f -

.PHONY: deploy
deploy: ensure-tools  manifests			## Deploy controller in the configured Kubernetes cluster in ~/.kube/config
	cd config/manager && ${REPO_ROOT}/bin/kustomize edit set image controller=${IMG}
	${REPO_ROOT}/bin/kustomize build config/default | kubectl apply -f -

# Generate manifests e.g. CRD, RBAC etc.
.PHONY: manifests
manifests: ensure-tools
	${REPO_ROOT}/bin/controller-gen $(CRD_OPTIONS) rbac:roleName=binary-role object webhook paths="./..." output:crd:artifacts:config=config/crd/bases

.PHONY: fmt
fmt:	## Run go fmt against code
	go fmt ./...

.PHONY: vet
vet:	## Run go vet against code
	go vet ./...

.PHONY: go-generate
go-generate: generate-generate
	go generate ./...

.PHONY: generate-generate
generate-generate:
	@${REPO_ROOT}/scripts/generate_generate.sh .

# Generate code
.PHONY: generate
generate: go-generate ensure-tools manifests generate-helm-crds		## Generate manifests, CRDs, static assets

.PHONY: generate-helm-crds
generate-helm-crds: ensure-tools	## Update the CRDs in our helm charts
	${REPO_ROOT}/bin/kustomize build config/crd > deploy/charts/${CHART_NAME}/crds/crds.yaml

.PHONY: docker-build
docker-build: test		## Build the docker image (to override image name please set IMG)
	docker build . --build-arg LDFLAGS="${LDFLAGS}" -f ${CURDIR}/Dockerfile -t ${IMG}

.PHONY: docker-push
docker-push:			## Push the docker image (to override image name please set IMG)
	docker push ${IMG}

bin/licensei: bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} bin/licensei
bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/goph/licensei/master/install.sh | bash -s v${LICENSEI_VERSION}
	@mv bin/licensei $@

.PHONY: license-check
license-check: bin/licensei ## Run license check
	bin/licensei check
	bin/licensei header
.PHONY: license-cache
license-cache: bin/licensei ## Generate license cache
	bin/licensei cache

MAKEFILE_LIST=Makefile

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

