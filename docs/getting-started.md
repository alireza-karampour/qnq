# Getting Started

This guide will help you get up and running with QnQ.

## Prerequisites

- **Go**: Version 1.24.5 or higher
- **Git**: For cloning the repository
- **Terminal**: A modern terminal emulator with Unicode support

## Installation

### From Source

1. **Clone the repository**:
   ```bash
   git clone https://github.com/alireza-karampour/qnq.git
   cd qnq
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Build the application**:
   ```bash
   go build -o qnq
   ```

4. **Run QnQ**:
   ```bash
   ./qnq
   ```

### Using Go Install

```bash
go install github.com/alireza-karampour/qnq@latest
```

## First Run

When you run QnQ for the first time, you'll see an interactive menu:

```
QnQ - Command & Conquer DevOps Center

What would you like to do?

> [ ] Server
  [ ] Client
  [ ] Config
  [ ] Exit

Press q to quit.
Use arrow keys or j/k to navigate, enter/space to select.
```

### Navigation

- **Arrow Keys** or **j/k**: Move up and down
- **Enter** or **Space**: Select/deselect an option
- **q** or **Ctrl+C**: Quit the application

## Configuration

QnQ uses Viper for configuration management. Configuration files are stored in your system's standard config directory.

### Configuration File Location

- **Linux/macOS**: `~/.config/.qnqrc`
- **Windows**: `%APPDATA%\.qnqrc`

### Environment Variables

All configuration options can be overridden using environment variables with the `QNQ_` prefix.

Example:
```bash
export QNQ_SERVER_PORT=8080
export QNQ_DATABASE_HOST=localhost
./qnq
```

Configuration keys with dots are converted to underscores in environment variables:
- Config: `server.port` → Env: `QNQ_SERVER_PORT`
- Config: `database.host` → Env: `QNQ_DATABASE_HOST`

## Basic Usage

### Interactive Mode (Default)

Simply run the application without arguments to launch the interactive TUI:

```bash
./qnq
```

### Using Subcommands

QnQ supports various subcommands for specific tasks:

```bash
# Server command (coming soon)
./qnq server

# View help
./qnq --help

# View subcommand help
./qnq server --help
```

## Quick Tips

1. **Terminal Compatibility**: QnQ works best with modern terminal emulators that support:
   - Unicode characters
   - ANSI color codes
   - Keyboard input handling

2. **Recommended Terminals**:
   - Linux: GNOME Terminal, Konsole, Alacritty, Kitty
   - macOS: iTerm2, Terminal.app, Alacritty
   - Windows: Windows Terminal, ConEmu

3. **Performance**: For the best experience, ensure your terminal is properly configured:
   - Enable true color support if available
   - Use a monospace font with good Unicode coverage
   - Disable unnecessary terminal features that might interfere with TUI

## Troubleshooting

### Build Issues

**Problem**: `go build` fails with dependency errors

**Solution**: 
```bash
go mod tidy
go build -o qnq
```

### Terminal Display Issues

**Problem**: UI appears garbled or incorrectly formatted

**Solution**:
- Ensure your terminal supports ANSI escape codes
- Try a different terminal emulator
- Check that your `TERM` environment variable is set correctly:
  ```bash
  echo $TERM
  # Should output something like: xterm-256color
  ```

### Configuration Not Loading

**Problem**: Environment variables or config file not being read

**Solution**:
- Verify the config file location
- Check environment variable names (must be prefixed with `QNQ_`)
- Ensure proper permissions on the config file

### Application Won't Start

**Problem**: Application exits immediately or shows errors

**Solution**:
- Check Go version: `go version`
- Rebuild the application: `go build -o qnq`
- Run with verbose output: `./qnq -v` (if verbose flag is implemented)

## Next Steps

- Read the [Commands](./commands.md) documentation to learn about available commands
- Check out the [UI Guide](./ui-guide.md) to understand the interface
- See [Development](./development.md) if you want to contribute

## Getting Help

- Check the documentation in the `docs/` directory
- Run `./qnq --help` for command-line help
- Review the source code in the `cmd/` directory

## Uninstallation

To remove QnQ:

1. Delete the binary:
   ```bash
   rm ./qnq
   # or if installed via go install:
   rm $(go env GOPATH)/bin/qnq
   ```

2. (Optional) Remove configuration:
   ```bash
   rm ~/.config/.qnqrc
   ```
