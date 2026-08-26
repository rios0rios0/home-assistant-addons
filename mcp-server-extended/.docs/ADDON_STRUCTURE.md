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
├── cmd/                     # Composition root
│   └── mcp-ha-extended/
│       ├── main.go          # Entry point
│       └── dig.go           # Dependency injection wiring
├── internal/                # Application code
│   ├── container.go         # Provider registration across layers
│   ├── domain/              # Contracts: entity and repository interface
│   │   ├── entities/
│   │   └── repositories/
│   └── infrastructure/      # Implementations
│       ├── controllers/     # MCP tool handlers, responses, mappers
│       └── repositories/    # Home Assistant REST client
├── test/                    # Shared test helpers
│   └── domain/
│       ├── builders/        # Test data builders
│       └── doubles/         # In-memory test doubles
├── go.mod                   # Module path and dependencies
├── go.sum                   # Dependency checksums
├── .golangci.yml            # Linter configuration
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
- Build arguments (bashio, tempio, s6-overlay versions, and `ADDON_VERSION`,
  which is stamped into the binary and reported as the MCP server version)

### Dockerfile
Multi-stage build that:
1. Cross-compiles the binary in a `--platform=${BUILDPLATFORM}` Go builder stage,
   so the target architecture is never emulated
2. Sets up the base environment
3. Installs S6 overlay, bashio, and tempio
4. Copies the static binary to `/usr/bin/mcp-ha-extended`
5. Sets up the rootfs overlay

No language runtime is installed in the final image — the binary is statically
linked and self-contained.

### rootfs/etc/services.d/mcp-server-extended/run
Service startup script that:
- Reads configuration from bashio
- Sets environment variables (HA_URL, HA_TOKEN)
- Validates configuration (both `HA_URL` and `HA_TOKEN` are required)
- Executes `/usr/bin/mcp-ha-extended`

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
