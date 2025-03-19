.PHONY: start build run update test clean docker-up docker-down help build-linux build-darwin-amd64 build-darwin-arm64 build-windows code-quality build-force check-go-version

# Default target when just 'make' is executed
help:
	@echo "Available commands:"
	@echo "  make check-go-version - Check if you have the required Go version (1.24+)"
	@echo "  make start           - Run the application directly from source"
	@echo "  make code-quality    - Run code quality checks"
	@echo "  make build           - Build the application binary (default: linux)"
	@echo "  make build-linux     - Build for Linux with code quality checks"
	@echo "  make build-darwin-amd64 - Build for macOS (Intel) with code quality checks"
	@echo "  make build-darwin-arm64 - Build for macOS (Apple Silicon) with code quality checks"
	@echo "  make build-windows   - Build for Windows with code quality checks"
	@echo "  make run             - Run the built binary"
	@echo "  make update          - Update dependencies"
	@echo "  make docker-up       - Start docker containers"
	@echo "  make docker-down     - Stop docker containers"
	@echo "  make unittests       - Run unit tests"
	@echo "  make unittests-verbose - Run unit tests with verbose output"
	@echo "  make unittests-coverage - Run unit tests with coverage"
	@echo "  make clean           - Clean cache and test data"
	@echo "  make setup-docker-dirs - Set up docker directories"
	@echo "  make build-force     - Build ignoring Go version constraints (use if Go version < 1.24)"

# Check Go version
check-go-version:
	@echo "Checking Go version..."
	@go version | grep -q "go1.2[4-9]" || (echo "Error: This project requires Go 1.24 or later. Current version:" && go version && echo "To update Go, visit: https://golang.org/dl/" && exit 1)
	@echo "Go version is compatible with this project"

# Code quality checks
code-quality: check-go-version
	go mod verify
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	golangci-lint run ./...

# Run directly from source
start: check-go-version
	go run cmd/server/main.go

# Build the application (default: linux)
build: build-linux

# Build for Linux
build-linux: check-go-version code-quality
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api ./cmd/server
	@echo "Linux build complete: ./bin/api"
	@echo "To install system-wide, run: sudo cp ./bin/api /usr/bin/"

# Build for macOS (Intel)
build-darwin-amd64: check-go-version code-quality
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api ./cmd/server
	@echo "macOS (Intel) build complete: ./bin/api"
	@echo "To install to your user bin, run: mkdir -p ~/bin && cp ./bin/api ~/bin/"

# Build for macOS (Apple Silicon)
build-darwin-arm64: check-go-version code-quality
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/api ./cmd/server
	@echo "macOS (Apple Silicon) build complete: ./bin/api"
	@echo "To install to your user bin, run: mkdir -p ~/bin && cp ./bin/api ~/bin/"

# Build for Windows
build-windows: check-go-version code-quality
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/api.exe ./cmd/server
	@echo "Windows build complete: ./bin/api.exe"

# Force build (ignoring Go version)
build-force:
	@echo "Force building (ignoring Go version requirements)..."
	@echo "Note: This might fail if your Go version is incompatible with dependencies"
	mkdir -p ./bin
	GO111MODULE=on GOPROXY=direct go build -mod=mod -ldflags="-s -w" -o ./bin/api ./cmd/server
	@echo "Force build complete: ./bin/api (version constraints ignored)"

# Run the built binary
run: build
	./bin/api

# Update dependencies
update:
	go mod tidy

# Run EAS script
run-script-eas:
	go run script/eas/main.go

# Set up Docker directories
setup-docker-dirs:
	sudo rm -rf ./.docker
	mkdir -p ./.docker/mongodb
	sudo chown -R 1001:1001 ./.docker
	sudo chmod -R 775 ./.docker

# Test targets
unittests:
	go clean -testcache && go test ./test/unittests/...

unittests-verbose:
	go clean -testcache && go test -v ./test/unittests/...

unittests-coverage:
	go clean -testcache && go test -v -coverpkg=./... -coverprofile=coverage.out ./test/unittests/...
	go tool cover -html=coverage.out

# Clean targets
clean: clean-go-cache clean-test-cache
	rm -rf ./bin

clean-go-cache:
	go clean -cache

clean-test-cache:
	go clean -testcache

# Docker compose targets
docker-up:
	docker compose up -d

docker-down:
	docker compose down

# Test setup docker targets
compose-setup-up:
	docker compose -f ./test/setup/docker-compose.yml up -d

compose-setup-down:
	docker compose -f ./test/setup/docker-compose.yml down

