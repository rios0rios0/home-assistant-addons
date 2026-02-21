# Home Assistant Addon Structure

This document describes the structure of the MCP Server Extended Home Assistant addon.

## Directory Structure

```
mcp-server-extended/
├── config.yaml              # Addon configuration and schema
├── build.yaml               # Build configuration for different architectures
├── Dockerfile               # Container build instructions
├── build.sh                 # Build script for manual building
├── rootfs/                  # Filesystem overlay for the container
│   └── etc/
│       └── services.d/
│           └── mcp-server-extended/
│               ├── run      # Service startup script
│               └── finish   # Service shutdown script
├── src/                     # Application source code
│   └── mcp_ha_extended/
│       ├── __init__.py
│       └── server.py
├── pyproject.toml           # Python project configuration
├── pdm.lock                 # Locked dependencies
└── .docs/                   # Documentation (see SUMMARY.md)
```

## Key Files

### config.yaml
Defines the addon metadata, configuration schema, and requirements:
- Addon name, version, description
- Supported architectures (aarch64, armv7, amd64)
- Configuration options (ha_url, ha_token, log_level)
- Home Assistant API requirements

### build.yaml
Specifies the base images for each architecture:
- Base images from Home Assistant's official registry
- Build arguments (bashio, tempio, s6-overlay versions)

### Dockerfile
Multi-stage build that:
1. Sets up the base environment
2. Installs system dependencies (Python, curl, etc.)
3. Installs S6 overlay, bashio, and tempio
4. Copies application code
5. Installs Python dependencies using PDM
6. Sets up the rootfs overlay

### rootfs/etc/services.d/mcp-server-extended/run
Service startup script that:
- Reads configuration from bashio
- Sets environment variables (HA_URL, HA_TOKEN)
- Validates configuration
- Starts the MCP server using PDM

### rootfs/etc/services.d/mcp-server-extended/finish
Service shutdown script for cleanup (if needed)

## Build Process

For detailed build instructions, see the [Addon Build Guide](ADDON_BUILD.md).

1. **Using GitHub Actions** (Recommended): Automatic builds and pushes to GitHub Container Registry on push to main/master or when tags are created
2. **Using build.sh**: Local script for manual building and testing
3. **Using Docker directly**: More control over build parameters

## Installation

For installation instructions, see [Addon Installation](ADDON_INSTALLATION.md).

The addon can be installed:
1. **GitHub repository** (Recommended): Add `https://github.com/rios0rios0/home-assistant-addons` as a repository in Home Assistant
2. **Local development**: Copy to `/config/addons/` and add as local repository for testing

## Configuration

Users configure the addon through Home Assistant's addon UI:
- **ha_url**: Home Assistant instance URL
- **ha_token**: Long-lived access token
- **log_level**: Logging verbosity

## Runtime

The addon runs as a service managed by S6 overlay:
- Automatically starts when the addon is started
- Automatically restarts on failure
- Logs are available through Home Assistant's addon logs interface
- Communicates via stdio using the MCP protocol
