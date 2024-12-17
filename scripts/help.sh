#!/bin/bash
if [ -f "./rustmaps" ]; then
    ./rustmaps --help
else
    go run cmd/rustmaps/main.go --help
fi 