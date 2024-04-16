.PHONY: all build clean run test vet

all: clean vet test build

build: clean
	@$(info Building mantas-cli...)
	@mkdir -p bin
	@go build -o bin/mantas-cli ./cmd

clean:
	@$(info Cleaning up...)
	@rm -rf ./bin

run: build
	@$(info Running mantas-cli...)
	@./bin/mantas-cli -h

test:
	@$(info Running tests...)
	@go test ./...

vet:
	@$(info Running go vet...)
	@go vet ./...