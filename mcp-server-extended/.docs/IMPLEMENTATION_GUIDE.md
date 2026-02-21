# Implementation Guide: Adding Automation Tools to MCP

## Overview

This guide explains the different approaches to adding automation management capabilities to your MCP setup.

## Approach Comparison

### Option 1: Standalone MCP Server (Current Implementation)
**Pros:**
- ✅ No need to modify existing code
- ✅ Can run alongside existing HA MCP server
- ✅ Easy to test and debug
- ✅ Can be version controlled separately

**Cons:**
- ❌ Requires running two MCP servers
- ❌ Slightly more complex configuration

**Best for:** Quick implementation, testing, or when you don't have access to modify the existing server.

### Option 2: Extend Existing MCP Server
**Pros:**
- ✅ Single server to manage
- ✅ Unified configuration
- ✅ Better performance (one process)

**Cons:**
- ❌ Requires access to source code
- ❌ Need to understand existing codebase
- ❌ May need to coordinate with maintainers

**Best for:** Long-term solution, if you control the server code.

### Option 3: Fork and Extend
**Pros:**
- ✅ Full control over features
- ✅ Can customize as needed
- ✅ Can contribute back upstream

**Cons:**
- ❌ Need to maintain fork
- ❌ May diverge from upstream

**Best for:** When you want to add many custom features.

## Implementation Details

### Home Assistant REST API

Home Assistant exposes automations via REST API:

```python
# List all automations
GET /api/automation
Authorization: Bearer {token}

# Get specific automation
GET /api/automation/{automation_id}

# Create automation
POST /api/automation
Content-Type: application/json
Body: { automation object }

# Update automation
PUT /api/automation/{automation_id}
Content-Type: application/json
Body: { automation object }

# Delete automation
DELETE /api/automation/{automation_id}

# Trigger automation
POST /api/automation/{automation_id}/trigger
```

### MCP Tool Definition

Each tool follows this pattern:

```python
Tool(
    name="tool_name",
    description="What the tool does",
    inputSchema={
        "type": "object",
        "properties": {
            "param_name": {
                "type": "string",
                "description": "Parameter description"
            }
        },
        "required": ["param_name"]
    }
)
```

### Error Handling

The server should handle:
- Network errors (connection refused, timeout)
- Authentication errors (401, 403)
- Not found errors (404)
- Validation errors (400)
- Server errors (500)

## Integration with Existing Tools

If extending the existing Home Assistant MCP server, you would:

1. **Find the tool registration** - Look for `@server.list_tools()` or similar
2. **Add new tools** - Add automation tools alongside device control tools
3. **Reuse HTTP client** - Use existing API client if available
4. **Follow patterns** - Match existing code style and error handling

## Testing Strategy

1. **Unit Tests**: Test API calls in isolation
2. **Integration Tests**: Test with real Home Assistant instance
3. **E2E Tests**: Test through MCP client (Cursor)

Example test:

```python
async def test_list_automations():
    result = await ha_api_call("GET", "/automation")
    assert isinstance(result, list)
    assert len(result) >= 0
```

## Security Considerations

1. **Token Storage**: Never commit tokens to git
2. **Token Permissions**: Use tokens with minimal required permissions
3. **Input Validation**: Validate all YAML before sending to HA
4. **Error Messages**: Don't expose sensitive info in errors
5. **Rate Limiting**: Consider rate limiting for API calls

## Performance Optimization

1. **Caching**: Cache automation list (with TTL)
2. **Batching**: Batch multiple operations when possible
3. **Async**: Use async/await for all I/O operations
4. **Connection Pooling**: Reuse HTTP connections

## Deployment

### Development
```bash
python server.py
```

### Production
- Use process manager (systemd, supervisor)
- Add logging
- Monitor health
- Set up alerts

### Docker (Optional)
```dockerfile
FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY server.py .
CMD ["python", "server.py"]
```

## Next Steps

1. **Test the standalone server** - Use `test_api.py` to verify connectivity
2. **Configure Cursor** - Add server to MCP configuration
3. **Try creating an automation** - Use the new tools
4. **Iterate** - Add more features as needed

## Contributing Back

If you extend the official Home Assistant MCP server:

1. Check if there's a GitHub repo
2. Look for contribution guidelines
3. Create a feature branch
4. Submit a pull request with:
   - Clear description
   - Tests
   - Documentation updates

## Related Documentation

- [Setup Guide](SETUP.md) - Development setup instructions
- [PDM Setup](PDM_SETUP.md) - Python dependency management
- [Usage Examples](USAGE_EXAMPLES.md) - Code examples
- [Addon Structure](ADDON_STRUCTURE.md) - Addon architecture
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
