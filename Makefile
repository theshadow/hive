VERSION ?=$(shell head -n 1 VERSION)
GITREF ?=$(shell git describe --always --dirty 2> /dev/null || echo 0000000)
BUILD ?=$(GITREF)

OUTPUT_PREFIX ?=hived-$(VERSION)

OUTPUT_LIB :=$(OUTPUT_PREFIX).so
OUTPUT_ARCHIVE :=$(OUTPUT_PREFIX).zip

LDFLAGS :=-ldflags "-X github.com/theshadow/hived.Version=$(VERSION) -X github.com/theshadow/hived.BuildID=$(BUILD)"

all: static-analysis tests build
.PHONY: build formatting static-analysis test _archive _docs

tests:
	go test .../..

static-analysis:
	go vet .../..

docker-container: _formatting
	docker build -t theshadow/hived .

build: _docs _archive

tag:
	git tag v$(VERSION)

_formatting:
	go fmt .../..

_docs:
	cd docs; sphinx-build -b html -D release=$(VERSION) . _build

_archive:
	mkdir -p _build/docs
	cp --recursive docs/_build/* _build/docs
	cp LICENSE _build
	cp VERSION _build
	cd _build; zip -r $(OUTPUT_ARCHIVE) .


