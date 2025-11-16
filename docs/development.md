# Development Guide

This guide is for developers who want to contribute to or extend QnQ.

## Development Setup

### Prerequisites

- **Go**: Version 1.24.5 or higher
- **Git**: For version control
- **Make**: (Optional) For build automation
- **golangci-lint**: (Optional) For code linting

### Getting the Code

```bash
# Clone the repository
git clone https://github.com/alireza-karampour/qnq.git
cd qnq

# Install dependencies
go mod download

# Build the project
go build -o qnq
```

## Project Structure

```
qnq/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command with TUI
│   └── server.go          # Server subcommand
├── docs/                  # Documentation (this directory)
│   ├── README.md          # Documentation index
│   ├── architecture.md    # Architecture overview
│   ├── commands.md        # Command reference
│   ├── development.md     # This file
│   ├── getting-started.md # User guide
│   └── ui-guide.md        # UI implementation details
├── main.go                # Application entry point
├── go.mod                 # Go module definition
├── go.sum                 # Dependency checksums
└── LICENSE                # License file
```

## Code Organization

### Adding a New Command

1. **Create a new file** in the `cmd/` directory:

```go
// cmd/mycommand.go
package cmd

import (
    "github.com/spf13/cobra"
)

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Brief description",
    Long:  `Detailed description of what this command does.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        // Implementation here
        return nil
    },
}

func init() {
    rootCmd.AddCommand(myCmd)
    
    // Add flags
    myCmd.Flags().StringP("flag", "f", "", "Flag description")
}
```

2. **Build and test**:

```bash
go build -o qnq
./qnq mycommand --help
```

### Extending the TUI

To add new views or modify the UI:

1. **Add new state** to the model:

```go
type model struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
    // Add new fields
    currentView string
    inputValue  string
}
```

2. **Update the Update function**:

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        // Handle new key bindings
        switch msg.String() {
        case "n":
            m.currentView = "new-view"
            return m, nil
        }
    }
    return m, nil
}
```

3. **Update the View function**:

```go
func (m model) View() string {
    switch m.currentView {
    case "new-view":
        return renderNewView(m)
    default:
        return renderMainView(m)
    }
}
```

## Building

### Standard Build

```bash
go build -o qnq
```

### Build with Version Info

```bash
VERSION=$(git describe --tags --always)
go build -ldflags "-X main.Version=$VERSION" -o qnq
```

### Cross-Platform Builds

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o qnq-linux-amd64

# macOS
GOOS=darwin GOARCH=amd64 go build -o qnq-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o qnq-darwin-arm64

# Windows
GOOS=windows GOARCH=amd64 go build -o qnq-windows-amd64.exe
```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Writing Tests

#### Testing Commands

```go
// cmd/mycommand_test.go
package cmd

import (
    "testing"
)

func TestMyCommand(t *testing.T) {
    // Test implementation
}
```

#### Testing the TUI

```go
// cmd/root_test.go
package cmd

import (
    "testing"
    tea "github.com/charmbracelet/bubbletea"
)

func TestModelUpdate(t *testing.T) {
    m := initialModel()
    
    // Test cursor movement
    msg := tea.KeyMsg{Type: tea.KeyDown}
    newModel, _ := m.Update(msg)
    
    if newModel.(model).cursor != 1 {
        t.Errorf("expected cursor at 1, got %d", newModel.(model).cursor)
    }
}

func TestModelView(t *testing.T) {
    m := initialModel()
    view := m.View()
    
    if view == "" {
        t.Error("view should not be empty")
    }
}
```

## Code Style

### Go Formatting

```bash
# Format code
go fmt ./...

# Check formatting
gofmt -l .
```

### Linting

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run
```

### Code Conventions

1. **Naming**:
   - Use camelCase for private functions/variables
   - Use PascalCase for exported functions/variables
   - Use descriptive names

2. **Comments**:
   - Document all exported functions
   - Use complete sentences
   - Start with the function name

```go
// ProcessData processes the input data and returns the result.
// It returns an error if the data is invalid.
func ProcessData(data string) (string, error) {
    // Implementation
}
```

3. **Error Handling**:
   - Always check errors
   - Wrap errors with context
   - Use meaningful error messages

```go
if err != nil {
    return fmt.Errorf("failed to process data: %w", err)
}
```

## Dependencies

### Managing Dependencies

```bash
# Add a new dependency
go get github.com/example/package

# Update dependencies
go get -u ./...

# Tidy dependencies
go mod tidy

# Vendor dependencies (optional)
go mod vendor
```

### Current Dependencies

