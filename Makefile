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
	protoc -I routeguide/ routeguide/route_guide.proto --go_out=plugins=grpc:routeguide
build: clean format
	go build -o bin/routeguide -v .
install:
	go install -v .
run: build
	./bin/routeguide