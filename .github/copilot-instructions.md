# Repository Instructions for GitHub Copilot

This document provides guidance for GitHub Copilot when working on this repository.

## Repository Overview

This is a **Home Assistant add-ons repository** that provides modular, containerized services for Home Assistant. The repository contains 19 add-ons across six categories:

- **AI & Machine Learning**: `mcp-server-extended`, `ollama-container`, `open-webui`, `localai`, `opencode`, `openclaw`
- **Voice & Speech**: `piper-tts`, `whisper-cpp`
- **Automation & Workflows**: `n8n`, `huginn`, `changedetection`
- **Monitoring & Notifications**: `uptime-kuma`, `ntfy`
- **Productivity & Utilities**: `stirling-pdf`, `it-tools`, `linkwarden`
- **Household Management**: `mealie`, `grocy`, `homebox`

Two add-ons are marked **stable** (`mcp-server-extended`, `ollama-container`); all others are **experimental**.

## Repository Structure

```
home-assistant-addons/
├── .github/
│   ├── copilot-instructions.md     # This file
│   └── workflows/
│       ├── build.yaml              # CI/CD pipeline (change detection + manifest creation)
│       └── _build-addon.yaml       # Reusable per-arch build workflow
├── <addon-name>/                   # One directory per add-on (19 total)
│   ├── config.yaml                 # Add-on metadata (name, version, arch, options…)
│   ├── build.yaml                  # Base images per architecture + build args
│   ├── Dockerfile                  # Multi-arch container with S6 overlay, Bashio, Tempio
│   ├── rootfs/                     # Files overlaid onto the container filesystem
│   │   └── etc/services.d/<slug>/
│   │       ├── run                 # S6 service startup script
│   │       └── finish              # S6 service finish script
│   └── README.md                   # Add-on documentation
├── mcp-server-extended/            # Python add-on (additional files)
│   ├── pyproject.toml              # Python project config (PDM)
│   ├── pdm.lock                    # Locked dependencies
│   ├── src/mcp_ha_extended/        # Python package source
│   ├── tests/                      # pytest test suite
│   └── .docs/                      # Extended documentation
├── ADDON_IDEAS.md                  # Backlog of potential future add-ons
├── CONTRIBUTING.md                 # Contribution guidelines
├── repository.json                 # Repository metadata (Home Assistant format)
├── README.md                       # Main repository documentation
├── LICENSE                         # MIT License
└── .gitignore                      # Git ignore rules
```

## Home Assistant Add-on Conventions

### Configuration Files

Each add-on directory must contain a `config.yaml` file following the Home Assistant add-on specification:

- **Required fields**: `name`, `version`, `slug`, `description`, `arch`, `url`, `startup`
- **Common optional fields**: `hassio_api`, `homeassistant`, `host_ipc`, `host_network`, `host_dbus`, `image`, `ingress`, `ingress_port`, `ingress_stream`, `init`, `map`, `options`, `schema`, `ports`, `ports_description`, `stage`
- **Version format**: Semantic versioning (e.g., `1.0.0`)
- **Supported architectures**: `amd64`, `aarch64`, `armv7` (all add-ons support amd64 and aarch64; six also support armv7)
- **Startup types**: `application`, `services`, `system`, `once`
- **Stage values**: `stable`, `experimental` (default to `experimental` for new add-ons)
- **Minimum HA version**: All add-ons require `homeassistant: 2024.6.0`

Example `config.yaml`:
```yaml
version: "1.0.0"
slug: my-addon
name: My Add-on
description: Short description of the add-on
url: https://github.com/rios0rios0/home-assistant-addons/tree/main/my-addon
arch:
  - amd64
  - aarch64
homeassistant: 2024.6.0
hassio_api: false
host_network: false
host_ipc: false
host_dbus: false
image: ghcr.io/rios0rios0/home-assistant-addons/my-addon
ingress: true
ingress_port: 8080
init: false
startup: application
stage: experimental
options:
  my_option: default_value
schema:
  my_option: str
```

### Build Configuration (`build.yaml`)

Each add-on must have a `build.yaml` file specifying architecture-specific base images and standard build arguments:

```yaml
build_from:
  aarch64: ghcr.io/home-assistant/aarch64-base:latest
  amd64: ghcr.io/home-assistant/amd64-base:latest
  armv7: ghcr.io/home-assistant/armv7-base:latest
args:
  BASHIO_VERSION: 0.17.1
  TEMPIO_VERSION: 2024.11.2
  S6_OVERLAY_VERSION: 3.1.6.2
```

### Add-on Documentation (`README.md`)

Each add-on should have comprehensive documentation including:

- Version and architecture badges (shields.io, `style=for-the-badge`)
- About section describing the add-on
- Features list
- Installation instructions
- Configuration examples
- Support information
- License reference

### Repository Metadata (`repository.json`)

The root `repository.json` file contains:

- Repository name
- Repository URL
- Maintainer information

