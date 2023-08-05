build:
		@go build -o bin/greenlight

run: build
		@./bin/greenlight

test:
		@go test -v ./...
