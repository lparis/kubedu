PACKAGE      := github.com/k8sland/lab2/prom/hangman
REGISTRY     := k8sland

SVC_VERSION  := 0.0.1
SVC_IMAGE    := go-hangman-prom-svc
SVC_APP      := hangman
SVC_PACKAGE  := ${PACKAGE}/${S_APP}
SVC_IMG_NAME := ${REGISTRY}/${SVC_IMAGE}:${SVC_VERSION}
CLI_VERSION  := 0.0.1
CLI_IMAGE    := go-hangman-cli
CLI_APP      := hangman_cli
CLI_PACKAGE  := ${PACKAGE}/${CLI_APP}
CLI_IMG_NAME := ${REGISTRY}/${CLI_IMAGE}:${CLI_VERSION}

default: help

sbuild:              ## Builds Hangman Svc
	@env GOOS=linux GOARCH=amd64 go build \
	-ldflags "-w -X ${SVC_PACKAGE}/cmd.Version=${SVC_VERSION}" \
	-a -tags netgo -o execs/${SVC_APP} cmd/svc/main.go

simg: sbuild         ## Build Hangman Svc Docker Image
	@docker build -f cmd/svc/Dockerfile --rm -t ${SVC_IMG_NAME} .

spush: simg          ## Push Hangman Svc Docker Image
	@docker push ${SVC_IMG_NAME}

cbuild:              ## Builds Hangman CLI
	@env GOOS=linux GOARCH=amd64 go build \
	-ldflags "-w -X ${CLI_PACKAGE}/cmd.Version=${CLI_VERSION}" \
	-a -tags netgo -o execs/${CLI_APP} cmd/cli/main.go

cimg: cbuild         ## Build Hangman CLI Docker Image
	@docker build -f cmd/cli/Dockerfile --rm -t ${CLI_IMG_NAME} .

cpush: cimg          ## Push Hangman CLI Docker Image
	@docker push ${CLI_IMG_NAME}

build: sbuild cbuild ## Build all execs

img: simg cimg       ## Push all images

push: spush cpush    ## Push all images


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;168m %s\n", $$1, $$2}'
