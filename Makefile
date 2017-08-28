.PHONY: build build-alpine clean test help default

BIN_NAME=ipfix-forwarder

VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
IMAGE_NAME := "none/ipfix-forwarder"

default: test

help:
	@echo 'Management commands for ipfix-forwarder:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs glide install, mostly used for ci.'
	
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X main.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X main.VersionPrerelease=DEV" -o bin/${BIN_NAME}

get-deps:
	@echo "GOPATH=${GOPATH}"
	glide install

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test -v $(glide nv)

coverage:
	go test -covermode=count -coverprofile=coverage.out
	go tool cover -func=coverage.out

coverage-html:
	go test -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	go get -u github.com/golang/lint/golint
	golint .

