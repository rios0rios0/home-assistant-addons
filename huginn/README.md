# Home Assistant Add-on: Huginn

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Huginn is a system for building agents that perform automated tasks online. Think of it as a hackable version of IFTTT or Zapier running on your own server. Huginn's agents create and consume events, propagating them along a directed graph. This add-on packages Huginn for Home Assistant with an embedded MySQL database.

## Features

- **Web Scraping Agents** - Extract data from websites on a schedule and detect changes.
- **Event Monitoring** - Watch for specific events, conditions, or triggers across the web.
- **Email/Webhook Notifications** - Send alerts via email, webhooks, or other notification channels.
- **Data Transformation** - Process, filter, and transform data between agents using a pipeline.
- **Scheduled Tasks** - Run agents on custom schedules to automate recurring workflows.

## Installation

1. Navigate to the Home Assistant Add-on Store.
2. Add this repository URL to your add-on repositories:
   ```
   https://github.com/rios0rios0/home-assistant-addons
   ```
3. Find "Huginn" in the add-on list and click **Install**.
4. Configure the add-on options (see below).
5. Start the add-on.

## Configuration

| Option              | Type   | Required | Default    | Description                                      |
|---------------------|--------|----------|------------|--------------------------------------------------|
| `app_secret_token`  | string | Yes      | `""`       | Secret token for Huginn application security.    |
| `timezone`          | string | Yes      | `Etc/UTC`  | Timezone for the Huginn instance.                |
| `invitation_code`   | string | No       | `""`       | Invitation code required for new user signups.   |
| `smtp_domain`       | string | No       | `""`       | SMTP domain for outgoing email.                  |
| `smtp_server`       | string | No       | `""`       | SMTP server hostname.                            |
| `smtp_port`         | int    | Yes      | `587`      | SMTP server port.                                |
| `smtp_user_name`    | string | No       | `""`       | SMTP authentication username.                    |
| `smtp_password`     | string | No       | `""`       | SMTP authentication password.                    |
| `log_level`         | list   | No       | `info`     | Log verbosity level.                             |

## Support

For bugs and feature requests, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

This add-on is provided under the same license as the parent repository. Huginn itself is licensed under the MIT License.
