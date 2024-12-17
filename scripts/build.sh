#!/bin/bash

echo "Building RustMaps CLI..."
cd ..
go build -o rustmaps ./cmd/rustmaps

if [ $? -eq 0 ]; then
    echo "Build successful! You can now use './rustmaps'"
else
    echo "Build failed! Please check the errors above"
fi
cd scripts