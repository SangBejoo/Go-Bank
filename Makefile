build:
	@go build -o Go-Bank

run:
	@go run .

test:
	@go test -v ./...