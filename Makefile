SHELL:=/bin/sh
.PHONY: build build_client build_server build_docker \
		test run fmt vet clean \
		mod_update vendor_from_mod vendor_clean

export GO111MODULE=on

# Path Related
MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
MKFILE_DIR := $(dir $(MKFILE_PATH))
RELEASE_DIR := ${MKFILE_DIR}bin

# Version
RELEASE?=1.0.0

# Git Related
GIT_REPO_INFO=$(shell cd ${MKFILE_DIR} && git config --get remote.origin.url)
ifndef GIT_COMMIT
  GIT_COMMIT := git-$(shell git rev-parse --short HEAD)
endif

# Build Flags
GO_LD_FLAGS= "-s -w -X github.com/megaease/easegress/pkg/version.RELEASE=${RELEASE} -X github.com/megaease/easegress/pkg/version.COMMIT=${GIT_COMMIT} -X github.com/megaease/easegress/pkg/version.REPO=${GIT_REPO_INFO}"

# Targets
TARGET_SERVER=${RELEASE_DIR}/easegress-server
TARGET_CLIENT=${RELEASE_DIR}/egctl

# Rules
build: build_client build_server

build_client:
	@echo "build client"
	cd ${MKFILE_DIR} && \
	CGO_ENABLED=0 go build -v -trimpath -ldflags ${GO_LD_FLAGS} \
	-o ${TARGET_CLIENT} ${MKFILE_DIR}cmd/client

build_server:
	@echo "build server"
	cd ${MKFILE_DIR} && \
	CGO_ENABLED=0 go build -v -trimpath -ldflags ${GO_LD_FLAGS} \
	-o ${TARGET_SERVER} ${MKFILE_DIR}cmd/server

dev_build_client:
	@echo "build dev client"
	cd ${MKFILE_DIR} && \
	go build -v -race -ldflags ${GO_LD_FLAGS} \
	-o ${TARGET_CLIENT} ${MKFILE_DIR}cmd/client

dev_build_server:
	@echo "build dev server"
	cd ${MKFILE_DIR} && \
	go build -v -race -ldflags ${GO_LD_FLAGS} \
	-o ${TARGET_SERVER} ${MKFILE_DIR}cmd/server

build_docker:
	docker build -t megaease/easegress:${RELEASE} -f ./build/package/Dockerfile .

test:
	@go list ./... | grep -v -E 'vendor' | xargs -n1 go test

clean:
	rm -rf ${RELEASE_DIR}

run: build_server

fmt:
	cd ${MKFILE_DIR} && go fmt ./...

vet:
	cd ${MKFILE_DIR} && go vet ./...

vendor_from_mod:
	cd ${MKFILE_DIR} && go mod vendor

vendor_clean:
	rm -rf ${MKFILE_DIR}vendor

mod_update:
	cd ${MKFILE_DIR} && go get -u
