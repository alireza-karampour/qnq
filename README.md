# QnQ - Command & Conquer DevOps Center

[![Go Version](https://img.shields.io/badge/Go-1.24.5-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-See%20LICENSE-green.svg)](./LICENSE)

> A simple Command & Conquer center for all your DevOps needs

QnQ is an interactive terminal-based application built with Go that provides a beautiful, keyboard-driven interface for managing DevOps tasks.

## ✨ Features

- 🎨 **Interactive TUI** - Beautiful full-screen terminal interface powered by Bubble Tea
- 🌈 **Stylish Design** - Color-coded elements with Lip Gloss styling
- ⚙️ **Configuration Management** - Flexible config with Viper
- ⌨️ **Keyboard-Driven** - Efficient navigation with vim-style keys
- 📱 **Responsive** - Adapts to terminal size and handles resizing
- 🔧 **Extensible** - Easy to add new commands and features
- 🚀 **Fast & Lightweight** - Written in Go for maximum performance

## 🚀 Quick Start

### Installation

```bash
# Clone the repository
git clone https://github.com/alireza-karampour/qnq.git
cd qnq

# Build the application
go build -o qnq

# Run QnQ
./qnq
```

### Using Go Install

```bash
go install github.com/alireza-karampour/qnq@latest
```

## 📖 Documentation

Comprehensive documentation is available in the [`docs/`](./docs) directory:

- **[Getting Started](./docs/getting-started.md)** - Installation and basic usage
- **[Commands Reference](./docs/commands.md)** - All available commands
- **[Architecture](./docs/architecture.md)** - System design and architecture
- **[UI Guide](./docs/ui-guide.md)** - Bubble Tea TUI implementation
- **[Development Guide](./docs/development.md)** - Contributing and development
- **[Changelog](./docs/CHANGELOG.md)** - Version history and changes

## 🎮 Usage

### Interactive Mode

Simply run `qnq` to launch the full-screen interactive interface:

```bash
./qnq
```

The application takes over your entire terminal with a beautiful interface:

```
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
 QnQ - Command & Conquer DevOps Center
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  What would you like to do?

  ▶ [ ] Server
    [ ] Client
    [ ] Config
    [ ] Exit

  ↑/k: up • ↓/j: down • enter/space: select • q/ctrl+c: quit
```

Features:
- **Full-screen mode** with alternate screen buffer
- **Purple-themed** title bar and cursor
- **Green checkmarks** for selected items
- **Centered content** that adapts to your terminal size
- **Preserved history** - your terminal is restored on exit

### Keyboard Controls

| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Enter` / `Space` | Select option |
| `q` / `Ctrl+C` | Quit |

## 🛠️ Technology Stack

- **[Cobra](https://github.com/spf13/cobra)** - CLI framework
- **[Viper](https://github.com/spf13/viper)** - Configuration management
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** - Terminal UI framework

## 📦 Project Structure

```
qnq/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command with TUI
│   └── server.go          # Server subcommand
├── docs/                  # Documentation
├── main.go                # Application entry point
├── go.mod                 # Go module definition
└── README.md              # This file
```

## 🔧 Configuration

QnQ uses Viper for configuration management. Configuration files are stored in:

- **Linux/macOS**: `~/.config/.qnqrc`
- **Windows**: `%APPDATA%\.qnqrc`

### Environment Variables

All configuration options can be overridden using environment variables with the `QNQ_` prefix:

```bash
export QNQ_SERVER_PORT=8080
export QNQ_DATABASE_HOST=localhost
./qnq
```

## 🚧 Development Status

QnQ is currently in active development. The following features are available:

- ✅ Interactive TUI with Bubble Tea
- ✅ Configuration management
- ✅ Command framework
- 🚧 Server command (in development)
- 🚧 Client command (planned)
- 🚧 Config command (planned)

## 🤝 Contributing

Contributions are welcome! Please see the [Development Guide](./docs/development.md) for details on:

- Setting up your development environment
- Code style and conventions
- Testing guidelines
- Pull request process

### Quick Contribution Guide

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Write tests
5. Update documentation
6. Commit your changes (`git commit -m 'feat: add amazing feature'`)
7. Push to the branch (`git push origin feature/amazing-feature`)
8. Open a Pull Request

## 📝 License

See the [LICENSE](./LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) - For the amazing Bubble Tea framework
- [Cobra](https://cobra.dev/) - For the excellent CLI framework
- [Viper](https://github.com/spf13/viper) - For configuration management

## 📬 Contact

- **Repository**: [github.com/alireza-karampour/qnq](https://github.com/alireza-karampour/qnq)
- **Issues**: [GitHub Issues](https://github.com/alireza-karampour/qnq/issues)

## 🗺️ Roadmap

### v0.1.0 (Next Release)
- Complete server command implementation
- Add client command
- Implement Lip Gloss styling
- Add comprehensive tests

### v0.2.0
- Add logging functionality
- Implement status command
- Plugin system
- Enhanced error handling

### v1.0.0
- Production-ready release
- Complete feature set
- Full test coverage
- Performance optimization

---

**Built with ❤️ using Go and Bubble Tea**
