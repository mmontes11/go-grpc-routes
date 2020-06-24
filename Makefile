.PHONY: clean deps deps-sync format proto server

all: server
clean:
	rm -rf bin/
deps:
	go get -u -v
deps-sync:
	go mod vendor
format:
	go fmt server/*.go
proto:
	protoc -I route/ --go_opt=paths=source_relative --go_out=plugins=grpc:route route/*.proto
build: proto clean format
	go build -o bin/route -v .
install:
	go install -v .
run: build
	./bin/route