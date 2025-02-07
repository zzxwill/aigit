BINARY_NAME=aigit
VERSION=$(shell git describe --tags --always --dirty)

# Build directories
BUILD_DIR=build
MACOS_AMD64=$(BUILD_DIR)/$(BINARY_NAME)_darwin_amd64_$(VERSION)
MACOS_ARM64=$(BUILD_DIR)/$(BINARY_NAME)_darwin_arm64_$(VERSION)
WINDOWS_AMD64=$(BUILD_DIR)/$(BINARY_NAME)_windows_amd64_$(VERSION).exe

.PHONY: all clean build-all build-macos-amd64 build-macos-arm64 build-windows

all: build-all

build-all: clean macos-amd64 macos-arm64 windows
	@echo "Build complete! Binaries are in the $(BUILD_DIR) directory"

macos-amd64:
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.Version=$(VERSION)" -o $(MACOS_AMD64) main.go

macos-arm64:
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.Version=$(VERSION)" -o $(MACOS_ARM64) main.go

windows:
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.Version=$(VERSION)" -o $(WINDOWS_AMD64) main.go

clean:
	@rm -rf $(BUILD_DIR)
