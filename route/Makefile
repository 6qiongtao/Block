GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:

	GOOS=linux GOARCH=amd64 go build -o  route-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t route-web:latest
