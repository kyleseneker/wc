.PHONY: build fmt clean lint

OS := $(shell go env GOOS)
BUILDCMD=env GOOS=$(OS) GOARCH=amd64 go build -v

build:
	$(BUILDCMD) -o wc cmd/wc/*.go

fmt:
	gofmt -w *.go

clean:
	go clean ./...
	rm -rf ./wc

lint:
	golangci-lint run ./...
