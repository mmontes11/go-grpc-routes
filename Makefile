.PHONY: clean deps deps-sync format build install run

all: build
clean:
	rm -rf bin/
deps:
	go get -u -v
deps-sync:
	go mod vendor
format:
	go fmt .
proto:
	protoc -I route/ --go_opt=paths=source_relative --go_out=plugins=grpc:route route/*.proto 
build: clean format
	go build -o bin/routeguide -v .
install:
	go install -v .
run: build
	./bin/routeguide