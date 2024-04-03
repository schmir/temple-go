default: lint test run

# Run tests
test:
    go test ./...

# Run golangci-lint
lint:
    golangci-lint run

# Run the main binary
run:
    go run ./
