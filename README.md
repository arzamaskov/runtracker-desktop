# Runtracker

Local-first desktop application for running training data.

Built with Go, Wails and Svelte. SQLite is planned for local storage.

## Requirements

- Go 1.26+
- Node.js 24+
- pnpm 11+
- Wails 2.15

Install Wails:

```sh
go install github.com/wailsapp/wails/v2/cmd/wails@v2.15.0
```

### macOS

```sh
xcode-select --install
```

### Ubuntu / WSL

```sh
sudo apt install build-essential libgtk-3-dev libwebkit2gtk-4.1-dev pkg-config
```

## Setup

```sh
git clone git@github.com:arzamaskov/runtracker-desktop.git
cd runtracker-desktop
make setup
```

## Development

```sh
make dev
```

## Build

```sh
make build
```

## License

MIT
