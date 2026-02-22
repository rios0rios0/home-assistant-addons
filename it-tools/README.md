# IT Tools

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/amd64-supported-green.svg)
![aarch64](https://img.shields.io/badge/aarch64-supported-green.svg)
![armv7](https://img.shields.io/badge/armv7-supported-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

Collection of 80+ developer and sysadmin utilities including encoders, converters, generators, and network tools. Served as a Home Assistant add-on with ingress support.

## About

IT Tools is a self-hosted web application that bundles a comprehensive set of handy utilities for developers and system administrators. This add-on packages the [IT Tools](https://github.com/CorentinTh/it-tools) project as a Home Assistant add-on, making it accessible directly from the Home Assistant sidebar via ingress.

## Features

- **Encoders / Decoders** - Base64, URL encoding, HTML entities, and more
- **Hash Generators** - MD5, SHA-1, SHA-256, SHA-512, and other hash functions
- **Text Utilities** - Lorem ipsum generator, text statistics, case converter
- **Converters** - Date/time, color, number base, YAML/JSON/TOML converters
- **Network Tools** - IPv4/IPv6 utilities, MAC address lookup, subnet calculator
- **UUID / ULID Generators** - Generate and validate UUIDs and ULIDs
- **JWT Debugger** - Decode, verify, and inspect JSON Web Tokens
- **Crypto Utilities** - Token generator, password strength checker, key-pair generation
- **Web Utilities** - HTTP status codes, MIME types, URL parser, OG meta generator
- **Math Tools** - Math evaluator, ETA calculator, percentage calculator
- **Measurement Converters** - Temperature, data storage, and other unit conversions
- **Development Helpers** - Crontab helper, JSON diff, SQL formatter, regex tester
- **And many more** - 80+ tools in total

## Installation

1. Add the repository to Home Assistant:
   ```
   https://github.com/rios0rios0/home-assistant-addons
   ```
2. Navigate to **Settings > Add-ons > Add-on Store**.
3. Find **IT Tools** in the list and click **Install**.
4. Start the add-on.
5. Access IT Tools from the Home Assistant sidebar (ingress).

## Configuration

This add-on requires minimal configuration. The only available option is the log level.

```yaml
log_level: info
```

### Options

| Option      | Description                              | Default |
|-------------|------------------------------------------|---------|
| `log_level` | Logging verbosity level for the add-on.  | `info`  |

**Accepted values for `log_level`:** `verbose`, `debug`, `info`, `warning`, `error`, `critical`

## Support

- Open an issue on [GitHub](https://github.com/rios0rios0/home-assistant-addons/issues) for bugs or feature requests.
- Check the [IT Tools upstream project](https://github.com/CorentinTh/it-tools) for application-specific questions.

## License

MIT License - see [LICENSE](../LICENSE) file for details.
