"""Home Assistant Automation MCP Server.

This MCP server provides tools to manage Home Assistant automations via the REST API.
"""

__version__ = "0.1.0"

from mcp_ha_extended.server import main

__all__ = ["main", "__version__"]
