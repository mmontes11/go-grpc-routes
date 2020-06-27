.PHONY: clean deps deps-sync fmt-srv fmt-cli fmt proto build install srv cli

all: build
clean:
	rm -rf bin/
deps:
	go get -u -v
deps-sync:
	go mod vendor
fmt-srv: 
	go fmt server/*.go 
fmt-cli: 
	go fmt client/*.go
fmt: fmt-srv fmt-cli
proto:
	protoc -I route/ --go_opt=paths=source_relative --go_out=plugins=grpc:route route/*.proto
build: fmt-srv proto clean
	go build -o bin/route -v .
install:
	go install -v .
srv: build
	./bin/route
cli: fmt-cli
	go run client/client.go
