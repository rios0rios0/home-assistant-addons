# OpenCode

![Version](https://img.shields.io/badge/version-0.1.0-blue.svg)
![Supports amd64 Architecture](https://img.shields.io/badge/amd64-yes-green.svg)
![Supports aarch64 Architecture](https://img.shields.io/badge/aarch64-yes-green.svg)
![Stage](https://img.shields.io/badge/stage-experimental-orange.svg)

## About

OpenCode is an AI coding agent built for the terminal. This add-on packages OpenCode with a web-based terminal (ttyd), making it accessible directly from your Home Assistant dashboard via ingress.

Write, debug, and refactor code using AI assistance — all from within your Home Assistant interface.

## Features

- **AI-Powered Coding**: Full terminal-based AI coding agent with file read/write/execute capabilities
- **Web Terminal Access**: Browser-based terminal via Home Assistant ingress — no SSH needed
- **Multiple AI Providers**: Supports Anthropic, OpenAI, Google Gemini, and Groq
- **Multi-Agent System**: Built-in "Build" (full access) and "Plan" (read-only analysis) agents
- **Shared Workspace**: Works on files in the `/share` directory, accessible by other add-ons

## Installation

1. Add this repository to your Home Assistant instance:
   - Navigate to **Settings** > **Add-ons** > **Add-on Store**
   - Click the menu icon in the top right
   - Select **Repositories**
   - Add the URL: `https://github.com/rios0rios0/home-assistant-addons`

2. Install the "OpenCode" add-on from the store

3. Configure at least one AI provider API key, then click **Start**

4. Access OpenCode from the add-on's **Web UI** tab in the sidebar

## Configuration

```yaml
anthropic_api_key: ""
openai_api_key: ""
gemini_api_key: ""
groq_api_key: ""
default_provider: anthropic
default_model: ""
working_directory: /share
log_level: info
```

### Options

| Option | Description | Default |
|--------|-------------|---------|
| `anthropic_api_key` | Anthropic API key for Claude models | _(empty)_ |
| `openai_api_key` | OpenAI API key for GPT models | _(empty)_ |
| `gemini_api_key` | Google Gemini API key | _(empty)_ |
| `groq_api_key` | Groq API key for fast inference | _(empty)_ |
| `default_provider` | Default AI provider to use | `anthropic` |
| `default_model` | Specific model to use (provider-dependent) | _(empty)_ |
| `working_directory` | Directory for OpenCode to operate in | `/share` |
| `log_level` | Logging verbosity | `info` |

## Provider Setup

You need at least one API key configured:

- **Anthropic**: Get your key at [console.anthropic.com](https://console.anthropic.com/)
- **OpenAI**: Get your key at [platform.openai.com](https://platform.openai.com/)
- **Google Gemini**: Get your key at [aistudio.google.com](https://aistudio.google.com/)
- **Groq**: Get your key at [console.groq.com](https://console.groq.com/)

## Support

For issues and feature requests, please use the [GitHub repository](https://github.com/rios0rios0/home-assistant-addons/issues).

## License

See the [LICENSE](../LICENSE) file in the repository root.
