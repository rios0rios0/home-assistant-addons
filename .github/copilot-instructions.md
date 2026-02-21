# Repository Instructions for GitHub Copilot

This document provides guidance for GitHub Copilot when working on this repository.

## Repository Overview

This is a **Home Assistant add-ons repository** that provides modular, containerized services for Home Assistant. The repository currently contains two add-ons:

1. **MCP Server Extended** - Model Context Protocol server for extended Home Assistant automation management
2. **Ollama Container** - Ollama integration for local AI capabilities

## Repository Structure

```
home-assistant-addons/
├── .github/
│   ├── copilot-instructions.md   # This file
│   └── workflows/
│       └── build.yaml            # CI/CD pipeline for Docker builds
├── mcp-server-extended/          # MCP Server Extended add-on
│   ├── config.yaml               # Add-on configuration (Home Assistant format)
│   ├── build.yaml                # Build configuration for architectures
│   ├── Dockerfile                # Container build instructions
│   ├── build.sh                  # Local build script
│   ├── pyproject.toml            # Python project configuration (PDM)
│   ├── pdm.lock                  # Locked dependencies
│   ├── config_example.json       # MCP client configuration example
│   ├── CHANGELOG.md              # Version history
│   ├── README.md                 # Add-on documentation
│   ├── rootfs/                   # S6 overlay filesystem
│   │   └── etc/services.d/
│   │       └── mcp-server-extended/
│   │           ├── run           # Service startup script
│   │           └── finish        # Service shutdown script
│   ├── src/
│   │   └── mcp_ha_extended/      # Python package
│   │       ├── __init__.py
│   │       └── server.py         # Main MCP server implementation
│   ├── tests/
│   │   ├── __init__.py
│   │   ├── test_server.py        # Server tests (pytest)
│   │   └── manual_test_api.py    # Manual testing utilities
│   └── .docs/                    # Comprehensive documentation
│       ├── SUMMARY.md            # Documentation index
│       ├── QUICK_START.md        # 5-minute setup guide
│       ├── SETUP.md              # Detailed setup
│       ├── USAGE_EXAMPLES.md     # Code examples
│       ├── IMPLEMENTATION_GUIDE.md
│       ├── PDM_SETUP.md
│       ├── ADDON_INSTALLATION.md
│       ├── ADDON_STRUCTURE.md
│       ├── ADDON_BUILD.md
│       └── ADDON_SETUP_SUMMARY.md
├── ollama-container/             # Ollama Container add-on
│   ├── config.json               # Add-on configuration
│   └── README.md                 # Add-on documentation
├── repository.json               # Repository metadata (Home Assistant format)
├── README.md                     # Main repository documentation
├── LICENSE                       # MIT License
└── .gitignore                    # Git ignore rules
```

## Home Assistant Add-on Conventions

### Configuration Files

Each add-on directory must contain a configuration file (`config.yaml` or `config.json`) following the Home Assistant add-on specification:

- **Required fields**: `name`, `version`, `slug`, `description`, `arch`, `url`, `startup`
- **Optional fields**: `boot`, `options`, `schema`, `ports`, `environment`, etc.
- **Version format**: Semantic versioning (e.g., "0.1.0")
- **Supported architectures**: `amd64`, `aarch64`, `armhf`, `armv7`
- **Startup types**: `application`, `services`, `system`, `once`

### Add-on Documentation (`README.md`)

Each add-on should have comprehensive documentation including:

- Version and architecture badges
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
- Framework: **pytest** with **pytest-asyncio**
- Run tests: `pdm run pytest` or `pdm run test`
- All async functions must use `@pytest.mark.asyncio` decorator
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
      @pytest.mark.asyncio
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
- Published to GHCR: `ghcr.io/rios0rios0/home-assistant-addons/mcp-server-extended/{arch}-addon`

## Common Tasks

### Adding a New Add-on

1. Create a new directory with a descriptive slug name (lowercase, hyphens)
2. Create `config.yaml` or `config.json` with required fields
3. Create `README.md` with comprehensive documentation
4. Add entry to main repository `README.md`
5. Ensure consistency with existing add-ons

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

- GitHub Actions builds Docker images on push to main/master
- Multi-architecture builds run in parallel
- Images tagged with branch name, version, and `latest`
- Only triggers on changes to `mcp-server-extended/` directory
- Automated testing should be added before merging breaking changes

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
