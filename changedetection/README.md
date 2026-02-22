# Changedetection.io

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Supports armv7 Architecture](https://img.shields.io/badge/armv7-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Website change detection and notification service for Home Assistant. Monitors web pages for content changes and triggers alerts, enabling you to stay informed when websites you care about are updated.

## Features

- **Website Monitoring**: Track changes on any web page by URL
- **Visual Diff**: See exactly what changed between checks with visual difference highlighting
- **Notification Support**: Get alerts via email, Slack, Discord, Telegram, and many more services through Apprise
- **CSS/XPath Selectors**: Monitor specific parts of a page using CSS selectors or XPath expressions
- **Scheduled Checks**: Configure automatic check intervals for each monitored page

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "Changedetection.io" add-on from the store

3. Configure the add-on options as needed, then click **Start**

## Configuration

```yaml
base_url: ""
notification_url: ""
fetch_workers: 2
data_dir: /share/changedetection
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `base_url` | Base URL for the Changedetection.io web interface (used for generating links in notifications) | _(empty)_ |
| `notification_url` | Default Apprise notification URL for all watches | _(empty)_ |
| `fetch_workers` | Number of parallel fetch workers for checking pages | `2` |
| `data_dir` | Directory for storing watch data and history | `/share/changedetection` |
| `log_level` | Logging verbosity (`verbose`, `debug`, `info`, `warning`, `error`, `critical`) | `info` |

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
