# RustMaps CLI

A command-line interface tool for interacting with the RustMaps API, built in Go.

[![Build Status](https://github.com/kWAYTV/rustmaps-downloader/actions/workflows/build.yml/badge.svg)](https://github.com/kWAYTV/rustmaps-downloader/actions/workflows/build.yml)
[![codecov](https://codecov.io/gh/kWAYTV/rustmaps-downloader/branch/main/graph/badge.svg)](https://codecov.io/gh/kWAYTV/rustmaps-downloader)
[![Go Report Card](https://goreportcard.com/badge/github.com/kWAYTV/rustmaps-downloader)](https://goreportcard.com/report/github.com/kWAYTV/rustmaps-downloader)
[![Latest Release](https://img.shields.io/github/v/release/kWAYTV/rustmaps-downloader?include_prereleases)](https://github.com/kWAYTV/rustmaps-downloader/releases/latest)

## Features

- CLI-based interface for easy interaction
- Downloads maps based on filter ID
- Updates server config files with map seeds
- Handles pagination automatically
- Implements rate limiting to respect API constraints
- Creates backups of existing map data
- Stores maps data in JSON format
- Preserves YAML formatting and comments

## Prerequisites

- Go 1.21 or higher
- RustMaps API key
- RustMaps filter ID

## Installation

### Pre-built binaries

Download the latest release for your platform from the [Releases page](https://github.com/kWAYTV/rustmaps-downloader/releases).

Available platforms:

- Windows (32/64-bit)
- Linux (32/64-bit)
- macOS (Intel/Apple Silicon)

### Building from source

1. Clone the repository:

```bash
git clone https://github.com/kWAYTV/rustmaps-downloader.git
cd rustmaps-downloader
```

2. Make scripts executable (Linux only):

```bash
chmod +x scripts/*.sh
```

3. Copy the example environment file and fill in your credentials:

```bash
cp .env.example .env
```

4. Edit `.env` with your API credentials:

```plaintext
RUSTMAPS_API_KEY=your_api_key_here
RUSTMAPS_FILTER_ID=your_filter_id_here
```

## Usage

### Using Scripts

#### Windows

```bash
# Build the CLI
scripts\build.bat

# Show help
scripts\help.bat
```

#### Linux

```bash
# Build the CLI
./scripts/build.sh

# Show help
./scripts/help.sh
```

### Using the Built CLI

After building with the build script:

1. Get help and see available commands:

```bash
# Windows
rustmaps.exe --help

# Linux
./rustmaps --help
```

2. Download maps using your filter:

```bash
# Windows
rustmaps.exe download

# Linux
./rustmaps download
```

3. Update server config with map seeds:

```bash
# Windows
rustmaps.exe update-config maps/rust_maps_[filter_id].json config.yml

# Linux
./rustmaps update-config maps/rust_maps_[filter_id].json config.yml
```

4. Get help for a specific command:

```bash
./rustmaps [command] --help
```

## Output

The application creates a `maps` directory and saves the fetched data in JSON format. Files are named with an incrementing counter to avoid overwriting:

- First run: `maps/rust_maps_[filter_id].json`
- Second run: `maps/rust_maps_[filter_id]_1.json`
- Third run: `maps/rust_maps_[filter_id]_2.json`
  And so on...

When using the `update-config` command, it will update the specified YAML config file while preserving all comments and formatting, only modifying the `world_seeds` section with the seeds from your maps JSON.

## Project Structure

```
rustmaps-downloader/
├── cmd/
│   └── rustmaps/
│       ├── main.go              # CLI entry point
│       └── commands/
│           ├── root.go          # Root command definition
│           ├── download.go      # Download command implementation
│           ├── update-config.go # Config update command implementation
│           └── version.go       # Version command implementation
├── .env.example                 # Example environment file
├── .env                         # Your environment file (git-ignored)
└── README.md                    # This file
```

## Adding New Commands

To add new functionality:

1. Create a new file in `cmd/rustmaps/commands/`
2. Define your command using cobra
3. Register it in the `init()` function

Example structure for a new command:

```go
package commands

import "github.com/spf13/cobra"

var newCmd = &cobra.Command{
    Use:   "commandname",
    Short: "Short description",
    Long:  `Longer description`,
    Run: func(cmd *cobra.Command, args []string) {
        // Command implementation
    },
}

func init() {
    RootCmd.AddCommand(newCmd)
}
```

## Troubleshooting

- If you get API errors, verify your credentials in the `.env` file
- Make sure Go is properly installed and in your system PATH
- For dependency issues, run `go mod tidy`

## Shell Completion

The CLI supports shell autocompletion. To enable it:

#### Bash

```bash
# Linux
./rustmaps completion bash > /etc/bash_completion.d/rustmaps
# or locally
./rustmaps completion bash > ~/.rustmaps-completion.bash
source ~/.rustmaps-completion.bash
```

#### PowerShell

```powershell
# Add to your PowerShell profile
rustmaps.exe completion powershell | Out-String | Invoke-Expression
# or temporarily
rustmaps.exe completion powershell > rustmaps.ps1
. .\rustmaps.ps1
```

#### Zsh

```bash
# Add to your zshrc
./rustmaps completion zsh > "${fpath[1]}/_rustmaps"
# or locally
./rustmaps completion zsh > ~/.rustmaps-completion.zsh
source ~/.rustmaps-completion.zsh
```

#### Fish

```bash
./rustmaps completion fish > ~/.config/fish/completions/rustmaps.fish
```

## License

[MIT](LICENSE)
