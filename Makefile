GOARCH=arm
GOOS=linux
GOARM=5
ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
NAME=golamp

all: lint test build

build:
	cd ${ROOT_DIR} && GOARCH=${GOARCH} GOOS=${GOOS} GOARM=${GOARM} go build ${EXTRA_GO_LINK_FLAGS} -o ${NAME} .

lint:
	goreportcard-cli -v

test:
	GOARCH=${GOARCH} GOOS=${GOOS} go test -cover -v ./...
