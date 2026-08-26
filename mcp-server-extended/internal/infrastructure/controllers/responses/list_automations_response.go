package responses

// AutomationSummaryResponse is the reduced view of an automation that listing
// returns. JSON tags live here rather than on the entity: the entity is
// framework-agnostic, and this is the wire shape.
//
// The fields are `any` so that an automation missing one serialises it as JSON
// null, which is what Home Assistant itself reports for an absent value.
type AutomationSummaryResponse struct {
	ID          any `json:"id"`
	Alias       any `json:"alias"`
	Enabled     any `json:"enabled"`
	Description any `json:"description"`
}

// ListAutomationsResponse is the payload the list_automations tool answers with.
type ListAutomationsResponse struct {
	Count       int                         `json:"count"`
	Automations []AutomationSummaryResponse `json:"automations"`
}

// StatusResponse acknowledges a mutation that has no body of its own.
type StatusResponse struct {
	Status       string `json:"status"`
	AutomationID string `json:"automation_id"`
}

// ResultResponse acknowledges a mutation that echoes the stored automation.
type ResultResponse struct {
	Status string `json:"status"`
	Result any    `json:"result"`
}

// ErrorResponse reports a failure as tool output rather than a protocol fault,
// so an MCP client shows the reason instead of a generic error.
type ErrorResponse struct {
	Error string `json:"error"`
}
