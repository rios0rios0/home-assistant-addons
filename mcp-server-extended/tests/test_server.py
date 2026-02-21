#!/usr/bin/env python3
"""Comprehensive tests for the MCP Home Assistant server."""

import asyncio
import json
import os
from unittest.mock import AsyncMock, MagicMock, patch

import pytest
import yaml
from mcp.types import TextContent

from mcp_ha_extended.server import (
    call_tool,
    ha_api_call,
    list_tools,
    server,
)


class TestHAAPICall:
    """Test the ha_api_call function."""

    @pytest.mark.asyncio
    async def test_ha_api_call_success(self):
        """Test successful API call."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            # Create mock response
            mock_response = AsyncMock()
            mock_response.status = 200
            mock_response.content_type = "application/json"
            mock_response.json = AsyncMock(return_value={"status": "ok"})
            mock_response.raise_for_status = MagicMock()

            # Make response an async context manager
            async def response_cm():
                return mock_response
            mock_response.__aenter__ = AsyncMock(return_value=mock_response)
            mock_response.__aexit__ = AsyncMock(return_value=None)

            # Create mock session
            mock_session = AsyncMock()
            # Make request return the response as async context manager
            mock_session.request = MagicMock(return_value=mock_response)
            mock_session.__aenter__ = AsyncMock(return_value=mock_session)
            mock_session.__aexit__ = AsyncMock(return_value=None)

            with patch("mcp_ha_extended.server.aiohttp.ClientSession", return_value=mock_session):
                result = await ha_api_call("GET", "/test")

                assert result == {"status": "ok"}
                mock_session.request.assert_called_once()
                call_args = mock_session.request.call_args
                assert call_args[0][0] == "GET"
                assert call_args[0][1] == "http://homeassistant.local:8123/api/test"
                assert "Authorization" in call_args[1]["headers"]
                assert call_args[1]["headers"]["Authorization"] == "Bearer test_token"

    @pytest.mark.asyncio
    async def test_ha_api_call_with_data(self):
        """Test API call with POST data."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            mock_response = AsyncMock()
            mock_response.status = 201
            mock_response.content_type = "application/json"
            mock_response.json = AsyncMock(return_value={"id": "123"})
            mock_response.raise_for_status = MagicMock()
            mock_response.__aenter__ = AsyncMock(return_value=mock_response)
            mock_response.__aexit__ = AsyncMock(return_value=None)

            mock_session = AsyncMock()
            mock_session.request = MagicMock(return_value=mock_response)
            mock_session.__aenter__ = AsyncMock(return_value=mock_session)
            mock_session.__aexit__ = AsyncMock(return_value=None)

            with patch("mcp_ha_extended.server.aiohttp.ClientSession", return_value=mock_session):
                result = await ha_api_call("POST", "/test", {"key": "value"})

                assert result == {"id": "123"}
                call_args = mock_session.request.call_args
                assert call_args[1]["json"] == {"key": "value"}

    @pytest.mark.asyncio
    async def test_ha_api_call_no_token(self):
        """Test API call without token raises error."""
        with patch("mcp_ha_extended.server.HA_TOKEN", ""):
            with pytest.raises(ValueError, match="HA_TOKEN"):
                await ha_api_call("GET", "/test")

    @pytest.mark.asyncio
    async def test_ha_api_call_non_json_response(self):
        """Test API call with non-JSON response."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            mock_response = AsyncMock()
            mock_response.status = 204
            mock_response.content_type = "text/plain"
            mock_response.raise_for_status = MagicMock()
            mock_response.__aenter__ = AsyncMock(return_value=mock_response)
            mock_response.__aexit__ = AsyncMock(return_value=None)

            mock_session = AsyncMock()
            mock_session.request = MagicMock(return_value=mock_response)
            mock_session.__aenter__ = AsyncMock(return_value=mock_session)
            mock_session.__aexit__ = AsyncMock(return_value=None)

            with patch("mcp_ha_extended.server.aiohttp.ClientSession", return_value=mock_session):
                result = await ha_api_call("DELETE", "/test")

                assert result == {"status": "success", "status_code": 204}


class TestListTools:
    """Test the list_tools function."""

    @pytest.mark.asyncio
    async def test_list_tools(self):
        """Test that all expected tools are listed."""
        tools = await list_tools()

        assert len(tools) == 8

        tool_names = [tool.name for tool in tools]
        expected_tools = [
            "list_automations",
            "get_automation",
            "create_automation",
            "update_automation",
            "delete_automation",
            "trigger_automation",
            "enable_automation",
            "disable_automation",
        ]

        for expected_tool in expected_tools:
            assert expected_tool in tool_names

    @pytest.mark.asyncio
    async def test_list_tools_schemas(self):
        """Test that tool schemas are correct."""
        tools = await list_tools()
        tools_dict = {tool.name: tool for tool in tools}

        # Test list_automations (no required params)
        assert tools_dict["list_automations"].inputSchema["type"] == "object"
        assert "required" not in tools_dict["list_automations"].inputSchema

        # Test get_automation (requires automation_id)
        get_auto = tools_dict["get_automation"]
        assert "automation_id" in get_auto.inputSchema["properties"]
        assert "automation_id" in get_auto.inputSchema["required"]

        # Test create_automation (requires automation_yaml)
        create_auto = tools_dict["create_automation"]
        assert "automation_yaml" in create_auto.inputSchema["properties"]
        assert "automation_yaml" in create_auto.inputSchema["required"]


class TestCallTool:
    """Test the call_tool function."""

    @pytest.mark.asyncio
    async def test_list_automations(self):
        """Test listing automations."""
        mock_automations = [
            {"id": "1", "alias": "Test Auto", "enabled": True, "description": "Test"},
            {"id": "2", "alias": "Another", "enabled": False},
        ]

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=mock_automations):
                result = await call_tool("list_automations", {})

                assert len(result) == 1
                assert isinstance(result[0], TextContent)
                data = json.loads(result[0].text)
                assert data["count"] == 2
                assert len(data["automations"]) == 2
                assert data["automations"][0]["id"] == "1"
                assert data["automations"][0]["alias"] == "Test Auto"

    @pytest.mark.asyncio
    async def test_get_automation(self):
        """Test getting a specific automation."""
        mock_automation = {
            "id": "123",
            "alias": "Test",
            "trigger": [{"platform": "time", "at": "12:00"}],
            "action": [{"service": "test.service"}],
        }

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=mock_automation):
                result = await call_tool("get_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["id"] == "123"
                assert data["alias"] == "Test"

    @pytest.mark.asyncio
    async def test_create_automation(self):
        """Test creating an automation."""
        automation_yaml = """
        alias: Test Automation
        trigger:
          - platform: time
            at: "12:00:00"
        action:
          - service: system_log.write
            data:
              message: Test
        """

        mock_response = {"id": "new_123", "alias": "Test Automation"}

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=mock_response) as mock_call:
                result = await call_tool("create_automation", {"automation_yaml": automation_yaml})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "created"
                assert "result" in data

                # Verify YAML was parsed and sent
                mock_call.assert_called_once()
                call_args = mock_call.call_args
                assert call_args[0][0] == "POST"
                assert call_args[0][1] == "/automation"
                assert isinstance(call_args[0][2], dict)  # data is third positional arg
                assert call_args[0][2]["alias"] == "Test Automation"

    @pytest.mark.asyncio
    async def test_update_automation(self):
        """Test updating an automation."""
        automation_yaml = """
        alias: Updated Automation
        trigger:
          - platform: time
            at: "13:00:00"
        """

        mock_response = {"id": "123", "alias": "Updated Automation"}

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=mock_response):
                result = await call_tool(
                    "update_automation", {"automation_id": "123", "automation_yaml": automation_yaml}
                )

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "updated"

    @pytest.mark.asyncio
    async def test_delete_automation(self):
        """Test deleting an automation."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=None):
                result = await call_tool("delete_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "deleted"
                assert data["automation_id"] == "123"

    @pytest.mark.asyncio
    async def test_trigger_automation(self):
        """Test triggering an automation."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", return_value=None):
                result = await call_tool("trigger_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "triggered"
                assert data["automation_id"] == "123"

    @pytest.mark.asyncio
    async def test_enable_automation(self):
        """Test enabling an automation."""
        mock_current = {"id": "123", "alias": "Test", "enabled": False}

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call") as mock_call:
                mock_call.side_effect = [mock_current, {"id": "123", "enabled": True}]

                result = await call_tool("enable_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "enabled"

                # Verify it was called twice: GET then PUT
                assert mock_call.call_count == 2
                # Check PUT call had enabled=True (third positional arg)
                put_call = mock_call.call_args_list[1]
                assert put_call[0][0] == "PUT"  # method
                assert put_call[0][2]["enabled"] is True  # data dict

    @pytest.mark.asyncio
    async def test_disable_automation(self):
        """Test disabling an automation."""
        mock_current = {"id": "123", "alias": "Test", "enabled": True}

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call") as mock_call:
                mock_call.side_effect = [mock_current, {"id": "123", "enabled": False}]

                result = await call_tool("disable_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert data["status"] == "disabled"

                # Verify PUT call had enabled=False (third positional arg)
                put_call = mock_call.call_args_list[1]
                assert put_call[0][0] == "PUT"  # method
                assert put_call[0][2]["enabled"] is False  # data dict

    @pytest.mark.asyncio
    async def test_unknown_tool(self):
        """Test calling an unknown tool."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            result = await call_tool("unknown_tool", {})

            assert len(result) == 1
            data = json.loads(result[0].text)
            assert "error" in data
            assert "Unknown tool" in data["error"]

    @pytest.mark.asyncio
    async def test_error_handling(self):
        """Test error handling in tool calls."""
        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call", side_effect=Exception("API Error")):
                result = await call_tool("get_automation", {"automation_id": "123"})

                assert len(result) == 1
                data = json.loads(result[0].text)
                assert "error" in data
                assert "API Error" in data["error"]


class TestServerConfiguration:
    """Test server configuration and initialization."""

    def test_server_instance(self):
        """Test that server instance is created."""
        assert server is not None
        assert server.name == "home-assistant-automations"

    @pytest.mark.asyncio
    async def test_yaml_parsing(self):
        """Test YAML parsing in create_automation."""
        automation_yaml = """
        alias: YAML Test
        description: Testing YAML parsing
        trigger:
          - platform: state
            entity_id: sensor.test
        action:
          - service: notify.test
            data:
              message: Hello
        """

        parsed = yaml.safe_load(automation_yaml)
        assert parsed["alias"] == "YAML Test"
        assert len(parsed["trigger"]) == 1
        assert len(parsed["action"]) == 1

    @pytest.mark.asyncio
    async def test_invalid_yaml(self):
        """Test handling of invalid YAML."""
        invalid_yaml = "alias: [unclosed"

        with patch("mcp_ha_extended.server.HA_TOKEN", "test_token"):
            with patch("mcp_ha_extended.server.ha_api_call") as mock_call:
                result = await call_tool("create_automation", {"automation_yaml": invalid_yaml})

                # Should return error
                assert len(result) == 1
                data = json.loads(result[0].text)
                assert "error" in data
                # YAML parsing error should be caught
                mock_call.assert_not_called()
