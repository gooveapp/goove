# Goove

A personal vinyl record collection manager built with Go. Catalog, organize, and manage your vinyl records with Discogs API integration.

## Features

- Browse and manage your vinyl record collection
- Add new records with detailed metadata
- Discogs API integration for fetching record information
- User preferences for currency, date format, and condition defaults
- Dark mode interface
- SQLite database for local storage

## Installation

### Option 1: Download Binary

Download the latest release binary for your platform from the [Releases](https://github.com/gooveapp/goove/releases) page.

```bash
# Make the binary executable (Linux/macOS)
chmod +x goove

# Run the application
./goove
```

### Option 2: Docker

Pull and run the Docker image:

```bash
docker pull ghcr.io/gooveapp/goove:latest
docker run -p 3000:3000 -v goove-data:/data ghcr.io/gooveapp/goove:latest
```

Or use Docker Compose:

```bash
docker-compose up
```

## Configuration

Create a `.env` file or set environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `3000` |
| `ENV` | Environment (development/production) | `development` |
| `DISCOGS_TOKEN` | Your Discogs API token | Required |
| `DISCOGS_USER_AGENT` | User agent for Discogs API | `Goove/1.0` |
| `DATABASE_PATH` | Path to SQLite database | `/data/goove.db` |
| `LOG_LEVEL` | Logging level | `info` |

### Getting a Discogs Token

1. Create a Discogs account at [discogs.com](https://www.discogs.com/)
2. Go to [Developer Settings](https://www.discogs.com/settings/developers)
3. Generate a personal access token

## Usage

Once running, access the application at `http://localhost:3000`.

## Tech Stack

- **Backend:** Go, Echo, SQLite
- **Frontend:** Templ, Tailwind CSS, HTMX

## Contributing

Contributions are welcome! Please read our [Contributing Guide](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
