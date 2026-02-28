# Contributing

Contributions are welcome. By participating, you agree to maintain a respectful and constructive environment.

For coding standards, testing patterns, architecture guidelines, commit conventions, and all
development practices, refer to the **[Development Guide](https://github.com/rios0rios0/guide/wiki)**.

## Prerequisites

- [Docker](https://docs.docker.com/get-docker/) 20.10+
- [Home Assistant](https://www.home-assistant.io/installation/) (for end-to-end testing)
- [yq](https://github.com/mikefarah/yq) (for parsing `config.yaml` and `build.yaml`)

## Add-on Structure

Each add-on lives in its own directory and follows the Home Assistant add-on format:

```
<addon-name>/
├── Dockerfile       # Multi-arch Docker image with S6 overlay, Bashio, and Tempio
├── config.yaml      # Add-on metadata (name, version, architectures, options)
├── build.yaml       # Base images per architecture and build arguments
├── rootfs/          # Files overlaid onto the container filesystem (S6 service scripts)
└── README.md        # Add-on-specific documentation
```

## Development Workflow

1. Fork and clone the repository
2. Create a branch: `git checkout -b feat/my-change`
3. Build a specific add-on image locally (e.g. `ntfy` for `amd64`):
   ```bash
   docker build \
     --build-arg BUILD_FROM=ghcr.io/home-assistant/amd64-base:latest \
     --build-arg BASHIO_VERSION=0.17.1 \
     --build-arg TEMPIO_VERSION=2024.11.2 \
     --build-arg S6_OVERLAY_VERSION=3.1.6.2 \
     --build-arg BUILD_ARCH=amd64 \
     -t local/ntfy:test \
     ./ntfy
   ```
4. Run the built image to verify it starts correctly:
   ```bash
   docker run --rm -it local/ntfy:test
   ```
5. Validate the add-on configuration:
   ```bash
   yq eval '.version' ntfy/config.yaml
   yq eval '.arch[]' ntfy/config.yaml
   ```
6. Test the add-on in Home Assistant by adding this repository as a local add-on repository
7. Commit following the [commit conventions](https://github.com/rios0rios0/guide/wiki/Life-Cycle/Git-Flow)
8. Open a pull request against `main`
