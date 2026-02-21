# PDM Setup Guide

This project uses [PDM](https://pdm.fming.dev/) for dependency management instead of pip and requirements.txt.

## Installation

1. Install PDM:
```bash
pip install pdm
```

Or using the official installer:
```bash
curl -sSL https://raw.githubusercontent.com/pdm-project/pdm/main/install-pdm.py | python -
```

2. Install project dependencies:
```bash
pdm install
```

3. Install with development dependencies:
```bash
pdm install --dev
```

## Usage

### Running the MCP Server

```bash
# Using PDM run
pdm run python -m mcp_ha_extended.server

# Or activate the virtual environment first
pdm shell
python -m mcp_ha_extended.server
```

### Running Tests

```bash
pdm run pytest
# or
pdm run python -m pytest tests/
```

### Adding Dependencies

```bash
# Add a production dependency
pdm add package-name

# Add a development dependency
pdm add -dG dev package-name
```

### Exporting to requirements.txt (if needed)

If you need a requirements.txt file for compatibility:
```bash
pdm export -o requirements.txt
```

## Project Structure

The project follows PEP 8 conventions with a `src/` layout:

```
mcp-server-extended/
├── src/
│   └── mcp_ha_extended/
│       ├── __init__.py
│       └── server.py
├── tests/
│   ├── __init__.py
│   └── test_api.py
├── pyproject.toml
└── README.md
```

## Configuration

All project configuration is in `pyproject.toml`, including:
- Project metadata
- Dependencies
- Development dependencies
- Build configuration
- Tool configurations (black, ruff, pytest)

## Related Documentation

- [Setup Guide](SETUP.md) - Complete development setup
- [Quick Start Guide](QUICK_START.md) - Fast setup guide
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical details
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
