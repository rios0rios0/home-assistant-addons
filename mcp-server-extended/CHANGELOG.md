# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- **BREAKING CHANGE:** `HA_URL` is now required and no longer defaults to
  `http://homeassistant.local:8123`. The old default hardcoded an unencrypted endpoint and, when
  the variable was unset, silently sent requests — bearer token included — to a guessed host.
  The add-on is unaffected because it always exports `HA_URL` from `config.yaml`; direct users of
  the MCP server must set it explicitly, as `.docs/SETUP.md` already instructs. A missing value
  now raises `ValueError: HA_URL environment variable must be set`.
- renamed the internal `_check_ha_token` guard to `_check_ha_config`, which now validates both
  `HA_URL` and `HA_TOKEN`
- added an autouse `ha_url` fixture to the test suite, since `HA_URL` no longer has a default and the
  existing tests patch only `HA_TOKEN`; added a test covering the missing-`HA_URL` error

## [0.1.0] - 2024-01-XX

### Added
- Initial release as Home Assistant addon
- Dockerfile for containerized deployment
- S6 overlay integration for process management
- Configuration schema for Home Assistant addon options
- Build scripts for multi-architecture support
- Service scripts for addon lifecycle management
- Support for aarch64, armv7, and amd64 architectures

### Features
- List all automations
- Get automation details
- Create new automations from YAML
- Update existing automations
- Delete automations
- Trigger automations manually
- Enable/disable automations
