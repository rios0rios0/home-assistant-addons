# Usage Examples

This document provides practical examples of using the MCP Server Extended tools.

For setup instructions, see:
- [Quick Start Guide](QUICK_START.md) - Fast setup
- [Setup Guide](SETUP.md) - Detailed setup
- [Addon Installation](ADDON_INSTALLATION.md) - Install as addon

## Example 1: List All Automations

```text
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

```text
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

```text
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

```text
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

```text
# Disable
disable_automation(automation_id="automation.morning_routine")

# Enable
enable_automation(automation_id="automation.morning_routine")
```

## Example 6: Trigger Automation Manually

```text
# Tool call
trigger_automation(automation_id="automation.morning_routine")

# Response
{
  "status": "triggered",
  "automation_id": "automation.morning_routine"
}
```

## Example 7: Delete Automation

```text
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

Example Go program for bulk enable, using the same repository the add-on uses:

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/repositories"
)

func main() {
	repository := repositories.NewHomeAssistantAutomationsRepository(
		os.Getenv("HA_URL"), os.Getenv("HA_TOKEN"), nil,
	)

	ctx := context.Background()

	automations, err := repository.List(ctx)
	if err != nil {
		panic(err)
	}

	for _, automation := range automations {
		if automation.IsEnabled() == true {
			continue
		}

		id, _ := automation.ID().(string)

		// Read the automation back before writing: enabling is a
		// read-modify-write, and the rest of the body has to survive it.
		current, err := repository.Get(ctx, id)
		if err != nil {
			panic(err)
		}

		current.SetEnabled(true)

		if _, err := repository.Update(ctx, id, current); err != nil {
			panic(err)
		}

		fmt.Printf("Enabled %s\n", id)
	}
}
```

## Example 9: Import from YAML Files

You can write a helper program to import all YAML files:

```go
package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/repositories"
)

func main() {
	repository := repositories.NewHomeAssistantAutomationsRepository(
		os.Getenv("HA_URL"), os.Getenv("HA_TOKEN"), nil,
	)

	matches, err := filepath.Glob(filepath.Join("../automations", "*.yaml"))
	if err != nil {
		panic(err)
	}

	for _, path := range matches {
		document, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		// Each file may hold a single automation or a list of them.
		var automations []entities.Automation
		if err := yaml.Unmarshal(document, &automations); err != nil {
			var single entities.Automation
			if err := yaml.Unmarshal(document, &single); err != nil {
				panic(err)
			}
			automations = []entities.Automation{single}
		}

		for _, automation := range automations {
			if _, err := repository.Insert(context.Background(), automation); err != nil {
				panic(err)
			}

			alias, ok := automation.Alias().(string)
			if !ok {
				alias = "unnamed"
			}
			fmt.Printf("Imported %s\n", alias)
		}
	}
}
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
