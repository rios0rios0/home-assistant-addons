# n8n

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/arch-amd64-green.svg)
![aarch64](https://img.shields.io/badge/arch-aarch64-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

n8n is a workflow automation platform that provides a visual flow editor for building complex multi-step automations. It offers 400+ integrations with popular services, AI capabilities, and support for webhook triggers and scheduled workflows. This add-on packages n8n as a Home Assistant add-on with ingress support for seamless access through the Home Assistant UI.

## Features

- **400+ Integrations** - Connect with a wide range of services and APIs out of the box
- **Visual Flow Editor** - Build and manage workflows using an intuitive drag-and-drop interface
- **AI Capabilities** - Leverage built-in AI nodes for intelligent automation workflows
- **Webhook Triggers** - Start workflows automatically via incoming HTTP requests
- **Scheduled Workflows** - Run automations on cron-based schedules or at fixed intervals

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. Find **n8n** in the add-on store and click **Install**

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `webhook_url` | string (optional) | `""` | External URL for webhook triggers. Leave empty if not using webhooks externally. |
| `timezone` | string | `Etc/UTC` | Timezone for scheduled workflows (e.g., `America/New_York`, `Europe/London`). |
| `data_dir` | string | `/share/n8n` | Directory where n8n stores workflow data and credentials. |
| `log_level` | string (optional) | `info` | Log verbosity level. One of: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

## Support

If you encounter any issues or have feature requests, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
