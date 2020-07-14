.PHONY: clean deps fmt-srv fmt-cli fmt proto build srv cli cli-tls

TLS_ARGS = -tls true

all: build
clean:
	rm -rf bin/
deps:
	go get -u -v
fmt-srv: 
	go fmt server/*.go 
fmt-cli: 
	go fmt client/*.go
fmt: fmt-srv fmt-cli
proto:
	protoc -I pb/ --go_opt=paths=source_relative --go_out=plugins=grpc:pb pb/*.proto
build: fmt-srv proto clean
	go build -o bin/route -v .
srv: build
	./bin/route
cli: fmt-cli
	go run client/*.go
cli-tls: fmt-cli
	go run client/*.go $(TLS_ARGS)
