package builders

import "github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"

// AutomationBuilder constructs automations for tests.
//
// The zero-value build deliberately omits `enabled`, which is the shape Home
// Assistant returns for an automation that has never been toggled.
type AutomationBuilder struct {
	automation entities.Automation
}

func NewAutomationBuilder() *AutomationBuilder {
	return &AutomationBuilder{automation: entities.Automation{}}
}

func (b *AutomationBuilder) WithID(id string) *AutomationBuilder {
	b.automation["id"] = id

	return b
}

func (b *AutomationBuilder) WithAlias(alias string) *AutomationBuilder {
	b.automation["alias"] = alias

	return b
}

func (b *AutomationBuilder) WithDescription(description string) *AutomationBuilder {
	b.automation["description"] = description

	return b
}

func (b *AutomationBuilder) WithEnabled(enabled bool) *AutomationBuilder {
	b.automation["enabled"] = enabled

	return b
}

func (b *AutomationBuilder) WithTriggers(triggers ...any) *AutomationBuilder {
	b.automation["triggers"] = triggers

	return b
}

func (b *AutomationBuilder) Build() entities.Automation {
	built := entities.Automation{}
	for key, value := range b.automation {
		built[key] = value
	}

	return built
}
