#!/bin/bash
echo "Building RustMaps CLI..."
go build -o rustmaps cmd/rustmaps/main.go
if [ $? -eq 0 ]; then
    echo "Build successful! You can now use './rustmaps'"
    chmod +x rustmaps
else
    echo "Build failed! Please check the errors above"
fi 