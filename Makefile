BINARY=webhash

all: clean test
	@export GOOS=linux GOARCH=386; $(MAKE) build;
	@export GOOS=linux GOARCH=amd64; $(MAKE) build;
	@export GOOS=linux GOARCH=arm; $(MAKE) build;
	@export GOOS=freebsd GOARCH=386; $(MAKE) build;
	@export GOOS=freebsd GOARCH=amd64; $(MAKE) build;
	@export GOOS=darwin GOARCH=amd64; $(MAKE) build;
	@export GOOS=windows GOARCH=386 EXT=.exe; $(MAKE) build;
	@export GOOS=windows GOARCH=amd64 EXT=.exe; $(MAKE) build;
build:
	@mkdir -p bin
	@echo "[${GOOS}_${GOARCH}] build ..." && \
	CGO_ENABLED=0 go build -o bin/$(BINARY)_$(GOOS)_$(GOARCH)${EXT} main/main.go
clean:
	go clean
	rm -rf bin
test:
	go test -v ./...
