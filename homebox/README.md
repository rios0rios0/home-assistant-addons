# Homebox

![Version](https://img.shields.io/badge/version-0.1.0-blue)
![amd64](https://img.shields.io/badge/arch-amd64-green)
![aarch64](https://img.shields.io/badge/arch-aarch64-green)
![Stage](https://img.shields.io/badge/stage-experimental-orange)

## About

Homebox is a home inventory and asset management system packaged as a Home Assistant add-on. It allows you to track your belongings, organize them by location, monitor warranties, generate QR code labels, and schedule maintenance tasks -- all from within your Home Assistant instance.

This add-on runs the [Homebox](https://github.com/sysadminsmedia/homebox) application using Home Assistant ingress, providing seamless access through the Home Assistant sidebar without exposing additional ports.

## Features

- **Inventory tracking** -- Catalog and manage all your household items with detailed metadata.
- **Location management** -- Organize items by rooms, buildings, or custom locations.
- **Warranty tracking** -- Keep track of warranty expiration dates and documentation.
- **QR code labels** -- Generate printable QR code labels for physical items.
- **Maintenance schedules** -- Set up recurring maintenance reminders for your assets.
- **Low memory footprint** -- Built with Go and SQLite, the application runs efficiently on minimal hardware.

## Installation

1. Add the repository to your Home Assistant instance:
   ```
   https://github.com/rios0rios0/home-assistant-addons
   ```
2. Navigate to **Settings > Add-ons > Add-on Store**.
3. Search for **Homebox** and click **Install**.
4. Configure the add-on options as needed (see below).
5. Start the add-on and access it via the Home Assistant sidebar.

## Configuration

| Option     | Type   | Default          | Description                                      |
|------------|--------|------------------|--------------------------------------------------|
| `data_dir` | string | `/share/homebox` | Directory where Homebox stores its data and database. |
| `timezone` | string | `Etc/UTC`        | Timezone for the application (e.g., `America/New_York`). |
| `log_level`| string | `info`           | Log verbosity level. One of: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example configuration

```yaml
data_dir: /share/homebox
timezone: America/New_York
log_level: info
```

## Support

- Open an issue on [GitHub](https://github.com/rios0rios0/home-assistant-addons/issues) for bugs or feature requests.
- Refer to the [Homebox documentation](https://homebox.software/) for application-specific questions.

## License

MIT License -- see [LICENSE](../LICENSE) file for details.
