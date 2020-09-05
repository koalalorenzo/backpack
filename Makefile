GIT_COMMIT_SHORT ?= $(shell git log -1 --pretty=format:"%h")
GOOS ?=
CGO_ENABLED ?=

clean:
	rm -rf vendor
	rm -rf out
.PHONY: clean

vendor:
	go mod vendor -v

build/backpack: vendor
	go build -race -x -a -o build/backpack main.go

build: build/backpack

test:
	go test -v ./...
.PHONY: test
