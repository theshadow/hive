VERSION ?= $(shell head -n 1 VERSION)

all: test static-analysis build
.PHONY: test static-analysis build formatting _docs _artifacts

test:
	go test .../..

formatting:
	go fmt .../..

static-analysis:
	go vet .../..

build: _artifacts _docs

_docs:
	cd docs; make;

_artifacts:
	go build -buildmode=plugin -o hived-$(VERSION).so .../..