## MCP Server Extended — Key Technologies & Dependencies

- **Python 3.10+**: Minimum required version
- **MCP SDK** (`mcp>=0.9.0`): Model Context Protocol implementation
- **aiohttp** (`>=3.9.0`): Async HTTP client for Home Assistant API
- **PyYAML** (`>=6.0`): YAML parsing for automation configurations
- **python-dotenv** (`>=1.0.0`): Environment variable management
- **PDM**: Package and dependency manager (NOT pip/poetry)

## Development Tools

### Package Management (MCP Server Extended)
- **ALWAYS use PDM** for dependency management
- Add dependencies: `pdm add <package>`
- Install dependencies: `pdm install`
- Update dependencies: `pdm update`
- **NEVER** use `pip install -r requirements.txt` (no requirements.txt exists)

### Testing
- Framework: **pytest** with **pytest-asyncio** (`asyncio_mode = "auto"` — no `@pytest.mark.asyncio` decorator needed)
- Run tests: `pdm run pytest` or `pdm run test`
- Mock external API calls (Home Assistant) using `unittest.mock`

### Code Quality
- **Black**: Code formatter (line length: 100)
- **Ruff**: Linter (targets Python 3.10+)
- Format code: `pdm run black .`
- Lint code: `pdm run ruff check .`

## Coding Standards and Best Practices

### General Guidelines

1. **Minimal changes**: Make the smallest possible changes to achieve the goal
2. **Consistency**: Follow existing patterns and conventions in the repository
3. **Documentation**: Update documentation when making changes to add-ons or features
4. **JSON formatting**: Use 2-space indentation for all JSON files
5. **Markdown formatting**: Follow existing markdown style with proper headers and lists

### Python Style (MCP Server Extended)
- Line length: **100 characters** (configured in pyproject.toml)
- Target: Python 3.10+ (use modern type hints: `str | None` instead of `Optional[str]`, `dict[str, Any]` for dictionaries)
- Use type hints for all function signatures
- Use descriptive variable names
- Follow PEP 8 conventions (enforced by Black and Ruff)

### Async/Await Patterns
- Use `async`/`await` for all I/O operations
- Use `aiohttp.ClientSession()` for HTTP requests (not `requests`)
- Always use async context managers: `async with session.request(...) as response:`
- Handle exceptions appropriately in async functions

### MCP Tool Definitions
Each tool follows this pattern:
```python
Tool(
    name="tool_name",
    description="Clear description of what the tool does",
    inputSchema={
        "type": "object",
        "properties": {
            "param_name": {
                "type": "string",  # or "object", "array", etc.
                "description": "Parameter description"
            }
        },
        "required": ["param_name"]  # List required parameters
    }
)
```

### Home Assistant API Patterns
- Base URL: `{HA_URL}/api`
- All requests require: `Authorization: Bearer {HA_TOKEN}`
- Common endpoints:
  - `GET /api/automation` - List automations
  - `GET /api/automation/{id}` - Get automation
  - `POST /api/automation` - Create automation
  - `PUT /api/automation/{id}` - Update automation
  - `DELETE /api/automation/{id}` - Delete automation
  - `POST /api/automation/{id}/trigger` - Trigger automation

## Environment Variables

- `HA_URL`: Home Assistant URL (e.g., `http://homeassistant.local:8123`)
- `HA_TOKEN`: Long-lived access token from Home Assistant
- Both are **required** for the MCP server to function

## Version Management

- Use semantic versioning (MAJOR.MINOR.PATCH)
- Increment PATCH for bug fixes
- Increment MINOR for new features (backward compatible)
- Increment MAJOR for breaking changes
- Update version in both `config.yaml` and `README.md` badges when releasing

## Testing Guidelines

### Unit Tests
- Mock external dependencies (Home Assistant API calls)
- Use `AsyncMock` for async functions
- Test both success and error paths
- This project uses test classes (not standalone functions)
- Example pattern:
  ```python
  class TestFeature:
      async def test_function_name(self):
          with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
              mock_response = AsyncMock()
              # ... setup mock
              result = await function_under_test()
              assert result == expected_value
  ```

### Manual Testing
- Use `tests/manual_test_api.py` for live testing
- Requires actual Home Assistant instance
- Test against real API endpoints

## Docker & Deployment

### Multi-Architecture Support
- Builds for: `aarch64`, `armv7`, `amd64`
- Uses QEMU for cross-compilation
- Base images defined in `build.yaml`

### Home Assistant Addon
- Configuration: `config.yaml`
- Startup scripts: `rootfs/etc/services.d/mcp-server-extended/`
- Published to GHCR: `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended`

## Common Tasks

### Adding a New Add-on

