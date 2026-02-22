# LocalAI

![Version](https://img.shields.io/badge/version-0.1.0-blue)
![amd64](https://img.shields.io/badge/arch-amd64-brightgreen)
![aarch64](https://img.shields.io/badge/arch-aarch64-brightgreen)
![Stage](https://img.shields.io/badge/stage-experimental-orange)

## About

LocalAI is a drop-in replacement for the OpenAI API that runs entirely on your local hardware. This Home Assistant add-on packages LocalAI so you can serve large language models, generate images, transcribe audio, synthesize speech, and compute embeddings without sending data to external services.

## Features

- OpenAI-compatible REST API (`/v1/chat/completions`, `/v1/embeddings`, `/v1/images/generations`, etc.)
- Run LLMs locally with GGUF model support (llama.cpp backend)
- Image generation via Stable Diffusion
- Speech-to-text transcription (Whisper)
- Text-to-speech synthesis
- Embedding generation for semantic search and RAG pipelines
- CPU-optimized inference -- no GPU required
- Persistent model storage on the Home Assistant shared volume

## Installation

1. Add this repository to Home Assistant: `https://github.com/rios0rios0/home-assistant-addons`
2. Navigate to **Settings > Add-ons > Add-on Store**.
3. Find **LocalAI** in the list and click **Install**.
4. Configure the add-on options (see below).
5. Start the add-on and check the logs.

## Configuration

| Option | Type | Default | Description |
|---|---|---|---|
| `threads` | int | `4` | Number of CPU threads used for inference. |
| `context_size` | int | `512` | Token context window size for the model. |
| `default_model` | str | `""` | Default model name returned when none is specified in the request. |
| `models_dir` | str | `/share/localai/models` | Directory where GGUF model files are stored. |
| `log_level` | str | `info` | Logging verbosity. One of `verbose`, `debug`, `info`, `warning`, `error`, `critical`. |

### Downloading Models

Place GGUF model files in the configured `models_dir` (default `/share/localai/models`). You can download models from [Hugging Face](https://huggingface.co/) or other sources and copy them into that directory via the Samba or SSH add-ons.

## API Usage

The add-on exposes an OpenAI-compatible API on port **8080**. Replace `<HA_IP>` with the IP address of your Home Assistant instance.

### Chat Completions

```bash
curl http://<HA_IP>:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "my-model",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "What is Home Assistant?"}
    ],
    "temperature": 0.7
  }'
```

### Embeddings

```bash
curl http://<HA_IP>:8080/v1/embeddings \
  -H "Content-Type: application/json" \
  -d '{
    "model": "my-embedding-model",
    "input": "Home Assistant is an open-source home automation platform."
  }'
```

## Support

- Open an issue on [GitHub](https://github.com/rios0rios0/home-assistant-addons/issues) for bugs or feature requests.
- See the [LocalAI documentation](https://localai.io/) for detailed API reference and model configuration.

## License

MIT License -- see [LICENSE](../LICENSE) for details.
