all: static-analysis tests
.PHONY: formatting static-analysis test coverage-report

coverage-report:
	go test -race -covermode=atomic -coverprofile=coverage.out

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



