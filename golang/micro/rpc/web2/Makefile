
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	go build -o web2-web *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t web2-web:latest
