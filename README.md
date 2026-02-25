<h1 align="center">Home Assistant Add-ons</h1>
<p align="center">
    <a href="https://github.com/rios0rios0/home-assistant-addons/releases/latest">
        <img src="https://img.shields.io/github/release/rios0rios0/home-assistant-addons.svg?style=for-the-badge&logo=github" alt="Latest Release"/></a>
    <a href="https://github.com/rios0rios0/home-assistant-addons/blob/main/LICENSE">
        <img src="https://img.shields.io/github/license/rios0rios0/home-assistant-addons.svg?style=for-the-badge&logo=github" alt="License"/></a>
    <a href="https://github.com/rios0rios0/home-assistant-addons/actions/workflows/build.yaml">
        <img src="https://img.shields.io/github/actions/workflow/status/rios0rios0/home-assistant-addons/build.yaml?branch=main&style=for-the-badge&logo=github" alt="Build Status"/></a>
    <a href="https://sonarcloud.io/summary/overall?id=rios0rios0_home-assistant-addons">
        <img src="https://img.shields.io/sonar/coverage/rios0rios0_home-assistant-addons?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonarqubecloud" alt="Coverage"/></a>
    <a href="https://sonarcloud.io/summary/overall?id=rios0rios0_home-assistant-addons">
        <img src="https://img.shields.io/sonar/quality_gate/rios0rios0_home-assistant-addons?server=https%3A%2F%2Fsonarcloud.io&style=for-the-badge&logo=sonarqubecloud" alt="Quality Gate"/></a>
</p>

A curated collection of 19 Home Assistant add-ons providing AI/ML services, productivity tools, monitoring, notifications, and household management. All add-ons are containerized and built for multiple architectures.

## Available Add-ons

### AI & Machine Learning

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [Ollama Container](./ollama-container/README.md) | Run large language models locally with privacy-focused inference | amd64, aarch64, armv7 | Stable |
| [Open WebUI](./open-webui/README.md) | Feature-rich chat UI for LLMs, pairs with Ollama Container | amd64, aarch64 | Experimental |
| [LocalAI](./localai/README.md) | Drop-in OpenAI API replacement for LLMs, image generation, and embeddings | amd64, aarch64 | Experimental |
| [OpenCode](./opencode/README.md) | AI coding agent in the terminal, accessible via web browser | amd64, aarch64 | Experimental |
| [OpenClaw](./openclaw/README.md) | Personal AI assistant with Telegram, Discord, Slack, and web chat | amd64, aarch64 | Experimental |
| [MCP Server Extended](./mcp-server-extended/README.md) | Model Context Protocol server for Home Assistant automation management | amd64, aarch64, armv7 | Stable |

### Voice & Speech

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [Piper TTS](./piper-tts/README.md) | Fast local neural text-to-speech with 30+ languages (Wyoming protocol) | amd64, aarch64 | Experimental |
| [Whisper.cpp](./whisper-cpp/README.md) | High-performance local speech-to-text using OpenAI Whisper models | amd64, aarch64 | Experimental |

### Automation & Workflows

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [n8n](./n8n/README.md) | Workflow automation with 400+ integrations and a visual flow editor | amd64, aarch64 | Experimental |
| [Huginn](./huginn/README.md) | Build agents that monitor websites, scrape data, and send notifications | amd64, aarch64 | Experimental |
| [Changedetection.io](./changedetection/README.md) | Monitor web pages for content changes and trigger alerts | amd64, aarch64, armv7 | Experimental |

### Monitoring & Notifications

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [Uptime Kuma](./uptime-kuma/README.md) | Monitor HTTP(s), TCP, DNS, and more with status pages and alerts | amd64, aarch64 | Experimental |
| [Ntfy](./ntfy/README.md) | Simple HTTP-based push notification server for phones and desktops | amd64, aarch64, armv7 | Experimental |

### Productivity & Utilities

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [Stirling PDF](./stirling-pdf/README.md) | PDF toolkit: merge, split, convert, compress, OCR, and more | amd64, aarch64 | Experimental |
| [IT Tools](./it-tools/README.md) | 80+ developer/sysadmin utilities (encoders, converters, generators) | amd64, aarch64, armv7 | Experimental |
| [Linkwarden](./linkwarden/README.md) | Bookmark and link manager with archiving, tagging, and collaboration | amd64, aarch64 | Experimental |

### Household Management

| Add-on | Description | Architectures | Stage |
|--------|-------------|---------------|-------|
| [Mealie](./mealie/README.md) | Recipe manager, meal planner, and shopping list with URL import | amd64, aarch64 | Experimental |
| [Grocy](./grocy/README.md) | Household ERP for groceries, chores, and batteries with barcode scanning | amd64, aarch64, armv7 | Experimental |
| [Homebox](./homebox/README.md) | Home inventory and asset management with QR labels and warranty tracking | amd64, aarch64 | Experimental |

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. The add-ons from this repository will now appear in your add-on store

## Architecture Support

All add-ons support **amd64** and **aarch64**. The following also support **armv7**:

- Changedetection.io, Grocy, IT Tools, MCP Server Extended, Ntfy, Ollama Container

## Contributing

Contributions are welcome. Please open an issue or submit a pull request on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

This project is licensed under the terms specified in the [LICENSE](LICENSE) file.
