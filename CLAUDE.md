# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Purpose

Home Assistant add-ons repository: 19 containerized services (AI/ML, voice, automation, monitoring, productivity, household). Consumers add the repository URL to their Home Assistant instance; the CI pipeline publishes multi-arch images to `ghcr.io/rios0rios0/home-assistant-addons/<addon>`.

## Architecture: Two Add-on Patterns

Every add-on lives in its own top-level directory and declares itself via `config.yaml`. Two distinct patterns exist:

### 1. Upstream-image wrapper (most add-ons)

The `Dockerfile` uses a multi-stage build: the first stage pulls the official upstream image (e.g., `binwiederhier/ntfy`, `n8nio/n8n`), the second stage starts from `${BUILD_FROM}` (an HA architecture base like `ghcr.io/home-assistant/amd64-base`) and copies the upstream binary across. It then installs **S6 overlay** (process supervisor), **Bashio** (HA add-on shell lib), and **Tempio** (template renderer) using the versions passed via `BASHIO_VERSION`, `TEMPIO_VERSION`, `S6_OVERLAY_VERSION` build args. Startup is handled by scripts in `rootfs/etc/services.d/<slug>/run` and `finish`, which read add-on options via `bashio::config`.

### 2. Python application (`mcp-server-extended`)

The only add-on that ships its own application code. It's a Python `3.10+` package managed with **PDM** (`pyproject.toml`, `pdm.lock`), built with `pdm-backend`, tested with `pytest` + `pytest-asyncio` (`asyncio_mode = "auto"` — no decorator needed). Formatting via `Black` (line length `100`), linting via `Ruff`. Source in `src/mcp_ha_extended/`, tests in `tests/`. The add-on exposes Home Assistant automation management over the **Model Context Protocol** to external MCP clients; it reads `HA_URL` and `HA_TOKEN` from env and talks to the HA REST API via `aiohttp`.

## CI/CD Architecture

Two workflows in `.github/workflows/`:

- **`build.yaml`** — orchestrator. Three jobs:
  1. `detect-changes`: discovers all `config.yaml` files (depth 2), computes the changed subset via `git diff` against the base SHA, reads each add-on's `arch[]` list, and emits a `{addon, arch}` build matrix. Pushes to `main`, tag pushes, and any change under `.github/workflows/` rebuild **all** add-ons.
  2. `build`: fans out to the reusable workflow per `{addon, arch}` pair.
  3. `manifest`: after all arch builds land, stitches per-arch digests into a multi-arch manifest and tags it as both the `config.yaml` version and `latest`. Skipped on PRs.
- **`_build-addon.yaml`** — reusable per-arch build (checkout, QEMU, parse `build.yaml`, Docker Buildx build, push digest artifact). PRs build-only, no push.

## Common Commands

### Building an add-on locally

Use the versions from the add-on's own `build.yaml`:

```bash
docker build \
  --build-arg BUILD_FROM=ghcr.io/home-assistant/amd64-base:latest \
  --build-arg BASHIO_VERSION=0.17.1 \
  --build-arg TEMPIO_VERSION=2024.11.2 \
  --build-arg S6_OVERLAY_VERSION=3.1.6.2 \
  --build-arg BUILD_ARCH=amd64 \
  -t 'local/<addon>:test' \
  ./<addon>
```

### Validating `config.yaml`

```bash
yq eval '.version' <addon>/config.yaml
yq eval '.arch[]' <addon>/config.yaml
```

### `mcp-server-extended` workflow

All commands from inside `mcp-server-extended/`:

```bash
pdm install                 # install deps (never pip install)
pdm run pytest              # run full test suite
pdm run pytest tests/test_server.py::TestFeature::test_name  # single test
pdm run test                # alias for pytest tests/
pdm run start               # run the MCP server locally
pdm run black .             # format
pdm run ruff check .        # lint
```

## Conventions Specific to This Repo

- **New add-ons default to `stage: experimental`** and `homeassistant: 2024.6.0`. Only promote to `stable` once proven.
- **Architectures**: every add-on must support `amd64` and `aarch64`; `armv7` is optional and only added when the upstream image publishes an ARMv7 variant.
- **S6 arch mapping** (in each `Dockerfile`): `armv7 → arm`, `i386 → i686`, `amd64 → x86_64`, `aarch64 → aarch64`. Preserve this mapping when copying patterns between add-ons.
- **Ingress**: most add-ons expose their UI via `ingress: true` + `ingress_port: <port>` rather than publishing host ports.
- **`rootfs/etc/services.d/<slug>/run`**: the service script runs under S6; `#!/usr/bin/with-contenv bashio` + `bashio::config 'key'` is the standard pattern for reading user-provided options.
- **Releases**: bump the add-on's `version` in `config.yaml` and any badge in its `README.md` together. The `manifest` job derives both the `<version>` and `latest` tags from `config.yaml`.
- **Changelog**: the root `CHANGELOG.md` tracks repository-level changes under `[Unreleased]`; `mcp-server-extended/` keeps its own nested `CHANGELOG.md` for the Python package.
- **n8n workflows**: `n8n/workflows/*.json` are importable workflow definitions (not code the container runs). The add-on itself is an upstream-image wrapper; the JSON files are shipped for users to import into their n8n instance.

## YAML Conventions

- Always `.yaml`, never `.yml` (CI workflows in this repo already follow this).
- Single-quote strings, double-quote only when interpolating; never quote booleans or numbers.

## What Not to Touch Without a Reason

- **Per-arch base image pins in `build.yaml`**: changing `BASHIO_VERSION` / `TEMPIO_VERSION` / `S6_OVERLAY_VERSION` rebuilds all add-ons and shifts the runtime surface — do only when intentional.
- **The `detect-changes` matrix logic**: any change under `.github/workflows/` triggers a full rebuild of all 19 add-ons.
