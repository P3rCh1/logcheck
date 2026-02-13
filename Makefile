.PHONY: build clean example

build:
	golangci-lint custom -v

clean:
	rm -rf bin

example: build
	@cd internal/analyzer/testdata && ../../../bin/logcheck run ./...