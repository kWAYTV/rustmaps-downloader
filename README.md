# RustMaps CLI

A command-line interface tool for interacting with the RustMaps API, built in Go.

## Features

- CLI-based interface for easy interaction
- Downloads maps based on filter ID
- Handles pagination automatically
- Implements rate limiting to respect API constraints
- Creates backups of existing map data
- Stores maps data in JSON format

## Prerequisites

- Go 1.21 or higher
- RustMaps API key
- RustMaps filter ID

## Installation

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

# Run in development mode
scripts\dev.bat download  # or any other command
```

#### Linux

```bash
# Build the CLI
./scripts/build.sh

# Show help
./scripts/help.sh

# Run in development mode
./scripts/dev.sh download  # or any other command
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

3. Get help for a specific command:

```bash
./rustmaps download --help
```

### Development Usage

During development, you can run commands directly without building:

```bash
go run cmd/rustmaps/main.go download
```

## Output

The application creates a `maps` directory and saves the fetched data in JSON format. Files are named with an incrementing counter to avoid overwriting:

- First run: `maps/rust_maps_[filter_id].json`
- Second run: `maps/rust_maps_[filter_id]_1.json`
- Third run: `maps/rust_maps_[filter_id]_2.json`
  And so on...

## Project Structure

```
rustmaps-downloader/
├── cmd/
│   └── rustmaps/
│       ├── main.go              # CLI entry point
│       └── commands/
│           ├── root.go          # Root command definition
│           └── download.go      # Download command implementation
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

## License

[MIT](LICENSE)

### Shell Completion

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
