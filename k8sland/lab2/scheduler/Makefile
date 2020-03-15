PACKAGE    := github.com/k8sland/lab2/scheduler
VERSION    := 0.0.1
IMAGE      := go-partysched
APP        := partysched
REGISTRY   := wrldsrf
IMAGE_NAME := ${REGISTRY}/${IMAGE}:${VERSION}

default: help

build:       ## Builds CLI
	@env GOOS=linux GOARCH=amd64 go build \
	-ldflags "-w -X ${PACKAGE}/cmd.Version=${VERSION}" \
	-a -tags netgo -o execs/${APP} main.go

img: build   ## Build Docker Image
	@docker build --rm -t ${IMAGE_NAME} .

push: img    ## Push Docker Image
	@docker push ${IMAGE_NAME}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;168m %s\n", $$1, $$2}'
