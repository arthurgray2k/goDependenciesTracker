.PHONY: all build test clean format vet

BINARY_NAME=goDependenciesTracker.exe

all: format vet test build

build:
	go build -o $(BINARY_NAME) ./cmd/goDependenciesTracker

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

format:
	go fmt ./...

vet:
	go vet ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -f coverage.out
