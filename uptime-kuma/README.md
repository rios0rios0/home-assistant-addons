# Uptime Kuma

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/amd64-supported-green.svg)
![aarch64](https://img.shields.io/badge/aarch64-supported-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Uptime Kuma is a self-hosted monitoring tool packaged as a Home Assistant add-on. It provides a clean and intuitive interface for monitoring the uptime of your services, websites, and network endpoints. With support for multiple protocols and a built-in notification system, it is a comprehensive solution for keeping track of your infrastructure health directly from your Home Assistant instance.

## Features

- HTTP(s), TCP, DNS, and other protocol monitoring
- Clean and responsive status pages for sharing uptime information
- Flexible notification system with support for multiple providers (Telegram, Discord, Slack, email, and many more)
- Real-time updates via WebSocket for instant status changes
- Multi-language support
- Ping chart and uptime history
- Certificate expiry monitoring
- Proxy support
- 2FA authentication

## Installation

1. Add the repository to your Home Assistant instance:
   ```
   https://github.com/rios0rios0/home-assistant-addons
   ```
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**.
3. Search for **Uptime Kuma** and click **Install**.
4. Configure the add-on options as needed (see below).
5. Start the add-on.
6. Access the Uptime Kuma web interface through the Home Assistant ingress sidebar.

## Configuration

| Option     | Type   | Default              | Description                                       |
|------------|--------|----------------------|---------------------------------------------------|
| `data_dir` | string | `/share/uptime-kuma` | Directory where Uptime Kuma stores its data files. |
| `log_level`| string | `info`               | Logging verbosity. One of `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example configuration

```yaml
data_dir: /share/uptime-kuma
log_level: info
```

## Support

- Open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues) for bugs or feature requests.
- Visit the [Uptime Kuma project](https://github.com/louislam/uptime-kuma) for upstream documentation and community resources.

## License

MIT License - see the [LICENSE](../LICENSE) file for details.
