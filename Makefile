build:
	@go build -o bin/go-bank-api

run: build
	@./bin/go-bank-api

test: 
	@go test -v ./...