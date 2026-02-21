# MCP Server Extended - Home Assistant Addon

This is a Home Assistant addon that provides a Model Context Protocol (MCP) server for extended Home Assistant automation management.

## Features

- List all automations
- Get automation details
- Create new automations from YAML
- Update existing automations
- Delete automations
- Trigger automations manually
- Enable/disable automations

## Installation

### Option 1: Manual Installation (Development)

1. Copy this directory to your Home Assistant `addons` directory:
   ```bash
   cp -r mcp-server-extended /config/addons/
   ```

2. In Home Assistant, go to **Settings** → **Add-ons** → **Add-on Store** → **⋮** (three dots) → **Repositories**

3. Add repository: `/config/addons`

4. The addon should appear in the addon store

### Option 2: GitHub Repository

1. In Home Assistant, go to **Settings** → **Add-ons** → **Add-on Store** → **⋮** (three dots) → **Repositories**

2. Add repository: `https://github.com/rios0rios0/home-assistant-addons`

3. Install the addon from the store

## Configuration

The addon requires the following configuration options:

- **ha_url** (required): The URL of your Home Assistant instance (e.g., `http://homeassistant.local:8123`)
- **ha_token** (required): A long-lived access token from Home Assistant
- **log_level** (optional): Logging level (`verbose`, `debug`, `info`, `warning`, `error`, `critical`). Default: `info`

### Getting a Long-Lived Access Token

1. In Home Assistant, go to your profile (click on your name in the sidebar)
2. Scroll down to **Long-lived access tokens**
3. Click **Create Token**
4. Give it a name (e.g., "MCP Addon")
5. Copy the token and paste it into the addon configuration

## Usage

Once installed and configured, the MCP server will run automatically. It communicates via stdio using the Model Context Protocol.

To use this with an MCP client (like Cursor IDE), configure it to connect to the addon's stdio interface.

## Building

For building instructions, see the [Addon Build Guide](ADDON_BUILD.md) which covers:
- Manual building with `build.sh`
- Building with Docker directly
- Automated builds with GitHub Actions

## Architecture

The addon uses:
- **S6 Overlay**: Process supervision and init system
- **bashio**: Home Assistant addon helper library
- **tempio**: Template engine for configuration files
- **Python 3.10+**: Runtime environment
- **PDM**: Python dependency management

## Troubleshooting

### Addon won't start

1. Check the logs: **Settings** → **Add-ons** → **MCP Server Extended** → **Logs**
2. Verify `ha_url` and `ha_token` are correctly configured
3. Ensure Home Assistant is accessible from the addon container

### Connection issues

- Verify the `ha_url` is correct and accessible
- Check that the long-lived access token is valid
- Ensure Home Assistant's REST API is enabled

## Development

For development setup and testing instructions, see:
- [Setup Guide](SETUP.md) - Detailed development setup
- [PDM Setup](PDM_SETUP.md) - Python dependency management
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical details

## Related Documentation

- [Addon Structure](ADDON_STRUCTURE.md) - Understanding the addon architecture
- [Addon Build Guide](ADDON_BUILD.md) - Building and deploying
- [Usage Examples](USAGE_EXAMPLES.md) - How to use the MCP tools
- [Documentation Index](SUMMARY.md) - Complete documentation navigation

## License

MIT License - see [LICENSE](../../LICENSE) file for details.
