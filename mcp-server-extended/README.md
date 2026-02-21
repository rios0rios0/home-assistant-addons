# MCP Server Extended

A Model Context Protocol (MCP) server that extends Home Assistant capabilities to support automation management. Available as both a standalone Python application and a Home Assistant addon.

## Features

- ✅ List all automations
- ✅ Get automation details
- ✅ Create new automations from YAML
- ✅ Update existing automations
- ✅ Delete automations
- ✅ Trigger automations manually
- ✅ Enable/disable automations

## Quick Start

### As a Home Assistant Addon (Recommended)

1. Add repository: `https://github.com/rios0rios0/home-assistant-addons`
2. Install "MCP Server Extended" from the addon store
3. Configure with your Home Assistant URL and access token
4. Use with any MCP-compatible client (like Cursor IDE)

See [Addon Installation Guide](.docs/ADDON_INSTALLATION.md) for detailed instructions.

### As a Standalone Python Application

1. Install dependencies: `pdm install`
2. Set environment variables: `HA_URL` and `HA_TOKEN`
3. Run: `pdm run python -m mcp_ha_extended.server`
4. Configure your MCP client to use the server

See [Quick Start Guide](.docs/QUICK_START.md) for a 5-minute setup, or [Setup Guide](.docs/SETUP.md) for detailed instructions.

## Documentation

All documentation is organized in the [`.docs`](.docs/) folder:

- **[Documentation Index](.docs/SUMMARY.md)** - Start here for navigation
- **[Quick Start](.docs/QUICK_START.md)** - Get running in 5 minutes
- **[Addon Installation](.docs/ADDON_INSTALLATION.md)** - Install as Home Assistant addon
- **[Setup Guide](.docs/SETUP.md)** - Detailed development setup
- **[Usage Examples](.docs/USAGE_EXAMPLES.md)** - Code examples and use cases
- **[Implementation Guide](.docs/IMPLEMENTATION_GUIDE.md)** - Technical details

## Architecture

```
┌─────────────────┐
│   MCP Client    │  (e.g., Cursor IDE)
│                 │
└────────┬────────┘
         │ MCP Protocol
         │
┌────────▼─────────────────────────┐
│  MCP Server Extended             │
│  (Python - this project)         │
└────────┬─────────────────────────┘
         │ HTTP REST API
         │
┌────────▼────────┐
│ Home Assistant   │
│ REST API        │
│ /api/automation │
└─────────────────┘
```

## Requirements

- Python 3.10+
- Home Assistant instance with REST API enabled
- Long-lived access token from Home Assistant
- MCP-compatible client (e.g., Cursor IDE)

## Installation Methods

### Option 1: Home Assistant Addon (Recommended)

Easiest installation method. The addon is automatically built and published to GitHub Container Registry.

**Repository**: `https://github.com/rios0rios0/home-assistant-addons`

See [Addon Installation Guide](.docs/ADDON_INSTALLATION.md) for details.

### Option 2: Standalone Python Application

For development or when you don't want to use Home Assistant addons.

See [Setup Guide](.docs/SETUP.md) for installation instructions.

## MCP Tools

Once configured, you'll have access to these tools:

1. **list_automations** - List all automations
2. **get_automation** - Get details of a specific automation
3. **create_automation** - Create a new automation from YAML
4. **update_automation** - Update an existing automation
5. **delete_automation** - Delete an automation
6. **trigger_automation** - Manually trigger an automation
7. **enable_automation** - Enable an automation
8. **disable_automation** - Disable an automation

See [Usage Examples](.docs/USAGE_EXAMPLES.md) for detailed examples.

## Development

### Prerequisites

- Python 3.10+
- [PDM](https://pdm.fming.dev/) for dependency management

### Setup

```bash
# Install dependencies
pdm install

# Run tests
pdm run pytest

# Run the server
pdm run python -m mcp_ha_extended.server
```

See [Setup Guide](.docs/SETUP.md) and [PDM Setup](.docs/PDM_SETUP.md) for more details.

### Building the Addon

The addon is automatically built via GitHub Actions. See [Addon Build Guide](.docs/ADDON_BUILD.md) for manual build instructions and GitHub Actions setup.

## Contributing

Contributions are welcome! Please read the [Implementation Guide](.docs/IMPLEMENTATION_GUIDE.md) to understand the architecture before contributing.

## License

MIT License - see [LICENSE](../LICENSE) file for details.

## Links

- **GitHub Repository**: https://github.com/rios0rios0/home-assistant-addons
- **Container Registry**: ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended
- **Documentation**: [.docs/SUMMARY.md](.docs/SUMMARY.md)

## Support

- Check the [troubleshooting sections](.docs/SETUP.md#troubleshooting) in the documentation
- Review [Usage Examples](.docs/USAGE_EXAMPLES.md) for common patterns
- Open an issue on GitHub for bugs or feature requests
