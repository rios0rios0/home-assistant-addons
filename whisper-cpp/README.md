# Whisper.cpp

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![amd64](https://img.shields.io/badge/arch-amd64-brightgreen.svg)
![aarch64](https://img.shields.io/badge/arch-aarch64-brightgreen.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Whisper.cpp is a high-performance local speech-to-text server packaged as a Home Assistant add-on. It uses OpenAI Whisper models compiled to C++ for fast inference, providing accurate transcription across 99 languages without any cloud dependency.

This add-on runs the [whisper.cpp](https://github.com/ggerganov/whisper.cpp) server inside a Home Assistant managed container, exposing an HTTP API on port 8080 for speech-to-text processing.

## Features

- Local speech-to-text with no cloud dependency
- Multiple model sizes to balance accuracy and resource usage
- Support for 99 languages with automatic language detection
- C++ optimized inference for fast transcription on consumer hardware
- OpenAI Whisper compatible model format

## Installation

1. Open your Home Assistant instance
2. Navigate to **Settings** > **Add-ons** > **Add-on Store**
3. Click the menu icon in the top right corner
4. Select **Repositories**
5. Add this repository URL: `https://github.com/rios0rios0/home-assistant-addons`
6. Click **Add**
7. Find **Whisper.cpp** in the add-on store and click **Install**
8. Once installed, click **Start** to launch the add-on
9. The selected model is downloaded automatically on first startup

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `model` | string | `base` | Whisper model size to use. See model sizes below. |
| `language` | string | `en` | Language code for transcription (e.g. `en`, `fr`, `de`, `es`). |
| `data_dir` | string | `/share/whisper-models` | Directory for storing downloaded model files. |
| `log_level` | string | `info` | Application log level. Options: `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Model Sizes

| Model | Disk Size | Description |
|---|---|---|
| `tiny` | ~39 MB | Fastest, lowest accuracy. Good for quick testing. |
| `base` | ~142 MB | Good balance of speed and accuracy for most use cases. |
| `small` | ~466 MB | Better accuracy with moderate resource usage. |
| `medium` | ~1.5 GB | High accuracy, requires more memory and processing time. |
| `large` | ~2.9 GB | Best accuracy, requires significant resources. |

### Example Configuration

```yaml
model: base
language: en
data_dir: /share/whisper-models
log_level: info
```

## API Usage

The add-on exposes an HTTP API on port 8080. You can send audio files for transcription:

```bash
curl http://<your-ha-host>:8080/inference \
    -H "Content-Type: multipart/form-data" \
    -F file="@audio.wav" \
    -F response_format="json"
```

## Data Storage

The add-on uses the following mapped directories:

- **Add-on Config** (`/addon_config`): Internal add-on configuration storage.
- **Share** (`/share`): Shared storage accessible by other add-ons. Whisper.cpp stores downloaded model files under the configured `data_dir`.

## Support

- For issues related to this Home Assistant add-on, please open an issue on the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).
- For issues related to whisper.cpp itself, refer to the [upstream project](https://github.com/ggerganov/whisper.cpp).

## License

This project is licensed under the terms specified in the [LICENSE](../LICENSE) file.
