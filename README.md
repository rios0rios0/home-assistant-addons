# Home Assistant Add-ons by rios0rios0

![Project Maintenance](https://img.shields.io/badge/maintainer-rios0rios0-blue.svg)

Home Assistant add-ons repository featuring mcp-server-extended and ollama-container. Built on the hassio-ihost-addon structure, it provides modular, containerized services for extended MCP server functionality and Ollama integration, enabling easy deployment and customization.

## Available Add-ons

### MCP Server Extended

Model Context Protocol (MCP) server for extended Home Assistant automation management. Provides 8 MCP tools for listing, creating, updating, deleting, triggering, and enabling/disabling automations via the Home Assistant REST API.

- **Version**: 0.1.0
- **Architectures**: amd64, aarch64, armv7
- **Startup**: Services
- [Documentation](./mcp-server-extended/README.md)

### Ollama Container

Ollama integration for Home Assistant, enabling easy deployment and customization of large language models for local AI capabilities.

- **Version**: 1.0.0
- **Architectures**: amd64, aarch64, armhf, armv7
- **Startup**: Application
- [Documentation](./ollama-container/README.md)

## Installation

To use these add-ons, follow these steps:

1. Open your Home Assistant instance
2. Navigate to **Settings** → **Add-ons** → **Add-on Store**
3. Click the menu icon (⋮) in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. The add-ons from this repository will now appear in your add-on store

## Support

If you encounter any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.
