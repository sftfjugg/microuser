.PHONY: proto
proto:
    
	protoc --micro_out=./proto/microuser --go_out=./proto/microuser ./proto/microuser/microuser.proto
    

.PHONY: build
build: proto

	go build -o microuser-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t microuser-service:latest
