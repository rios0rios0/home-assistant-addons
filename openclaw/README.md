# OpenClaw

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

OpenClaw is a personal AI assistant that you deploy on your own hardware. This add-on integrates OpenClaw with Home Assistant, giving you a multi-channel AI assistant accessible through WhatsApp, Telegram, Slack, Discord, and a built-in web chat interface.

## Features

- **Multi-Channel Support**: Connect through WhatsApp, Telegram, Slack, Discord, Google Chat, Signal, and more
- **Web Chat Interface**: Built-in webchat accessible via Home Assistant ingress
- **Multiple AI Providers**: Supports Anthropic Claude, OpenAI GPT, and other providers
- **Privacy-First**: Self-hosted on your own hardware, your data stays local
- **Home Assistant Integration**: Access HA API for automation-aware responses
- **WebSocket Gateway**: Real-time bidirectional communication

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "OpenClaw" add-on from the store

3. Configure at least one AI provider API key and enable desired channels

4. Click **Start** and access the web chat from the add-on's **Web UI** tab

## Configuration

```yaml
anthropic_api_key: ""
openai_api_key: ""
default_model: ""
telegram_bot_token: ""
telegram_enabled: false
discord_bot_token: ""
discord_enabled: false
slack_bot_token: ""
slack_enabled: false
webchat_enabled: true
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `anthropic_api_key` | Anthropic API key for Claude models | _(empty)_ |
| `openai_api_key` | OpenAI API key for GPT models | _(empty)_ |
| `default_model` | AI model to use (e.g., `anthropic/claude-sonnet-4-20250514`) | _(empty)_ |
| `telegram_bot_token` | Telegram Bot API token | _(empty)_ |
| `telegram_enabled` | Enable Telegram channel | `false` |
| `discord_bot_token` | Discord bot token | _(empty)_ |
| `discord_enabled` | Enable Discord channel | `false` |
| `slack_bot_token` | Slack bot token | _(empty)_ |
| `slack_enabled` | Enable Slack channel | `false` |
| `webchat_enabled` | Enable built-in web chat interface | `true` |
| `log_level` | Logging verbosity | `info` |

## Channel Setup

### Telegram
1. Create a bot via [@BotFather](https://t.me/BotFather) on Telegram
2. Copy the bot token and paste it in the `telegram_bot_token` field
3. Set `telegram_enabled` to `true`

### Discord
1. Create an application at the [Discord Developer Portal](https://discord.com/developers/applications)
2. Create a bot under your application and copy the token
3. Paste it in the `discord_bot_token` field
4. Set `discord_enabled` to `true`

### Slack
1. Create a Slack App at [api.slack.com](https://api.slack.com/apps)
2. Add the Bot Token Scopes your assistant needs
3. Install the app to your workspace and copy the Bot User OAuth Token
4. Paste it in the `slack_bot_token` field
5. Set `slack_enabled` to `true`

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