- **cobra**: CLI framework
- **viper**: Configuration management
- **bubbletea**: TUI framework
- **lipgloss**: Terminal styling (indirect)

## Debugging

### Debug Logging

Since Bubble Tea takes over the terminal, log to a file:

```go
import (
    "log"
    "os"
)

func init() {
    f, err := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    log.SetOutput(f)
}

// In your code
log.Printf("Debug: cursor=%d", m.cursor)
```

### Using Delve Debugger

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug the application
dlv debug

# Set breakpoint
(dlv) break main.main
(dlv) continue
```

## Git Workflow

### Branching Strategy

- `main`: Stable, production-ready code
- `develop`: Integration branch for features
- `feature/*`: Feature branches
- `bugfix/*`: Bug fix branches

### Commit Messages

Follow conventional commits:

```
type(scope): subject

body

footer
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting)
- `refactor`: Code refactoring
- `test`: Adding tests
- `chore`: Maintenance tasks

Example:
```
feat(ui): add color support to TUI

Implement Lip Gloss styling for better visual appearance.
Add color configuration options.

Closes #123
```

### Pull Request Process

1. Create a feature branch
2. Make your changes
3. Write tests
4. Update documentation
5. Submit pull request
6. Address review comments
7. Merge after approval

## Documentation

### Updating Documentation

When making changes, update relevant documentation:

1. **Code Changes**: Update inline comments
2. **New Features**: Update `docs/commands.md` and `docs/getting-started.md`
3. **Architecture Changes**: Update `docs/architecture.md`
4. **UI Changes**: Update `docs/ui-guide.md`

### Documentation Standards

- Use clear, concise language
- Include code examples
- Add diagrams where helpful
- Keep documentation in sync with code

## Release Process

### Version Numbering

Follow semantic versioning (SemVer):
- `MAJOR.MINOR.PATCH`
- `MAJOR`: Breaking changes
- `MINOR`: New features (backward compatible)
- `PATCH`: Bug fixes

### Creating a Release

1. **Update version**:
   ```bash
   git tag -a v1.0.0 -m "Release version 1.0.0"
   ```

2. **Build binaries**:
   ```bash
   ./scripts/build-release.sh  # If available
   ```

3. **Push tag**:
   ```bash
   git push origin v1.0.0
   ```

4. **Create GitHub release**:
   - Go to GitHub releases
   - Create new release from tag
   - Add release notes
   - Upload binaries

## Performance Optimization

### Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof

# Memory profiling
go test -memprofile=mem.prof -bench=.
go tool pprof mem.prof
```

### Optimization Tips

1. **Avoid allocations** in hot paths
2. **Use string builders** for string concatenation
3. **Profile before optimizing**
4. **Keep the TUI responsive** (avoid blocking operations)

## Common Development Tasks

### Adding Configuration Options

1. **Define in Viper**:
```go
viper.SetDefault("myoption", "default-value")
```

2. **Access in code**:
```go
value := viper.GetString("myoption")
```

3. **Document** in `docs/getting-started.md`

### Adding a Subcommand

1. Create `cmd/subcommand.go`
2. Register with root command
3. Add tests
4. Update documentation

### Modifying the TUI

1. Update model structure
2. Modify Update function
3. Update View function
4. Test thoroughly
5. Update `docs/ui-guide.md`

## Troubleshooting

### Common Issues

**Issue**: Import cycle detected

**Solution**: Restructure packages to avoid circular dependencies

---

**Issue**: Tests fail in CI but pass locally

**Solution**: Check for environment-specific code, ensure deterministic tests

---

**Issue**: TUI rendering issues

**Solution**: Test in different terminals, check ANSI code support

## Contributing

### Before Contributing

1. Check existing issues and PRs
2. Discuss major changes in an issue first
3. Follow the code style guide
4. Write tests for new features
5. Update documentation

### Code Review Checklist

- [ ] Code follows style guide
- [ ] Tests added/updated
- [ ] Documentation updated
- [ ] No unnecessary dependencies
- [ ] Error handling is appropriate
- [ ] Performance considerations addressed

## Resources

### Go Resources

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

### Framework Documentation

- [Cobra](https://github.com/spf13/cobra)
- [Viper](https://github.com/spf13/viper)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)

### Tools

- [golangci-lint](https://golangci-lint.run/)
- [Delve Debugger](https://github.com/go-delve/delve)
- [Go Modules](https://golang.org/ref/mod)

## Getting Help

- Read the documentation in `docs/`
- Check existing issues on GitHub
- Ask questions in discussions
- Review the source code

## License

See the LICENSE file in the root directory for licensing information.
