#!/bin/bash

# Exit if anything fails
set -e

# Find the first .go file (can be improved for specific patterns like main.go)
GOFILE=$(find . -name '*.go' | grep -m 1 'main.go')

if [ -z "$GOFILE" ]; then
    echo "Error: No main.go file found in the project."
    exit 1
fi

# Get directory of the main file
DIR=$(dirname "$GOFILE")
cd "$DIR"

# Name of output binary
BINARY_NAME="forum"

echo "Building Go project from $GOFILE..."
go build -o "$BINARY_NAME"
chmod +x "$BINARY_NAME"

echo "Running $BINARY_NAME..."
./"$BINARY_NAME"
