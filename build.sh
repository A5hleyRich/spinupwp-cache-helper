#!/usr/bin/env bash

go mod tidy

echo "Building for AMD64 Ubuntu..."
GOOS=linux GOARCH=amd64 go build -o builds/cache-amd64
echo "Done"

echo "Building for ARM64 Ubuntu..."
GOOS=linux GOARCH=arm64 go build -o builds/cache-arm64
echo "Done"
