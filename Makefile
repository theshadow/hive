all: static-analysis tests
.PHONY: formatting static-analysis test

tests:
	go test -test.v ./... ./game

static-analysis:
	go vet ./...

docker-container:
	docker build .

tag:
	git tag v$(VERSION)

formatting:
	go fmt ./...



