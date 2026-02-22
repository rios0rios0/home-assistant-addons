# Stirling PDF

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/arch-amd64-brightgreen.svg)
![aarch64](https://img.shields.io/badge/arch-aarch64-brightgreen.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Stirling PDF is a self-hosted PDF manipulation toolkit packaged as a Home Assistant add-on. It provides a comprehensive web-based interface for performing a wide range of PDF operations without relying on external services, keeping your documents private and secure.

This add-on runs the [Stirling PDF](https://github.com/Stirling-Tools/Stirling-PDF) application inside a Home Assistant managed container with ingress support, allowing seamless access through the Home Assistant UI.

## Features

- Merge multiple PDF files into a single document
- Split PDFs by page ranges or extract individual pages
- Convert PDFs to and from various formats (images, Word, etc.)
- Compress PDFs to reduce file size
- OCR (Optical Character Recognition) for scanned documents
- Add, remove, or edit watermarks on PDF pages
- Rotate, reorder, and remove pages
- Edit PDF metadata (title, author, subject, keywords)
- Password protect and unlock PDF files
- Web-based UI accessible through Home Assistant ingress

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. Find **Stirling PDF** in the add-on store and click **Install**
8. Once installed, click **Start** to launch the add-on
9. Access the UI by clicking **Open Web UI** or through the sidebar if ingress is enabled

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `enable_login` | bool | `false` | Enable or disable the login screen for Stirling PDF. |
| `default_locale` | string | `en-US` | Default language/locale for the web interface. |
| `max_file_size` | int | `2000` | Maximum upload file size in megabytes. |
| `log_level` | string | `info` | Application log level. Options: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example Configuration

```yaml
enable_login: false
default_locale: en-US
max_file_size: 2000
log_level: info
```

## Data Storage

The add-on uses the following mapped directories:

- **Add-on Config** (`/addon_config`): Internal add-on configuration storage.
- **Share** (`/share`): Shared storage accessible by other add-ons. Stirling PDF stores its configs and logs under `/share/stirling-pdf/`.

## Support

- For issues related to this Home Assistant add-on, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).
- For issues related to Stirling PDF itself, refer to the [upstream project](https://github.com/Stirling-Tools/Stirling-PDF).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
