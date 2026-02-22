# Piper TTS

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/arch-amd64-brightgreen.svg)
![aarch64](https://img.shields.io/badge/arch-aarch64-brightgreen.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Piper TTS is a fast, local neural text-to-speech engine packaged as a Home Assistant add-on. It supports many languages and voices, and is fully compatible with the Wyoming protocol for seamless integration with Home Assistant voice assistants.

This add-on runs the [Piper](https://github.com/rhasspy/piper) application inside a Home Assistant managed container, exposing the Wyoming TTS protocol on port 10200 for use with Home Assistant voice pipelines.

## Features

- Local neural text-to-speech with no cloud dependency
- Support for 30+ languages with multiple voice options per language
- Wyoming protocol compatible for native Home Assistant voice integration
- Fast inference optimized for low-power hardware
- Multiple voice quality levels (low, medium, high) for balancing speed and quality

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. Find **Piper TTS** in the add-on store and click **Install**
8. Once installed, click **Start** to launch the add-on
9. Voice models are downloaded automatically on first use

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `piper_voice` | string | `en_US-lessac-medium` | Voice model to use. Format: `{language}_{REGION}-{name}-{quality}`. |
| `piper_length` | string | `1.0` | Length scale for speech speed. Lower values produce faster speech. |
| `piper_noise` | string | `0.667` | Noise scale controlling speech variation. Higher values add more variation. |
| `piper_speaker` | int | `0` | Speaker ID for multi-speaker voice models. |
| `data_dir` | string | `/share/piper-tts` | Directory for storing downloaded voice models. |
| `log_level` | string | `info` | Application log level. Options: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Example Configuration

```yaml
piper_voice: en_US-lessac-medium
piper_length: "1.0"
piper_noise: "0.667"
piper_speaker: 0
data_dir: /share/piper-tts
log_level: info
```

## Home Assistant Integration

After starting the add-on, add it as a Wyoming TTS provider in Home Assistant:

1. Navigate to **Settings** > **Devices & Services**
2. Click **Add Integration**
3. Search for **Wyoming Protocol**
4. Enter the add-on host and port `10200`
5. The Piper TTS engine will now be available as a text-to-speech provider in your voice pipelines

## Data Storage

The add-on uses the following mapped directories:

- **Add-on Config** (`/addon_config`): Internal add-on configuration storage.
- **Share** (`/share`): Shared storage accessible by other add-ons. Piper TTS stores downloaded voice models under the configured `data_dir`.

## Support

- For issues related to this Home Assistant add-on, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).
- For issues related to Piper itself, refer to the [upstream project](https://github.com/rhasspy/piper).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
