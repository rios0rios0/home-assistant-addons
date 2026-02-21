# Ollama Container

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Supports armv7 Architecture](https://img.shields.io/badge/armv7-yes-green.svg)

## About

Ollama integration for Home Assistant, providing easy deployment and customization of large language models. Run AI models locally on your Home Assistant instance for enhanced privacy and offline capabilities.

## Features

- **Local AI Models**: Run large language models directly on your Home Assistant hardware
- **Privacy-Focused**: All processing happens locally, no data sent to external services
- **Multiple Model Support**: Compatible with various Ollama-supported models (Llama, Mistral, Gemma, etc.)
- **Auto Model Pull**: Optionally configure a default model to be pulled on startup
- **REST API**: Exposes the Ollama API on port 11434 for integration with other add-ons and services
- **Multi-Architecture Support**: Compatible with amd64, aarch64, and armv7 platforms

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "Ollama Container" add-on from the store

3. Configure the add-on options as needed, then click **Start**

## Configuration

```yaml
ollama_host: "0.0.0.0:11434"
ollama_models: /share/ollama/models
ollama_keep_alive: 5m
ollama_num_parallel: 1
ollama_max_loaded_models: 1
default_model: ""
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `ollama_host` | Bind address for the Ollama API server | `0.0.0.0:11434` |
| `ollama_models` | Directory for storing downloaded models | `/share/ollama/models` |
| `ollama_keep_alive` | Duration to keep models loaded in memory | `5m` |
| `ollama_num_parallel` | Number of parallel requests per model | `1` |
| `ollama_max_loaded_models` | Maximum number of models loaded simultaneously | `1` |
| `default_model` | Model to auto-pull on startup (e.g., `llama3.2`, `mistral`) | _(empty)_ |
| `log_level` | Logging verbosity | `info` |

## Model Management

Once the add-on is running, you can manage models via the Ollama API:

```bash
# Pull a model
curl http://homeassistant.local:11434/api/pull -d '{"name": "llama3.2"}'

# List downloaded models
curl http://homeassistant.local:11434/api/tags

# Generate a response
curl http://homeassistant.local:11434/api/generate -d '{"model": "llama3.2", "prompt": "Hello!"}'

# Chat with a model
curl http://homeassistant.local:11434/api/chat -d '{"model": "llama3.2", "messages": [{"role": "user", "content": "Hello!"}]}'

# Delete a model
curl -X DELETE http://homeassistant.local:11434/api/delete -d '{"name": "llama3.2"}'
```

## Integration with Home Assistant

The Ollama API can be used with the [Ollama integration](https://www.home-assistant.io/integrations/ollama/) in Home Assistant for conversation agents and automations.

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
