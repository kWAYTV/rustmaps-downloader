# Rustmaps Fetcher

A Go application that fetches and stores Rust game maps data from the RustMaps API.

## Features

- Fetches all available maps from RustMaps API using filter ID
- Handles pagination automatically
- Implements rate limiting to respect API constraints
- Creates backups of existing map data
- Stores maps data in JSON format

## Prerequisites

- Go 1.23.2 or higher
- RustMaps API key
- RustMaps filter ID

## Setup

1. Clone the repository:

```bash
git clone https://github.com/yourusername/rustmaps-downloader.git
cd rustmaps-downloader
```

2. Copy the example environment file and fill in your credentials:

```bash
cp .env.example .env
```

3. Edit `.env` with your API credentials:

```plaintext
RUSTMAPS_API_KEY=your_api_key_here
RUSTMAPS_FILTER_ID=your_filter_id_here
```

4. Make the start script executable (Linux/MacOS only):

```bash
chmod +x scripts/start.sh
```

## Usage

### Using Scripts

#### Windows

Run the start script:

```bash
scripts\start.bat
```

#### Linux/MacOS

Run the start script:

```bash
./scripts/start.sh
```

### Manual Run

Alternatively, you can run the application directly:

```bash
go run main.go
```

## Output

The application creates a `maps` directory and saves the fetched data in JSON format:

- Main output file: `maps/rust_maps_[filter_id].json`
- Backup files (if existing): `maps/rust_maps_[filter_id]_[timestamp].backup.json`

## Troubleshooting

- If you get a permission denied error on Linux/MacOS, make sure you've run the chmod command from step 4
- If you get an API error, verify your credentials in the `.env` file
- Make sure Go is properly installed and in your system PATH

## License

MIT
