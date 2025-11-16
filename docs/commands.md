# Commands

This document describes all available commands in QnQ.

## Root Command

The root command launches the interactive TUI interface.

### Usage

```bash
qnq
```

### Description

Launches the QnQ interactive terminal interface with a menu-driven navigation system.

### Features

- Interactive menu with keyboard navigation
- Configuration initialization
- Visual feedback for selections
- Easy-to-use interface

### Keyboard Controls

| Key | Action |
|-----|--------|
| `↑` / `k` | Move cursor up |
| `↓` / `j` | Move cursor down |
| `Enter` / `Space` | Select/deselect option |
| `q` / `Ctrl+C` | Quit application |

### Menu Options

1. **Server**: Server-related operations (in development)
2. **Client**: Client-related operations (in development)
3. **Config**: Configuration management (in development)
4. **Exit**: Exit the application

### Examples

```bash
# Launch interactive mode
./qnq

# View help
./qnq --help
```

## Subcommands

### server

Server command for managing server-related operations.

**Status**: 🚧 In Development

#### Usage

```bash
qnq server [flags]
```

#### Description

A placeholder command for server-related functionality. This command will be expanded in future versions to include:
- Starting/stopping servers
- Server configuration
- Server status monitoring
- Log management

#### Flags

Currently no flags are defined. This will be expanded as the command is developed.

#### Examples

```bash
# Run server command
./qnq server
```

## Global Flags

Global flags can be used with any command.

### --help, -h

Display help information for any command.

```bash
qnq --help
qnq server --help
```

## Future Commands

The following commands are planned for future releases:

### client

Client-related operations and management.

**Planned Features**:
- Connect to remote servers
- Execute remote commands
- File transfer operations
- Client configuration

### config

Configuration management interface.

**Planned Features**:
- View current configuration
- Edit configuration values
- Validate configuration
- Export/import configuration
- Reset to defaults

### init

Initialize a new QnQ project or configuration.

**Planned Features**:
- Create default configuration
- Setup project structure
- Initialize required directories
- Generate sample files

### status

Display system status and health checks.

**Planned Features**:
- Server status
- Connection status
- Resource usage
- Active operations

### logs

View and manage application logs.

**Planned Features**:
- View logs with filtering
- Export logs
- Log rotation management
- Real-time log streaming

## Command Structure

QnQ follows a hierarchical command structure:

```
qnq
├── server
│   ├── start
│   ├── stop
│   └── status
├── client
│   ├── connect
│   └── disconnect
├── config
│   ├── get
│   ├── set
│   └── list
└── logs
    ├── view
    └── export
```

## Exit Codes

QnQ uses standard exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 2 | Misuse of command |
| 130 | Terminated by Ctrl+C |

## Command Aliases

Currently, no command aliases are defined. This may be added in future versions for commonly used commands.

## Environment Variables

Commands respect the following environment variables:

- `QNQ_*`: Any configuration value can be overridden
- `NO_COLOR`: Disable colored output (standard)
- `TERM`: Terminal type detection

## Best Practices

1. **Use Interactive Mode**: For exploratory tasks, use the interactive TUI
2. **Script with Subcommands**: For automation, use specific subcommands
3. **Check Help**: Always run `--help` to see current options
4. **Environment Variables**: Use env vars for CI/CD and automation

## Examples

### Basic Usage

```bash
# Interactive mode
./qnq

# Get help
./qnq --help

# Run specific command
./qnq server
```

### With Environment Variables

```bash
# Override configuration
QNQ_SERVER_PORT=8080 ./qnq server

# Multiple overrides
QNQ_SERVER_HOST=0.0.0.0 QNQ_SERVER_PORT=8080 ./qnq server
```

### Scripting

```bash
#!/bin/bash
# Example automation script

# Check if qnq is available
if ! command -v qnq &> /dev/null; then
    echo "qnq not found"
    exit 1
fi

# Run qnq commands
qnq server start
qnq status
```

## Command Development

For information on developing new commands, see the [Development Guide](./development.md).

### Adding a New Command

1. Create a new file in `cmd/` directory
2. Define the command using Cobra
3. Register it with the root command in `init()`
4. Implement the command logic in `RunE`
5. Add tests
6. Update documentation

## Notes

- Commands are implemented using the Cobra framework
- All commands support `--help` flag
- Command structure is designed to be extensible
- Future versions will add more sophisticated command options
