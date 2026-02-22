# Open WebUI

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

Feature-rich web UI for interacting with large language models. Pairs naturally with the Ollama Container add-on as a chat frontend, providing a polished browser-based interface for managing conversations, models, and prompts on your Home Assistant instance.

## Features

- **Chat Interface**: Full-featured chat UI for conversing with large language models
- **Ollama Integration**: Connect directly to a local or remote Ollama instance
- **OpenAI API Support**: Optionally use OpenAI-compatible API endpoints
- **Conversation History**: Persistent storage of chat sessions
- **Multi-Model Support**: Switch between different models within the same interface
- **Home Assistant Ingress**: Access the UI directly from the Home Assistant sidebar
- **Multi-Architecture Support**: Compatible with amd64 and aarch64 platforms

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "Open WebUI" add-on from the store

3. Configure the `ollama_base_url` option to point to your Ollama instance (e.g., the Ollama Container add-on)

4. Click **Start** to launch the add-on

## Configuration

```yaml
ollama_base_url: "http://localhost:11434"
openai_api_key: ""
data_dir: /share/open-webui
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `ollama_base_url` | URL of the Ollama API server to connect to | `http://localhost:11434` |
| `openai_api_key` | Optional OpenAI API key for OpenAI-compatible endpoints | _(empty)_ |
| `data_dir` | Directory for persistent data storage (conversations, settings) | `/share/open-webui` |
| `log_level` | Logging verbosity | `info` |

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
