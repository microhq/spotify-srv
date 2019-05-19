GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc protoc --proto_path=. --go_out=. --micro_out=. proto/spotify/spotify.proto

.PHONY: build
build: proto

	go build -o spotify-srv main.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t microhq/spotify-srv:latest

.PHONY: docker-multistage
docker:
	docker build -f Dockerfile.multistage . -t microhq/spotify-srv:latest