1. Create a new directory with a descriptive slug name (lowercase, hyphens)
2. Add the standard files: `Dockerfile`, `config.yaml`, `build.yaml`, `rootfs/`, `README.md`
3. Use the standard build args in `build.yaml` (see `build.yaml` for current `BASHIO_VERSION`, `TEMPIO_VERSION`, and `S6_OVERLAY_VERSION` values)
4. Build locally to verify (replace version values with those from the add-on's `build.yaml`):
   ```bash
   docker build \
     --build-arg BUILD_FROM=ghcr.io/home-assistant/amd64-base:latest \
     --build-arg BASHIO_VERSION=<BASHIO_VERSION> \
     --build-arg TEMPIO_VERSION=<TEMPIO_VERSION> \
     --build-arg S6_OVERLAY_VERSION=<S6_OVERLAY_VERSION> \
     --build-arg BUILD_ARCH=amd64 \
     -t local/<addon-name>:test \
     ./<addon-name>
   ```
5. Validate the configuration:
   ```bash
   yq eval '.version' <addon-name>/config.yaml
   yq eval '.arch[]' <addon-name>/config.yaml
   ```
6. Add an entry to the main repository `README.md`
7. Ensure consistency with existing add-ons

### Adding a New MCP Tool

1. Add `Tool` definition in `list_tools()` function
2. Add handler in `call_tool()` function
3. Implement the API logic (use `ha_api_call()` helper)
4. Add unit tests in `tests/test_server.py`
5. Update documentation in `.docs/USAGE_EXAMPLES.md`

### Updating Dependencies

1. `pdm add <package>@<version>` or `pdm add <package>`
2. Verify with `pdm install`
3. Run tests: `pdm run pytest`
4. Update `pyproject.toml` if needed (usually automatic)

### Debugging

1. Check environment variables are set
2. Test Home Assistant API manually: `curl -H "Authorization: Bearer $HA_TOKEN" $HA_URL/api/`
3. Run server with debug logging
4. Check MCP client logs for protocol errors

## Security Considerations

- Never hardcode `HA_TOKEN` or `HA_URL`
- Always validate user input before sending to Home Assistant
- Use environment variables for sensitive configuration
- Token should have minimal necessary permissions in Home Assistant
- Sanitize YAML input to prevent injection attacks
- Do not commit files containing tokens or secrets (.env, credentials.json, etc.)

## CI/CD

GitHub Actions builds Docker images on every push to `main`, on tag pushes (`v*`), on pull requests targeting `main`, and on manual dispatch.

### Workflow Architecture

- **`build.yaml`** — Orchestration workflow with three jobs:
  1. **`detect-changes`**: Discovers which add-ons changed (via `git diff`), reads their supported architectures from `config.yaml`, and produces a build matrix. On pushes to `main` or tags, all add-ons are rebuilt. If any file under `.github/workflows/` changes, all add-ons are rebuilt.
  2. **`build`**: Fans out to the reusable `_build-addon.yaml` workflow for each `{addon, arch}` pair in the matrix.
  3. **`manifest`**: Creates multi-arch manifests after all `build` jobs succeed (skipped for PRs). Tags images with the version from `config.yaml` and `latest`.

- **`_build-addon.yaml`** — Reusable per-arch build workflow. Steps: checkout → QEMU setup → parse `build.yaml` for base images/args → Docker Buildx build → push digest artifact. PRs validate only (no push).
- **`claude.yaml`** — Claude Code agent triggered by issue comments, PR review comments, opened/assigned issues, and submitted PR reviews. Delegates to a reusable workflow in `rios0rios0/.github`.
- **`claude-code-review.yaml`** — Automated PR review via Claude Code on PR open/sync/reopen. Delegates to a reusable workflow in `rios0rios0/.github`.
- **`release.yaml`** — Triggered on pushes to `main`. Delegates to a reusable workflow in `rios0rios0/pipelines` to create Git tags from merged bump PRs.

### Registry

Images are published to GHCR: `ghcr.io/rios0rios0/home-assistant-addons/<addon-name>`

### Image Tags

- `<version>` (from `config.yaml`)
- `latest`

## Best Practices

1. **Minimal changes**: Make surgical, focused changes
2. **Test-driven**: Write tests before or alongside code changes
3. **Type safety**: Use type hints consistently
4. **Error handling**: Provide clear error messages
5. **Async-first**: Use async/await for all I/O operations
6. **Documentation**: Update docs when adding features
7. **Security**: Never commit tokens or secrets
8. **PDM-only**: Do not introduce pip/poetry workflows for MCP Server Extended

## Preferred Solutions

- **HTTP Client**: aiohttp (not requests)
- **Package Manager**: PDM (not pip or poetry)
- **Testing**: pytest + pytest-asyncio (not unittest)
- **Formatting**: Black (configured in pyproject.toml)
- **Linting**: Ruff (not flake8 or pylint)
- **Type Checking**: Built-in type hints (consider adding mypy if needed)

## Support and Contribution

- **Issues**: Report bugs and feature requests via GitHub Issues
- **Pull Requests**: Welcome for bug fixes, features, and documentation improvements
- **Maintainer**: @rios0rios0
- **License**: See LICENSE file in repository root
