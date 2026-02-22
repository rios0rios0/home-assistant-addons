# Mealie

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Mealie is a self-hosted recipe manager, meal planner, and shopping list application with auto-import from recipe websites. This add-on packages Mealie for seamless integration with Home Assistant, running behind the built-in ingress proxy for secure, authenticated access.

## Features

- **Recipe Import from URLs**: Automatically scrape and import recipes from thousands of supported websites
- **Meal Planning**: Plan your meals for the week with a built-in calendar and scheduling interface
- **Shopping Lists**: Generate shopping lists from recipes and manage them collaboratively
- **Nutritional Information**: View nutritional data for your recipes
- **Multi-User Support**: Multiple users with individual preferences and permissions
- **REST API**: Full API access for integration with other services and automations

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "Mealie" add-on from the store

3. Configure the add-on options as needed, then click **Start**

## Configuration

```yaml
base_url: ""
allow_signup: true
timezone: Etc/UTC
max_workers: 1
data_dir: /share/mealie
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `base_url` | External URL for Mealie (used for links in emails and notifications) | _(empty)_ |
| `allow_signup` | Allow new users to register accounts | `true` |
| `timezone` | Timezone for the application | `Etc/UTC` |
| `max_workers` | Number of worker processes for the API server | `1` |
| `data_dir` | Directory for storing Mealie data (recipes, database, images) | `/share/mealie` |
| `log_level` | Logging verbosity (`verbose`, `debug`, `info`, `warning`, `error`, `critical`) | `info` |

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
