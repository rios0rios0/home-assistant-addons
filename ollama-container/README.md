# Ollama Container

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Supports armhf Architecture](https://img.shields.io/badge/armhf-yes-green.svg)
![Supports armv7 Architecture](https://img.shields.io/badge/armv7-yes-green.svg)

## About

Ollama integration for Home Assistant, providing easy deployment and customization of large language models. Run AI models locally on your Home Assistant instance for enhanced privacy and offline capabilities.

## Features

- **Local AI Models**: Run large language models directly on your Home Assistant hardware
- **Privacy-Focused**: All processing happens locally, no data sent to external services
- **Multiple Model Support**: Compatible with various Ollama-supported models
- **Easy Deployment**: Simple installation and configuration process
- **Multi-Architecture Support**: Compatible with various hardware platforms

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** → **Add-ons** → **Add-on Store**
   - Click the menu icon (⋮) in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "Ollama Container" add-on:
   - Find "Ollama Container" in the add-on store
   - Click on it and press **Install**

3. Start the add-on:
   - After installation, click **Start**
   - Optionally enable "Start on boot" for automatic startup

## Configuration

This add-on currently uses default configuration. Advanced options for model selection and resource allocation will be added in future releases.

```json
{}
```

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
