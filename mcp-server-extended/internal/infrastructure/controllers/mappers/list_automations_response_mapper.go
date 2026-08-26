package mappers

import (
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers/responses"
)

// ListAutomationsResponseMapper converts automations into the listing payload.
type ListAutomationsResponseMapper struct{}

func NewListAutomationsResponseMapper() *ListAutomationsResponseMapper {
	return &ListAutomationsResponseMapper{}
}

func (m ListAutomationsResponseMapper) MapToResponse(automation entities.Automation) responses.AutomationSummaryResponse {
	return responses.AutomationSummaryResponse{
		ID:          automation.ID(),
		Alias:       automation.Alias(),
		Enabled:     automation.IsEnabled(),
		Description: automation.Description(),
	}
}

func (m ListAutomationsResponseMapper) MapToResponses(automations []entities.Automation) responses.ListAutomationsResponse {
	summaries := make([]responses.AutomationSummaryResponse, 0, len(automations))
	for _, automation := range automations {
		summaries = append(summaries, m.MapToResponse(automation))
	}

	return responses.ListAutomationsResponse{
		Count:       len(summaries),
		Automations: summaries,
	}
}
