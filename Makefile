.PHONY: build
build: vet
	@echo "Building todo-api..."
	@go build -o bin/todo-api ./cmd/todo-api

.PHONY: run
run: build
	@echo "Starting todo-api on port 8080..."
	@DATABASE_URL="postgres://localhost/todo?sslmode=disable" ./bin/todo-api

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@go clean

.PHONY: check
check:
	@echo "Formatting code..."
	go fmt ./...
	@echo "Running go vet..."
	go vet ./...
	@echo "Running tests..."
	go test -v ./...
	@echo "All checks passed!"