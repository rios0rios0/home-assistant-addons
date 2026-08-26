# Quick Start Guide

Get automation management working in 5 minutes!

## Prerequisites

- Home Assistant running and accessible
- Long-lived access token (get from HA profile → Long-lived access tokens)
- [Go 1.27+](https://go.dev/dl/) — only to build the binary; the add-on install needs nothing

## Step 1: Build

```bash
cd mcp-server-extended
CGO_ENABLED=0 go build -o mcp-ha-extended ./cmd/mcp-ha-extended
```

`CGO_ENABLED=0` is what makes the result genuinely static; without it the toolchain
may link against the host's C library. To cross-compile
for another machine, set `GOOS`/`GOARCH`:

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o mcp-ha-extended ./cmd/mcp-ha-extended
```

## Step 2: Configure

Both variables are **required** — the server refuses to start without them rather
than guessing a host and sending your token to it:

```bash
export HA_URL="http://192.168.1.100:8123"
export HA_TOKEN="your_token_here"
```

## Step 3: Test

Check that the token and URL work before wiring the server into a client:

```bash
curl -H "Authorization: Bearer $HA_TOKEN" "$HA_URL/api/"
```

Should return:

```json
{"message": "API running."}
```

Then confirm the binary starts — it speaks MCP on stdin/stdout and logs to stderr:

```bash
./mcp-ha-extended
```

You should see a `serving MCP over stdio` line on stderr. Press `Ctrl+C` to stop.

## Step 4: Configure Cursor

Add to Cursor settings (JSON):

```json
{
  "mcpServers": {
    "home-assistant-automations": {
      "command": "/absolute/path/to/mcp-server-extended/mcp-ha-extended",
      "args": [],
      "env": {
        "HA_URL": "http://192.168.1.100:8123",
        "HA_TOKEN": "your_token_here"
      }
    }
  }
}
```

The command is the binary itself — there is no interpreter to point at.

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

**Server exits immediately?**

- The binary requires both `HA_URL` and `HA_TOKEN`. A missing one is reported on
  stderr as `HA_URL environment variable must be set`.

**Server not found?**

- Use an absolute path to the binary in the client config
- Make sure it is executable: `chmod +x mcp-ha-extended`

**Tools not showing?**

- Restart Cursor
- Check Cursor logs for MCP errors

## Next Steps

- [Setup Guide](SETUP.md) - Detailed setup instructions
- [Usage Examples](USAGE_EXAMPLES.md) - Code examples and use cases
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical deep dive
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
