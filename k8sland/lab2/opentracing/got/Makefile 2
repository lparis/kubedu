PACKAGE         := github.com/k8sland/code2/tracing/bin
CASTLE_VERSION  := 0.0.1
CASTLE_IMAGE    := ot-castle-svc-go
CASTLE_APP      := castle
CASTLE_PACKAGE  := ${PACKAGE}/${CASTLE_APP}
KNIGHT_VERSION  := 0.0.1
KNIGHT_IMAGE    := ot-knight-svc-go
KNIGHT_APP      := knight
KNIGTH_PACKAGE  := ${PACKAGE}/${KNIGHT_APP}
REGISTRY        := <<!!!CHANGE_ME!!!>>

default: help

cbuild:              ## Builds Castle CLI
	@env GOOS=linux GOARCH=amd64 go build \
	-ldflags "-w -X ${CASTLE_PACKAGE}/cmd.Version=${CASTLE_VERSION}" \
	-a -tags netgo -o execs/${CASTLE_APP} bin/castle/main.go

cimg: cbuild         ## Build Castle Docker Image
	@docker build -f bin/castle/Dockerfile --rm -t $(REGISTRY)/$(CASTLE_IMAGE):$(CASTLE_VERSION) .

cpush: cimg          ## Push Castle Docker Image
	@docker push $(REGISTRY)/$(CASTLE_IMAGE):$(CASTLE_VERSION)

kbuild:              ## Builds Knight CLI
	@env GOOS=linux GOARCH=amd64 go build \
	-ldflags "-w -X ${KNIGHT_PACKAGE}/cmd.Version=${KNIGHT_VERSION}" \
	-a -tags netgo -o execs/${KNIGHT_APP} bin/knight/main.go

kimg: kbuild         ## Build Knight Docker Image
	@docker build -f bin/knight/Dockerfile --rm -t $(REGISTRY)/$(KNIGHT_IMAGE):$(KNIGHT_VERSION) .

kpush: kimg          ## Push Knight Docker Image
	@docker push $(REGISTRY)/$(KNIGHT_IMAGE):$(KNIGHT_VERSION)

build: cbuild kbuild ## Build all execs

img: cimg kimg       ## Push all images

push: cpush kpush    ## Push all images


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[38;5;69m%-30s\033[38;5;168m %s\n", $$1, $$2}'
