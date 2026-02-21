# Setup Guide: Home Assistant Automation MCP Server

## Step 1: Get Home Assistant Access Token

1. Open Home Assistant web interface
2. Go to your profile (bottom left)
3. Scroll down to "Long-lived access tokens"
4. Click "Create Token"
5. Give it a name (e.g., "MCP Automation Server")
6. Copy the token (you won't see it again!)

## Step 2: Install Dependencies

```bash
cd mcp_ha_automations
pip install -r requirements.txt
```

## Step 3: Configure Environment

Create a `.env` file:

```bash
HA_URL=http://homeassistant.local:8123
HA_TOKEN=your_token_here
```

Or export environment variables:

```bash
export HA_URL="http://homeassistant.local:8123"
export HA_TOKEN="your_token_here"
```

**Note:** Replace `homeassistant.local` with your actual Home Assistant URL (could be an IP address like `192.168.1.100:8123`)

## Step 4: Test the Server

Test the server manually:

```bash
python server.py
```

The server communicates via stdio, so you'll need to test it through an MCP client.

## Step 5: Configure Cursor/IDE

### For Cursor IDE:

1. Open Cursor Settings
2. Search for "MCP" or "Model Context Protocol"
3. Add server configuration:

```json
{
  "mcpServers": {
    "home-assistant-automations": {
      "command": "python",
      "args": [
        "/absolute/path/to/mcp_ha_automations/server.py"
      ],
      "env": {
        "HA_URL": "http://homeassistant.local:8123",
        "HA_TOKEN": "your_token_here"
      }
    }
  }
}
```

### For VS Code with MCP Extension:

Similar configuration in VS Code settings.

## Step 6: Verify Installation

Once configured, you should be able to use these tools in Cursor:

- `list_automations` - List all automations
- `get_automation` - Get automation details
- `create_automation` - Create new automation
- `update_automation` - Update automation
- `delete_automation` - Delete automation
- `trigger_automation` - Trigger automation
- `enable_automation` - Enable automation
- `disable_automation` - Disable automation

## Troubleshooting

### Connection Issues

1. **Check HA_URL**: Make sure it's accessible from your machine
   ```bash
   curl http://homeassistant.local:8123/api/
   ```

2. **Check Token**: Test with curl:
   ```bash
   curl -H "Authorization: Bearer YOUR_TOKEN" \
        http://homeassistant.local:8123/api/automation
   ```

3. **Check Firewall**: Ensure port 8123 is accessible

### MCP Server Not Found

- Make sure Python path is correct in config
- Check that all dependencies are installed
- Verify the server.py file is executable: `chmod +x server.py`

### API Errors

- Check Home Assistant logs: `/config/home-assistant.log`
- Verify your token has proper permissions
- Ensure Home Assistant REST API is enabled (it is by default)

## Related Documentation

- [Quick Start Guide](QUICK_START.md) - Fast setup guide
- [Usage Examples](USAGE_EXAMPLES.md) - Code examples
- [PDM Setup](PDM_SETUP.md) - Python dependency management
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical details
- [Addon Installation](ADDON_INSTALLATION.md) - Install as Home Assistant addon
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
