# Quick Start Guide

Get automation management working in 5 minutes!

## Prerequisites

- Home Assistant running and accessible
- Long-lived access token (get from HA profile → Long-lived access tokens)

## Step 1: Install

```bash
cd mcp_ha_automations
pip install -r requirements.txt
```

## Step 2: Configure

Create `.env` file:

```bash
HA_URL=http://192.168.1.100:8123  # Your HA URL
HA_TOKEN=eyJ0eXAiOiJKV1QiLCJh...   # Your token
```

Or export:

```bash
export HA_URL="http://192.168.1.100:8123"
export HA_TOKEN="your_token_here"
```

## Step 3: Test

```bash
python test_api.py
```

Should see:
```
✓ API connection successful
✓ Automation API accessible
```

## Step 4: Configure Cursor

Add to Cursor settings (JSON):

```json
{
  "mcpServers": {
    "home-assistant-automations": {
      "command": "python",
      "args": ["/absolute/path/to/mcp_ha_automations/server.py"],
      "env": {
        "HA_URL": "http://192.168.1.100:8123",
        "HA_TOKEN": "your_token_here"
      }
    }
  }
}
```

## Step 5: Use!

In Cursor, you can now:

- "List all my automations"
- "Create an automation to turn on lights at sunset"
- "Update the morning routine automation"
- "Show me the bedroom automation details"

## Troubleshooting

**Connection failed?**
- Check `HA_URL` is correct (try in browser)
- Verify token works: `curl -H "Authorization: Bearer $HA_TOKEN" $HA_URL/api/`

**Server not found?**
- Check Python path in config
- Make sure `server.py` is executable: `chmod +x server.py`

**Tools not showing?**
- Restart Cursor
- Check Cursor logs for MCP errors

## Next Steps

- [Setup Guide](SETUP.md) - Detailed setup instructions
- [Usage Examples](USAGE_EXAMPLES.md) - Code examples and use cases
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical deep dive
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
