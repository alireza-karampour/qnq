# Changelog

All notable changes to the QnQ project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Interactive TUI using Bubble Tea framework
- Root command with menu-driven interface
- Configuration management with Viper
- Server subcommand (placeholder)
- Comprehensive documentation in `docs/` directory
  - Architecture documentation
  - Getting started guide
  - Commands reference
  - UI implementation guide
  - Development guide

### Changed
- Root command now launches interactive TUI instead of just initializing config

### Technical Details
- Integrated Bubble Tea v1.3.10 for terminal UI
- Implemented Model-View-Update (MVU) pattern
- Added keyboard navigation (arrow keys, vim-style j/k)
- Added selection mechanism with visual feedback

## [0.0.1] - Initial Release

### Added
- Initial project setup with Go modules
- Cobra CLI framework integration (v1.10.1)
- Viper configuration management (v1.21.0)
- Basic project structure
- License file

### Technical Details
- Go version: 1.24.5
- Module: github.com/alireza-karampour/qnq
- Basic command structure with root command

---

## Version History

### Version Format

Versions follow semantic versioning: `MAJOR.MINOR.PATCH`

- **MAJOR**: Incompatible API changes
- **MINOR**: Backward-compatible functionality additions
- **PATCH**: Backward-compatible bug fixes

### Release Notes

#### Unreleased (Current Development)

The current development version includes a complete rewrite of the root command to use an interactive TUI. This provides a much better user experience with visual feedback and keyboard-driven navigation.

**Key Features:**
- Menu-based navigation
- Visual cursor and selection indicators
- Multiple quit options (q, Ctrl+C)
- Vim-style navigation support
- Extensible architecture for future views

**Breaking Changes:**
- None (first major release)

**Migration Guide:**
- No migration needed for new project

#### 0.0.1 (Initial)

Initial project setup with basic CLI framework. This version established the foundation with Cobra and Viper integration.

---

## Future Releases

### Planned for v0.1.0
- [ ] Complete server command implementation
- [ ] Add client command
- [ ] Add config command for interactive configuration
- [ ] Implement Lip Gloss styling for better visuals
- [ ] Add unit tests for all commands
- [ ] Add integration tests for TUI

### Planned for v0.2.0
- [ ] Add logging functionality
- [ ] Implement status command
- [ ] Add support for plugins
- [ ] Enhanced error handling and reporting
- [ ] Configuration file templates

### Planned for v1.0.0
- [ ] Complete feature set
- [ ] Comprehensive test coverage (>80%)
- [ ] Full documentation
- [ ] Production-ready stability
- [ ] Performance optimization
- [ ] Security audit

---

## Contributing

When contributing, please:
1. Update this changelog with your changes
2. Follow the format specified above
3. Place changes under the `[Unreleased]` section
4. Use appropriate categories (Added, Changed, Deprecated, Removed, Fixed, Security)

### Changelog Categories

- **Added**: New features
- **Changed**: Changes in existing functionality
- **Deprecated**: Soon-to-be removed features
- **Removed**: Removed features
- **Fixed**: Bug fixes
- **Security**: Security vulnerability fixes

---

[Unreleased]: https://github.com/alireza-karampour/qnq/compare/v0.0.1...HEAD
[0.0.1]: https://github.com/alireza-karampour/qnq/releases/tag/v0.0.1
