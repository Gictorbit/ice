# Run the whole application stack with Docker Compose
run:
	docker compose up --build

# Run unit tests
test:
	go test ./...

# Run benchmarks
benchmark:
	go test -bench=. ./...

# Build the Go binary using the local Go environment (not Docker)
build:
	go build -o bin/todoservice ./cmd/main.go

# Build Docker image
docker:
	docker build -t todoservice:latest .

# Clean up compiled binaries
clean:
	rm -rf bin/
# generate go commands and mocks
generate:
	go generate ./...