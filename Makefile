BINARY=webhash
GOOS=linux darwin freebsd windows
GOARCH=386 amd64

all: test build
build:
	mkdir -p bin
	go build -o bin/$(BINARY) main/main.go
test:
	go test -v ./...
clean:
	go clean
	rm -rf bin
buildall:
	$(foreach OS, $(GOOS),\
	$(foreach ARCH, $(GOARCH),\
	$(shell export GOOS=$(OS); export GOARCH=$(ARCH); go build -o bin/$(BINARY)_$(OS)_$(ARCH) main/main.go)))
