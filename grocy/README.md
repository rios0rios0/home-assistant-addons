# Grocy

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/arch-amd64-brightgreen.svg)
![aarch64](https://img.shields.io/badge/arch-aarch64-brightgreen.svg)
![armv7](https://img.shields.io/badge/arch-armv7-brightgreen.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Grocy is an ERP-style household management solution packaged as a Home Assistant add-on. It provides a comprehensive web-based interface for tracking groceries, managing chores, monitoring batteries, and more, all with barcode scanning support.

This add-on runs the [Grocy](https://grocy.info/) application inside a Home Assistant managed container with ingress support, allowing seamless access through the Home Assistant UI.

## Features

- Grocery tracking with stock management and expiration date monitoring
- Chore management with scheduling and assignment capabilities
- Battery tracking to keep tabs on device battery replacement cycles
- Barcode scanning support for quick product identification and inventory
- Meal planning with recipe management and automatic shopping list generation

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. Find **Grocy** in the add-on store and click **Install**
8. Once installed, click **Start** to launch the add-on
9. Access the UI by clicking **Open Web UI** or through the sidebar if ingress is enabled

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `timezone` | string | `Etc/UTC` | Timezone for the Grocy application (e.g. `America/New_York`). |
| `log_level` | string | `info` | Application log level. Options: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example Configuration

```yaml
timezone: Etc/UTC
log_level: info
```

## Data Storage

The add-on uses the following mapped directories:

- **Add-on Config** (`/addon_config`): Internal add-on configuration storage.

## Support

- For issues related to this Home Assistant add-on, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).
- For issues related to Grocy itself, refer to the [upstream project](https://github.com/grocy/grocy).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
