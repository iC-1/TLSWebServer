SHELL:=/bin/bash
GOBIN:=$(GOROOT)/bin

COMMIT := $(shell git rev-parse --short dev)
VERSION := $(shell git tag --points-at dev)
BUILDTIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
BRANCH := $(shell git branch | grep \* | cut -d ' ' -f2)


GOLDFLAGS += -X main.Version=$(VERSION)
GOLDFLAGS += -X main.Commit=$(COMMIT)
GOLDFLAGS += -X main.Buildtime=$(BUILDTIME)
GOLDFLAGS += -X main.Branch=$(BRANCH)
GOLDFLAGS += -w -s
GOFLAGS = -ldflags "$(GOLDFLAGS)"

run: build
	mkdir -p tls
	go run cmd/gencert/generate_cert.go --rsa-bits=2048 --host=localhost
	mv cert.pem tls/
	mv key.pem tls/
	mkdir -p www
	echo "i am the index page" > www/index.html
	./TLSWebServer

build:
	$(GOBIN)/go build $(GOFLAGS) .

