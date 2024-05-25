.PHONY: build fmt clean lint

OS := $(shell go env GOOS)
BUILDCMD=env GOOS=$(OS) GOARCH=amd64 go build -v

build:
	@$(BUILDCMD) -o wc cmd/wc/*.go 

fmt:
	@go fmt ./...

clean:
	@go clean ./...
	@rm -rf ./wc

lint:
	@golangci-lint run ./...

benchmark: build
	@hyperfine --warmup 2 './wc examples/test1.txt examples/test2.txt' 'wc examples/test1.txt examples/test2.txt'
	@$(MAKE) clean
