@echo off
echo Building RustMaps CLI...
cd ..
go build -o rustmaps.exe ./cmd/rustmaps
if %ERRORLEVEL% EQU 0 (
    echo Build successful! You can now use 'rustmaps.exe'
) else (
    echo Build failed! Please check the errors above
)
cd scripts