# MCP Server Extended

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Supports armhf Architecture](https://img.shields.io/badge/armhf-yes-green.svg)
![Supports armv7 Architecture](https://img.shields.io/badge/armv7-yes-green.svg)

## About

Extended MCP (Model Context Protocol) server functionality for Home Assistant. This add-on provides modular and containerized services for advanced integrations, enabling enhanced communication between AI models and your Home Assistant instance.

## Features

- **Model Context Protocol Support**: Full implementation of the MCP specification
- **Modular Architecture**: Easy to extend and customize
- **Containerized Deployment**: Isolated and secure execution environment
- **Multi-Architecture Support**: Compatible with various hardware platforms

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** → **Add-ons** → **Add-on Store**
   - Click the menu icon (⋮) in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "MCP Server Extended" add-on:
   - Find "MCP Server Extended" in the add-on store
   - Click on it and press **Install**

3. Start the add-on:
   - After installation, click **Start**
   - Optionally enable "Start on boot" for automatic startup

## Configuration

This add-on currently uses default configuration. Advanced options will be added in future releases.

```json
{}
```

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
