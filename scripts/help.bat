@echo off
if exist rustmaps.exe (
    rustmaps.exe --help
) else (
    go run cmd/rustmaps/main.go --help
) 