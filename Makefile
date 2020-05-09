BINARY_NAME=webhash
BINARY_UNIX=$(BINARY_NAME)_unix
GOOS:=linux darwin freebsd windows
GOARCH:=386 amd64

all: test build
build:
	mkdir -p dist
	go build -o dist/$(BINARY_NAME) main/main.go
test:
	go test -v ./...
clean:
	go clean
	rm -rf dist
