CGO_ENABLED ?= 0

.EXPORT_ALL_VARIABLES:

GIT_COMMIT_SHORT ?= $(shell git log -1 --pretty=format:"%h")
ifdef GOOS
ifdef GOARCH
GO_BUILD_OUT_FILE := build/backpack_${GIT_COMMIT_SHORT}_${GOOS}_${GOARCH}
endif
endif

_GO_PACKAGE ?= gitlab.com/qm64/backpack
GO_BUILD_OUT_FILE ?= build/backpack

.DEFAULT_GOAL := build

ifneq (${VERSION},)
BUILD_LDFLAGS := -X ${_GO_PACKAGE}/cmd.version=${VERSION} -X ${_GO_PACKAGE}/cmd.versionGitHash=${GIT_COMMIT_SHORT}
endif
BUILD_LDFLAGS ?= -X ${_GO_PACKAGE}/cmd.versionGitHash=${GIT_COMMIT_SHORT}

clean:
	rm -rf vendor
	rm -rf build
	rm -rf backpack
.PHONY: clean

vendor:
	go mod vendor -v

build: vendor
	go build -a -ldflags "${BUILD_LDFLAGS}" -o ${GO_BUILD_OUT_FILE} main.go
.PHONY: build

install:
	go install -i
.PHONY: install

test:
	go test -v ./...
.PHONY: test
