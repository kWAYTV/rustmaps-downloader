#!/bin/bash

# Ensure the script fails on any error
set -e

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    exit 1
fi

# Check if .env file exists
if [ ! -f .env ]; then
    echo "Error: .env file not found"
    echo "Please copy .env.example to .env and configure your credentials"
    exit 1
fi

# Run the application
go run main.go 