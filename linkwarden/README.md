# Home Assistant Add-on: Linkwarden

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Linkwarden is a self-hosted, open-source bookmark and link manager. It allows you to collect, organize, and archive web pages with full-text search, tagging, and collaboration features. This add-on packages Linkwarden for Home Assistant and requires an external PostgreSQL database.

## Features

- **Bookmark Management** - Save, organize, and search your bookmarks from a clean web interface.
- **Link Archiving** - Automatically archive web pages to preserve content even if the original goes offline.
- **Tagging** - Organize links with tags and collections for easy retrieval.
- **Collaboration** - Share collections with other users and collaborate on link curation.
- **Browser Extension Support** - Use the Linkwarden browser extension to save links directly from your browser.

## Prerequisites

This add-on requires an **external PostgreSQL database**. Linkwarden does not include an embedded database. You must provide a PostgreSQL connection string in the configuration.

You can use any PostgreSQL instance accessible from your Home Assistant host, for example:

- A PostgreSQL add-on running on the same Home Assistant instance.
- A PostgreSQL server running on another machine in your network.
- A managed PostgreSQL service.

The connection string format is:

```
postgresql://USER:PASSWORD@HOST:PORT/DATABASE
```

## Installation

1. Navigate to the Home Assistant Add-on Store.
2. Add this repository URL to your add-on repositories:
   ```
   https://github.com/rios0rios0/home-assistant-addons
   ```
3. Find "Linkwarden" in the add-on list and click **Install**.
4. Configure the add-on options (see below).
5. Start the add-on.

## Configuration

| Option            | Type   | Required | Default              | Description                                         |
|-------------------|--------|----------|----------------------|-----------------------------------------------------|
| `nextauth_secret` | string | Yes      | `""`                 | Secret key for NextAuth session encryption.         |
| `database_url`    | string | Yes      | `""`                 | PostgreSQL connection string.                       |
| `data_dir`        | string | Yes      | `/share/linkwarden`  | Directory for storing archived link data.           |
| `log_level`       | list   | No       | `info`               | Log verbosity level.                                |

## Support

For bugs and feature requests, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

This add-on is provided under the same license as the parent repository. Linkwarden itself is licensed under the AGPLv3 License.
