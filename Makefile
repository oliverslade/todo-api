.PHONY: generate
generate:
	@echo "Generating protobuf code..."
	@mkdir -p proto/todo/v1
	protoc --go_out=. --go-grpc_out=. proto/todo.proto
	@if [ -d "github.com" ]; then \
		mv github.com/oliverslade/todo-api/proto/todo/v1/* proto/todo/v1/ && \
		rm -rf github.com; \
	fi

.PHONY: build
build: generate
	@echo "Building todo-api..."
	@go build -o bin/todo-api ./cmd/todo-api

.PHONY: run
run: build
	@echo "Starting todo-api on port 8080..."
	./bin/todo-api

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -f proto/todo/v1/*.pb.go
	@go clean

.PHONY: check
check: generate
	@echo "Formatting code..."
	go fmt ./...
	@echo "Running go vet..."
	go vet ./...
	@echo "Running tests..."
	go test -v ./...
	@echo "All checks passed!"