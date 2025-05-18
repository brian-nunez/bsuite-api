#!/bin/bash

set -e

echo "🚀 Building binaries..."

# Linux (amd64)
echo "🔨 Building Linux amd64 binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o binaries/bsuite-api-linux ./cmd/main.go

# Mac (Intel x86_64)
echo "🔨 Building Mac Intel (amd64) binary..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o binaries/bsuite-api-mac-intel ./cmd/main.go

# Mac (Apple Silicon ARM64)
echo "🔨 Building Mac ARM64 (M1/M2/M3) binary..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o binaries/bsuite-api-mac-arm64 ./cmd/main.go

echo "✅ All binaries built successfully!"
