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
stubs:
	protoc -I routeguide/ --go_opt=paths=source_relative --go_out=plugins=grpc:routeguide routeguide/*.proto 
build: clean format
	go build -o bin/routeguide -v .
install:
	go install -v .
run: build
	./bin/routeguide