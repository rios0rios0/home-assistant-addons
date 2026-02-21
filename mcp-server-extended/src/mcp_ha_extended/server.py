#!/usr/bin/env python3
"""Home Assistant Automation MCP Server.

This MCP server provides tools to manage Home Assistant automations via the REST API.
"""

import asyncio
import json
import os
from typing import Any, Sequence

import aiohttp
import yaml
from mcp.server import Server
from mcp.server.stdio import stdio_server
from mcp.types import TextContent, Tool

# Configuration
HA_URL = os.getenv("HA_URL", "http://homeassistant.local:8123")
HA_TOKEN = os.getenv("HA_TOKEN", "")


def _check_ha_token():
    """Check if HA_TOKEN is set, raise error if not."""
    if not HA_TOKEN:
        raise ValueError("HA_TOKEN environment variable must be set")

# MCP Server instance
server = Server("home-assistant-automations")


async def ha_api_call(method: str, endpoint: str, data: dict | None = None) -> dict:
    """Make an authenticated API call to Home Assistant."""
    _check_ha_token()
    url = f"{HA_URL}/api{endpoint}"
    headers = {
        "Authorization": f"Bearer {HA_TOKEN}",
        "Content-Type": "application/json",
    }

    async with aiohttp.ClientSession() as session:
        async with session.request(method, url, headers=headers, json=data) as response:
            response.raise_for_status()
            if response.content_type == "application/json":
                return await response.json()
            return {"status": "success", "status_code": response.status}


@server.list_tools()
async def list_tools() -> list[Tool]:
    """List all available tools."""
    return [
        Tool(
            name="list_automations",
            description="List all automations in Home Assistant",
            inputSchema={
                "type": "object",
                "properties": {},
            },
        ),
        Tool(
            name="get_automation",
            description="Get details of a specific automation by ID",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to retrieve",
                    }
                },
                "required": ["automation_id"],
            },
        ),
        Tool(
            name="create_automation",
            description="Create a new automation from YAML configuration",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_yaml": {
                        "type": "string",
                        "description": "YAML configuration for the automation (single automation object)",
                    },
                    "alias": {
                        "type": "string",
                        "description": "Optional alias/name for the automation",
                    },
                },
                "required": ["automation_yaml"],
            },
        ),
        Tool(
            name="update_automation",
            description="Update an existing automation",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to update",
                    },
                    "automation_yaml": {
                        "type": "string",
                        "description": "Updated YAML configuration for the automation",
                    },
                },
                "required": ["automation_id", "automation_yaml"],
            },
        ),
        Tool(
            name="delete_automation",
            description="Delete an automation",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to delete",
                    },
                },
                "required": ["automation_id"],
            },
        ),
        Tool(
            name="trigger_automation",
            description="Manually trigger an automation",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to trigger",
                    },
                },
                "required": ["automation_id"],
            },
        ),
        Tool(
            name="enable_automation",
            description="Enable an automation",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to enable",
                    },
                },
                "required": ["automation_id"],
            },
        ),
        Tool(
            name="disable_automation",
            description="Disable an automation",
            inputSchema={
                "type": "object",
                "properties": {
                    "automation_id": {
                        "type": "string",
                        "description": "The automation ID to disable",
                    },
                },
                "required": ["automation_id"],
            },
        ),
    ]


@server.call_tool()
async def call_tool(name: str, arguments: dict) -> Sequence[TextContent]:
    """Handle tool calls."""

    try:
        if name == "list_automations":
            result = await ha_api_call("GET", "/automation")
            automations = result if isinstance(result, list) else result.get("automations", [])
            return [
                TextContent(
                    type="text",
                    text=json.dumps(
                        {
                            "count": len(automations),
                            "automations": [
                                {
                                    "id": auto.get("id"),
                                    "alias": auto.get("alias"),
                                    "enabled": auto.get("enabled", True),
                                    "description": auto.get("description"),
                                }
                                for auto in automations
                            ],
                        },
                        indent=2,
                    ),
                )
            ]

        elif name == "get_automation":
            automation_id = arguments["automation_id"]
            result = await ha_api_call("GET", f"/automation/{automation_id}")
            return [TextContent(type="text", text=json.dumps(result, indent=2))]

        elif name == "create_automation":
            automation_yaml = arguments["automation_yaml"]

            # Parse YAML to dict
            automation_dict = yaml.safe_load(automation_yaml)

            # Home Assistant expects the automation object directly
            result = await ha_api_call("POST", "/automation", automation_dict)
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "created", "result": result}, indent=2),
                )
            ]

        elif name == "update_automation":
            automation_id = arguments["automation_id"]
            automation_yaml = arguments["automation_yaml"]

            # Parse YAML to dict
            automation_dict = yaml.safe_load(automation_yaml)

            result = await ha_api_call("PUT", f"/automation/{automation_id}", automation_dict)
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "updated", "result": result}, indent=2),
                )
            ]

        elif name == "delete_automation":
            automation_id = arguments["automation_id"]
            await ha_api_call("DELETE", f"/automation/{automation_id}")
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "deleted", "automation_id": automation_id}, indent=2),
                )
            ]

        elif name == "trigger_automation":
            automation_id = arguments["automation_id"]
            await ha_api_call("POST", f"/automation/{automation_id}/trigger")
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "triggered", "automation_id": automation_id}, indent=2),
                )
            ]

        elif name == "enable_automation":
            automation_id = arguments["automation_id"]
            # Get current automation, update enabled flag
            current = await ha_api_call("GET", f"/automation/{automation_id}")
            current["enabled"] = True
            await ha_api_call("PUT", f"/automation/{automation_id}", current)
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "enabled", "automation_id": automation_id}, indent=2),
                )
            ]

        elif name == "disable_automation":
            automation_id = arguments["automation_id"]
            # Get current automation, update enabled flag
            current = await ha_api_call("GET", f"/automation/{automation_id}")
            current["enabled"] = False
            await ha_api_call("PUT", f"/automation/{automation_id}", current)
            return [
                TextContent(
                    type="text",
                    text=json.dumps({"status": "disabled", "automation_id": automation_id}, indent=2),
                )
            ]

        else:
            raise ValueError(f"Unknown tool: {name}")

    except Exception as e:
        return [
            TextContent(
                type="text",
                text=json.dumps({"error": str(e), "type": type(e).__name__}, indent=2),
            )
        ]


async def main():
    """Run the MCP server."""
    _check_ha_token()
    async with stdio_server() as (read_stream, write_stream):
        await server.run(
            read_stream,
            write_stream,
            server.create_initialization_options(),
        )


if __name__ == "__main__":
    asyncio.run(main())
