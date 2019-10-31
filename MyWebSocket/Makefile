
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/MyWebSocket/MyWebSocket.proto

.PHONY: build
build: proto

	GOOS=linux GOARCH=amd64 go build -o MyWebSocket-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t MyWebSocket-srv:latest
