package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"gopkg.in/yaml.v3"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/repositories"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers/mappers"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers/responses"
)

// AutomationsController exposes Home Assistant automations as MCP tools.
type AutomationsController struct {
	repository repositories.AutomationsRepository
	mapper     *mappers.ListAutomationsResponseMapper
}

func NewAutomationsController(
	repository repositories.AutomationsRepository,
	mapper *mappers.ListAutomationsResponseMapper,
) *AutomationsController {
	return &AutomationsController{repository: repository, mapper: mapper}
}

// Tool argument shapes. The `jsonschema` tags become the tool input schema the
// SDK advertises, so the descriptions here are what an MCP client shows.
type (
	emptyArgs struct{}

	automationIDArgs struct {
		AutomationID string `json:"automation_id" jsonschema:"The automation ID"`
	}

	createAutomationArgs struct {
		AutomationYAML string `json:"automation_yaml" jsonschema:"YAML configuration for the automation (single automation object)"`
		Alias          string `json:"alias,omitempty" jsonschema:"Optional alias/name for the automation"`
	}

	updateAutomationArgs struct {
		AutomationID   string `json:"automation_id" jsonschema:"The automation ID to update"`
		AutomationYAML string `json:"automation_yaml" jsonschema:"Updated YAML configuration for the automation"`
	}
)

// Register wires every tool onto the server.
//
// Each handler declares `any` as its output type, which tells the SDK to omit
// an output schema and pass the text content through untouched — the tools
// answer with formatted JSON, as the previous implementation did.
func (c *AutomationsController) Register(server *mcp.Server) {
	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_automations",
		Description: "List all automations in Home Assistant",
	}, c.List)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_automation",
		Description: "Get details of a specific automation by ID",
	}, c.Get)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "create_automation",
		Description: "Create a new automation from YAML configuration",
	}, c.Insert)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "update_automation",
		Description: "Update an existing automation",
	}, c.Update)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "delete_automation",
		Description: "Delete an automation",
	}, c.Delete)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "trigger_automation",
		Description: "Manually trigger an automation",
	}, c.Trigger)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "enable_automation",
		Description: "Enable an automation",
	}, c.Enable)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "disable_automation",
		Description: "Disable an automation",
	}, c.Disable)
}

func (c *AutomationsController) List(ctx context.Context, _ *mcp.CallToolRequest, _ emptyArgs) (*mcp.CallToolResult, any, error) {
	automations, err := c.repository.List(ctx)
	if err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(c.mapper.MapToResponses(automations)), nil, nil
}

func (c *AutomationsController) Get(ctx context.Context, _ *mcp.CallToolRequest, args automationIDArgs) (*mcp.CallToolResult, any, error) {
	automation, err := c.repository.Get(ctx, args.AutomationID)
	if err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(automation), nil, nil
}

func (c *AutomationsController) Insert(ctx context.Context, _ *mcp.CallToolRequest, args createAutomationArgs) (*mcp.CallToolResult, any, error) {
	automation, err := decodeAutomation(args.AutomationYAML)
	if err != nil {
		return errorResult(err), nil, nil
	}

	// `alias` is accepted for convenience: callers routinely name the
	// automation outside the YAML body, and dropping it silently would lose
	// the name. YAML already carrying an alias wins.
	if args.Alias != "" {
		if _, ok := automation["alias"]; !ok {
			automation["alias"] = args.Alias
		}
	}

	created, err := c.repository.Insert(ctx, automation)
	if err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(responses.ResultResponse{Status: "created", Result: created}), nil, nil
}

func (c *AutomationsController) Update(ctx context.Context, _ *mcp.CallToolRequest, args updateAutomationArgs) (*mcp.CallToolResult, any, error) {
	automation, err := decodeAutomation(args.AutomationYAML)
	if err != nil {
		return errorResult(err), nil, nil
	}

	updated, err := c.repository.Update(ctx, args.AutomationID, automation)
	if err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(responses.ResultResponse{Status: "updated", Result: updated}), nil, nil
}

func (c *AutomationsController) Delete(ctx context.Context, _ *mcp.CallToolRequest, args automationIDArgs) (*mcp.CallToolResult, any, error) {
	if err := c.repository.Delete(ctx, args.AutomationID); err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(responses.StatusResponse{Status: "deleted", AutomationID: args.AutomationID}), nil, nil
}

func (c *AutomationsController) Trigger(ctx context.Context, _ *mcp.CallToolRequest, args automationIDArgs) (*mcp.CallToolResult, any, error) {
	if err := c.repository.Trigger(ctx, args.AutomationID); err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(responses.StatusResponse{Status: "triggered", AutomationID: args.AutomationID}), nil, nil
}

func (c *AutomationsController) Enable(ctx context.Context, _ *mcp.CallToolRequest, args automationIDArgs) (*mcp.CallToolResult, any, error) {
	return c.setEnabled(ctx, args.AutomationID, true, "enabled")
}

func (c *AutomationsController) Disable(ctx context.Context, _ *mcp.CallToolRequest, args automationIDArgs) (*mcp.CallToolResult, any, error) {
	return c.setEnabled(ctx, args.AutomationID, false, "disabled")
}

func (c *AutomationsController) setEnabled(ctx context.Context, id string, enabled bool, status string) (*mcp.CallToolResult, any, error) {
	current, err := c.repository.Get(ctx, id)
	if err != nil {
		return errorResult(err), nil, nil
	}

	current.SetEnabled(enabled)

	if _, err := c.repository.Update(ctx, id, current); err != nil {
		return errorResult(err), nil, nil
	}

	return jsonResult(responses.StatusResponse{Status: status, AutomationID: id}), nil, nil
}

func decodeAutomation(document string) (entities.Automation, error) {
	var automation entities.Automation
	if err := yaml.Unmarshal([]byte(document), &automation); err != nil {
		return nil, fmt.Errorf("parsing automation YAML: %w", err)
	}

	if automation == nil {
		return nil, fmt.Errorf("parsing automation YAML: document is empty")
	}

	return automation, nil
}

// jsonResult renders a value as the indented JSON text the tools answer with.
func jsonResult(value any) *mcp.CallToolResult {
	encoded, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return errorResult(fmt.Errorf("encoding result: %w", err))
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: string(encoded)}},
	}
}

// errorResult reports a failure to the caller as tool output rather than as a
// protocol error, so an MCP client shows the reason instead of a generic fault.
func errorResult(err error) *mcp.CallToolResult {
	encoded, marshalErr := json.MarshalIndent(responses.ErrorResponse{Error: err.Error()}, "", "  ")
	if marshalErr != nil {
		encoded = []byte(`{"error": "` + err.Error() + `"}`)
	}

	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{&mcp.TextContent{Text: string(encoded)}},
	}
}
