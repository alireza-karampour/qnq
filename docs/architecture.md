# Architecture

This document describes the architecture and design decisions of the QnQ project.

## Overview

QnQ is built as a command-line application with a terminal user interface (TUI). The architecture follows a modular design pattern with clear separation of concerns.

## Technology Stack

### Core Libraries

- **[Cobra](https://github.com/spf13/cobra)** (v1.10.1) - Command-line interface framework
- **[Viper](https://github.com/spf13/viper)** (v1.21.0) - Configuration management
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** (v1.3.10) - Terminal UI framework

### Supporting Libraries

- **Lip Gloss** (v1.1.0) - Terminal styling (via Bubble Tea)
- **Various terminal utilities** - For cross-platform terminal support

## Project Structure

```
qnq/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command with TUI
│   └── server.go          # Server subcommand
├── docs/                  # Documentation
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── go.sum                 # Dependency checksums
```

## Component Architecture

### 1. Main Entry Point (`main.go`)

The application entry point is minimal and delegates to the command structure:

```
main() → cmd.Execute()
```

### 2. Command Layer (`cmd/`)

#### Root Command (`root.go`)

The root command serves as the main entry point for the application and includes:

- **Configuration Initialization**: Sets up Viper for configuration management
- **TUI Launch**: Initializes and runs the Bubble Tea interface
- **Model-View-Update Pattern**: Implements the Elm architecture via Bubble Tea

**Key Components:**

- `model` struct: Holds the application state
  - `choices`: Available menu options
  - `cursor`: Current cursor position
  - `selected`: Map of selected items

- `Init()`: Initializes the model (no commands needed at startup)
- `Update()`: Handles user input and state transitions
- `View()`: Renders the UI based on current state

#### Subcommands

- `server.go`: Server-related commands (placeholder for future implementation)

### 3. Configuration Management

Configuration is handled through Viper with the following setup:

- **Config Location**: User's config directory (`~/.config` on Linux/macOS)
- **Config File**: `.qnqrc`
- **Environment Variables**: Prefixed with `QNQ_`
- **Key Mapping**: Dots (`.`) in config keys map to underscores (`_`) in env vars

Example:
```
Config key: database.host
Env var: QNQ_DATABASE_HOST
```

## Design Patterns

### Model-View-Update (MVU)

The Bubble Tea framework implements the Elm architecture:

1. **Model**: Application state
2. **Update**: State transitions based on messages (events)
3. **View**: Pure function that renders the model

```
User Input → Message → Update(model, msg) → New Model → View(model) → Display
```

### Command Pattern

Cobra provides a command pattern implementation where each command:
- Has a specific responsibility
- Can have subcommands
- Supports flags and arguments
- Has a `RunE` function for execution with error handling

## Key Design Decisions

### 1. Why Bubble Tea?

- **Composable**: Build complex UIs from simple components
- **Testable**: Pure functions make testing straightforward
- **Cross-platform**: Works consistently across different terminals
- **Modern**: Follows functional programming principles

### 2. Why Cobra + Viper?

- **Industry Standard**: Widely used in Go CLI applications
- **Feature Rich**: Flags, subcommands, help generation
- **Integration**: Cobra and Viper work seamlessly together
- **Flexibility**: Easy to extend with new commands

### 3. Configuration Strategy

- User-level configuration in standard OS locations
- Environment variable override support
- Consistent naming conventions
- Automatic environment variable binding

## Future Architecture Considerations

### Planned Enhancements

1. **Plugin System**: Allow extending functionality through plugins
2. **API Layer**: Separate business logic from UI
3. **State Management**: More sophisticated state handling for complex workflows
4. **Testing Framework**: Comprehensive test coverage
5. **Logging**: Structured logging with configurable levels

### Scalability

The current architecture supports:
- Adding new commands without modifying existing code
- Extending the TUI with new views and components
- Configuration expansion without breaking changes
- Multiple deployment modes (CLI, server, daemon)

## Dependencies Graph

```
main.go
  └── cmd.Execute()
      └── rootCmd.Execute()
          ├── Viper (config)
          └── Bubble Tea (TUI)
              └── model (MVU pattern)
```

## Security Considerations

1. **Configuration**: Sensitive data should use environment variables
2. **Input Validation**: All user input should be validated
3. **Error Handling**: Errors should not leak sensitive information
4. **Dependencies**: Regular updates to address security vulnerabilities

## Performance

- **Startup Time**: Minimal initialization, fast startup
- **Memory**: Low memory footprint for TUI applications
- **Responsiveness**: Event-driven architecture ensures UI responsiveness
- **Resource Usage**: Efficient terminal rendering with Bubble Tea

## Conclusion

The architecture is designed to be:
- **Simple**: Easy to understand and modify
- **Extensible**: New features can be added without major refactoring
- **Maintainable**: Clear separation of concerns
- **User-Friendly**: Interactive and intuitive interface
