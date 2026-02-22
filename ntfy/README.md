# Ntfy - Home Assistant Add-on

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/amd64-supported-green.svg)
![aarch64](https://img.shields.io/badge/aarch64-supported-green.svg)
![armv7](https://img.shields.io/badge/armv7-supported-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

[Ntfy](https://ntfy.sh/) (pronounced "notify") is a simple HTTP-based pub-sub notification service. It allows you to send push notifications to your phone or desktop via simple PUT/POST HTTP requests, making it a perfect self-hosted solution for Home Assistant notification workflows.

This add-on packages the official ntfy server as a Home Assistant add-on, providing a fully self-hosted notification infrastructure that integrates directly into your smart home setup.

## Features

- **HTTP-based notifications** -- Send notifications using simple PUT/POST requests from automations, scripts, or any HTTP client.
- **Topic subscriptions** -- Subscribe to topics and receive notifications in real time via the web UI, mobile apps, or the CLI.
- **Mobile app support** -- Works with the official ntfy Android and iOS apps for push notifications on the go.
- **Self-hosted** -- Runs entirely on your Home Assistant instance with no external dependencies or third-party services required.
- **Ingress support** -- Access the ntfy web UI directly through the Home Assistant sidebar.
- **Persistent cache** -- Messages are cached on disk so subscribers can retrieve missed notifications.
- **Access control** -- Configure default access levels (read-write, read-only, write-only, deny) for topic security.
- **Proxy-ready** -- Designed to work behind reverse proxies with proper header forwarding.

## Installation

1. Open your Home Assistant instance.
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**.
3. Click the menu icon in the top right corner and select **Repositories**.
4. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
5. Click **Add**, then close the dialog.
6. Search for **Ntfy** in the add-on store and click on it.
7. Click **Install** and wait for the installation to complete.
8. Configure the add-on options (see below) and click **Start**.

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `base_url` | string | `""` | The public-facing base URL of the ntfy server (e.g., `https://ntfy.example.com`). Leave empty if not exposing externally. |
| `behind_proxy` | boolean | `true` | Set to `true` if ntfy is running behind a reverse proxy. Ensures correct client IP forwarding. |
| `auth_default_access` | string | `read-write` | Default access level for topics. Options: `read-write`, `read-only`, `write-only`, `deny`. |
| `cache_duration` | string | `12h` | Duration to keep messages in the cache. Supports units like `h` (hours), `m` (minutes), `d` (days). |
| `log_level` | string | `info` | Logging verbosity. Options: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example configuration

```yaml
base_url: "https://ntfy.example.com"
behind_proxy: true
auth_default_access: read-write
cache_duration: 12h
log_level: info
```

## Usage

### Sending a notification

Use a simple `curl` command or any HTTP client to publish a message to a topic:

```bash
curl -d "Garage door opened" http://homeassistant.local:8080/home-alerts
```

### Sending a notification with a title

```bash
curl -H "Title: Security Alert" \
     -d "Motion detected in the backyard" \
     http://homeassistant.local:8080/security
```

### Sending a notification with priority and tags

```bash
curl -H "Title: Temperature Warning" \
     -H "Priority: high" \
     -H "Tags: warning,thermometer" \
     -d "Living room temperature exceeded 30C" \
     http://homeassistant.local:8080/home-alerts
```

### Using in Home Assistant automations

You can integrate ntfy into your Home Assistant automations using the `rest_command` integration. Add the following to your `configuration.yaml`:

```yaml
rest_command:
  ntfy_notify:
    url: "http://localhost:8080/{{ topic }}"
    method: POST
    headers:
      Title: "{{ title }}"
      Priority: "{{ priority | default('default') }}"
      Tags: "{{ tags | default('') }}"
    payload: "{{ message }}"
```

Then use it in an automation:

```yaml
automation:
  - alias: "Notify on door open"
    trigger:
      - platform: state
        entity_id: binary_sensor.front_door
        to: "on"
    action:
      - service: rest_command.ntfy_notify
        data:
          topic: "home-alerts"
          title: "Door Alert"
          message: "The front door was opened"
          priority: "high"
          tags: "door,warning"
```

### Subscribing to notifications

- **Web UI**: Access the ntfy web interface through the Home Assistant sidebar (ingress) or at `http://homeassistant.local:8080`.
- **Mobile apps**: Install the [ntfy Android app](https://play.google.com/store/apps/details?id=io.heckel.ntfy) or [ntfy iOS app](https://apps.apple.com/us/app/ntfy/id1625396347) and subscribe to your server topics.
- **CLI**: Subscribe from the command line with `ntfy subscribe http://homeassistant.local:8080/home-alerts`.

## Data Storage

- **Cache database**: `/share/ntfy/cache/cache.db` -- Stores cached messages for offline retrieval.
- **Server configuration**: `/share/ntfy/etc/server.yml` -- Auto-generated from add-on options on each startup.

Both paths are mapped to the Home Assistant `/share` directory, ensuring data persists across add-on restarts and updates.

## Support

- For bugs or feature requests, open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).
- For ntfy-specific documentation, visit [docs.ntfy.sh](https://docs.ntfy.sh/).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
