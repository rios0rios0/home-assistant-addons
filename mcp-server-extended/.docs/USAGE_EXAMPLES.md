# Usage Examples

This document provides practical examples of using the MCP Server Extended tools.

For setup instructions, see:
- [Quick Start Guide](QUICK_START.md) - Fast setup
- [Setup Guide](SETUP.md) - Detailed setup
- [Addon Installation](ADDON_INSTALLATION.md) - Install as addon

## Example 1: List All Automations

```python
# Tool call
list_automations()

# Response
{
  "count": 5,
  "automations": [
    {
      "id": "automation.morning_routine",
      "alias": "Morning Routine - Turn On Key Lights",
      "enabled": true,
      "description": "Automatically turns on lights..."
    },
    ...
  ]
}
```

## Example 2: Create Automation from YAML

```python
# Tool call
create_automation(
  automation_yaml="""
- id: 'test_automation'
  alias: 'Test Automation'
  description: 'A test automation'
  trigger:
    - platform: time
      at: '08:00:00'
  action:
    - service: light.turn_on
      target:
        entity_id: light.office_bulb
      data:
        brightness: 255
  mode: single
  """
)

# Response
{
  "status": "created",
  "result": {
    "id": "automation.test_automation",
    ...
  }
}
```

## Example 3: Update Existing Automation

```python
# Tool call
update_automation(
  automation_id="automation.morning_routine",
  automation_yaml="""
- id: 'morning_routine'
  alias: 'Morning Routine - Updated'
  trigger:
    - platform: sun
      event: sunrise
      offset: '+00:45:00'  # Changed from 30 to 45 minutes
  action:
    - service: light.turn_on
      target:
        entity_id:
          - light.kitchen_bulb
          - light.entrance_bulb
      data:
        brightness: 200  # Changed brightness
  mode: single
  """
)
```

## Example 4: Get Automation Details

```python
# Tool call
get_automation(automation_id="automation.morning_routine")

# Response
{
  "id": "automation.morning_routine",
  "alias": "Morning Routine - Turn On Key Lights",
  "description": "...",
  "enabled": true,
  "trigger": [...],
  "action": [...],
  "condition": [...],
  "mode": "single"
}
```

## Example 5: Enable/Disable Automation

```python
# Disable
disable_automation(automation_id="automation.morning_routine")

# Enable
enable_automation(automation_id="automation.morning_routine")
```

## Example 6: Trigger Automation Manually

```python
# Tool call
trigger_automation(automation_id="automation.morning_routine")

# Response
{
  "status": "triggered",
  "automation_id": "automation.morning_routine"
}
```

## Example 7: Delete Automation

```python
# Tool call
delete_automation(automation_id="automation.test_automation")

# Response
{
  "status": "deleted",
  "automation_id": "automation.test_automation"
}
```

## Example 8: Bulk Operations

While the MCP server doesn't have bulk operations built-in, you can:

1. **List all automations** to get IDs
2. **Loop through** and call individual operations
3. **Or extend the server** to add bulk operations

Example Python script for bulk enable:

```python
import asyncio
from server import ha_api_call

async def enable_all_automations():
    automations = await ha_api_call("GET", "/automation")
    for auto in automations:
        auto_id = auto.get("id")
        if not auto.get("enabled"):
            current = await ha_api_call("GET", f"/automation/{auto_id}")
            current["enabled"] = True
            await ha_api_call("PUT", f"/automation/{auto_id}", current)
            print(f"Enabled {auto_id}")

asyncio.run(enable_all_automations())
```

## Example 9: Import from YAML Files

You can create a helper script to import all YAML files:

```python
import os
import yaml
import asyncio
from server import ha_api_call

async def import_automations_from_directory(directory):
    for filename in os.listdir(directory):
        if filename.endswith('.yaml'):
            with open(os.path.join(directory, filename), 'r') as f:
                automations = yaml.safe_load(f)
                if isinstance(automations, list):
                    for automation in automations:
                        result = await ha_api_call("POST", "/automation", automation)
                        print(f"Imported {automation.get('alias', 'unnamed')}")

# Usage
asyncio.run(import_automations_from_directory("../automations"))
```

## Example 10: Using with Cursor AI

Once configured, you can ask Cursor:

```
"List all my automations"
"Create an automation that turns on the office light at 9 AM"
"Update the morning routine to start 15 minutes later"
"Disable the evening routine automation"
"Show me the details of the bedroom automation"
```

Cursor will use the MCP tools automatically!

## Related Documentation

- [Quick Start Guide](QUICK_START.md) - Fast setup guide
- [Setup Guide](SETUP.md) - Detailed setup instructions
- [Implementation Guide](IMPLEMENTATION_GUIDE.md) - Technical details
- [Addon Installation](ADDON_INSTALLATION.md) - Install as Home Assistant addon
- [Documentation Index](SUMMARY.md) - Complete documentation navigation
