@echo off

REM Check if Go is installed
where go >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo Error: Go is not installed
    exit /b 1
)

REM Check if .env file exists
if not exist .env (
    echo Error: .env file not found
    echo Please copy .env.example to .env and configure your credentials
    exit /b 1
)

REM Run the application
go run main.go 