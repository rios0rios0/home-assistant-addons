# Repository Instructions for GitHub Copilot

This document provides guidance for GitHub Copilot when working on this repository.

## Repository Overview

This is a **Home Assistant add-ons repository** that provides modular, containerized services for Home Assistant. The repository currently contains two add-ons:

1. **MCP Server Extended** - Extended Model Context Protocol server functionality
2. **Ollama Container** - Ollama integration for local AI capabilities

## Repository Structure

```
home-assistant-addons/
├── .github/               # GitHub-specific files and configurations
├── mcp-server-extended/   # MCP Server Extended add-on
│   ├── config.json       # Add-on configuration (Home Assistant format)
│   └── README.md         # Add-on documentation
├── ollama-container/      # Ollama Container add-on
│   ├── config.json       # Add-on configuration (Home Assistant format)
│   └── README.md         # Add-on documentation
├── repository.json        # Repository metadata (Home Assistant format)
├── README.md             # Main repository documentation
└── LICENSE               # Project license
```

## Home Assistant Add-on Conventions

### Configuration Files (`config.json`)

Each add-on directory must contain a `config.json` file following the Home Assistant add-on specification:

- **Required fields**: `name`, `version`, `slug`, `description`, `arch`, `url`, `startup`
- **Optional fields**: `boot`, `options`, `schema`, `ports`, `environment`, etc.
- **Version format**: Semantic versioning (e.g., "1.0.0")
- **Supported architectures**: `amd64`, `aarch64`, `armhf`, `armv7`
- **Startup types**: `application`, `services`, `system`, `once`

### Add-on Documentation (`README.md`)

Each add-on should have comprehensive documentation including:

- Version and architecture badges
- About section describing the add-on
- Features list
- Installation instructions
- Configuration examples
- Support information
- License reference

### Repository Metadata (`repository.json`)

The root `repository.json` file contains:

- Repository name
- Repository URL
- Maintainer information

## Coding Standards and Best Practices

### General Guidelines

1. **Minimal changes**: Make the smallest possible changes to achieve the goal
2. **Consistency**: Follow existing patterns and conventions in the repository
3. **Documentation**: Update documentation when making changes to add-ons or features
4. **JSON formatting**: Use 2-space indentation for all JSON files
5. **Markdown formatting**: Follow existing markdown style with proper headers and lists

### Version Management

- Use semantic versioning (MAJOR.MINOR.PATCH)
- Increment PATCH for bug fixes
- Increment MINOR for new features (backward compatible)
- Increment MAJOR for breaking changes
- Update version in both `config.json` and `README.md` badges when releasing

### Add-on Development

When creating or modifying add-ons:

1. Ensure `config.json` is valid JSON and follows Home Assistant schema
2. Test configuration on supported architectures (when possible)
3. Update README.md with any new features or configuration options
4. Maintain consistency with existing add-on structure
5. Include clear installation and configuration instructions

### Documentation

- Keep documentation clear, concise, and up-to-date
- Use badges for visual representation of versions and architecture support
- Include practical examples in configuration sections
- Link to GitHub issues for support

## Build and Test

Currently, this repository does not have automated build or test infrastructure. When making changes:

1. Validate JSON files for syntax correctness
2. Review documentation for accuracy
3. Ensure links are working and point to correct locations
4. Verify that changes maintain consistency with Home Assistant add-on conventions

### JSON Validation

To validate JSON files:
```bash
# Using python3
python3 -m json.tool config.json

# Using jq
jq . config.json
```

## Common Tasks

### Adding a New Add-on

1. Create a new directory with a descriptive slug name (lowercase, hyphens)
2. Create `config.json` with required fields
3. Create `README.md` with comprehensive documentation
4. Add entry to main repository `README.md`
5. Ensure consistency with existing add-ons

### Updating an Add-on

1. Update version in `config.json`
2. Update version badge in `README.md`
3. Document changes in README if applicable
4. Ensure backward compatibility unless major version bump

### Documentation Updates

1. Keep formatting consistent with existing style
2. Update all relevant files (add-on README, main README)
3. Verify all links are functional
4. Maintain badge consistency

## Support and Contribution

- **Issues**: Report bugs and feature requests via GitHub Issues
- **Pull Requests**: Welcome for bug fixes, features, and documentation improvements
- **Maintainer**: @rios0rios0
- **License**: See LICENSE file in repository root

## Important Notes for Copilot

1. This is a Home Assistant add-ons repository - follow Home Assistant conventions strictly
2. Do not modify working functionality unless necessary or fixing security issues
3. Maintain multi-architecture support (amd64, aarch64, armhf, armv7)
4. Keep changes minimal and focused on the specific task
5. Validate JSON syntax before committing
6. Update documentation to reflect any configuration or feature changes
7. Respect the existing structure and patterns in the repository
8. Do not introduce new dependencies or build tools without explicit need
