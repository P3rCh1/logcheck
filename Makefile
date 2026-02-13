.PHONY: build clean run

build:
	golangci-lint custom -v

clean:
	rm -rf bin

run: build
	./bin/logcheck run ./...

version: build
	./bin/logcheck version

example: build
	@cd internal/analyzer/testdata && ../../../bin/logcheck run ./...