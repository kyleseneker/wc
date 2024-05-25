.PHONY: build fmt clean

OS := $(shell go env GOOS)
BUILDCMD=env GOOS=$(OS) GOARCH=amd64 go build -v

build:
	$(BUILDCMD) -o wc cmd/*.go

fmt:
	gofmt -w *.go

clean:
	go clean
