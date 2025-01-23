help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build		- Build the application"
	@echo "  clean		- Clean the application"
	@echo "  env		- Load environment variables"
	@echo "  run		- Run the application"
	@echo "  restart	- Restart the application in prod"
	@echo "  test		- Run tests"
	@echo "  lint		- Run linter"
	@echo "  logs		- Display logs"
	@echo "  help		- Display this help message"
	@echo ""
	@echo "For more information, RTFM!"


build:
	@echo "Building the application..."
	@go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go
	@echo "Application built successfully"

clean:
	@echo "Cleaning the application..."
	@rm -rf bin
	@echo "Application cleaned successfully"

env:
	@echo "Loading environment variables..."
	@source .env
	@echo "Environment variables loaded successfully"

run:
	@echo "Running the application..."
	@go run cmd/$(APP_NAME)/main.go

restart:
	@echo "Restarting the application..."
	@kill -9 $$(lsof -t -i:$(PORT))
	@go run cmd/$(APP_NAME)/main.go

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@go fmt ./...