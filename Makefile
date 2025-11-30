# ZLauncher Makefile

.PHONY: all build build-android run clean test lint

# Default target
all: build

# Build for current platform (development)
build:
	go build -o bin/zlauncher ./cmd/zlauncher

# Build for Android
build-android:
	GOOS=android GOARCH=arm64 go build -o bin/zlauncher-android ./cmd/zlauncher

# Run the desktop version (for development)
run: build
	./bin/zlauncher

# Clean build artifacts
clean:
	rm -rf bin/
	go clean

# Run tests
test:
	go test -v ./...

# Run linter
lint:
	go vet ./...
	@if command -v staticcheck > /dev/null; then staticcheck ./...; fi

# Download dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Generate Android APK (requires gogio tool)
apk:
	@echo "Building APK with gogio..."
	@if command -v gogio > /dev/null; then \
		gogio -target android -appid com.zlauncher.app -o bin/zlauncher.apk ./cmd/zlauncher; \
	else \
		echo "gogio not found. Install with: go install gioui.org/cmd/gogio@latest"; \
	fi
