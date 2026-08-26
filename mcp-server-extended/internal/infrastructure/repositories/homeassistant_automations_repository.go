package repositories

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
)

const (
	defaultTimeout = 30 * time.Second

	// automationsEndpoint is the collection path; per-automation paths are
	// built from it so the literal lives in exactly one place.
	automationsEndpoint = "/automation"
)

// HomeAssistantAutomationsRepository talks to Home Assistant's REST API.
type HomeAssistantAutomationsRepository struct {
	baseURL string
	token   string
	client  *http.Client
}

// NewHomeAssistantAutomationsRepository wires a repository against the given
// instance. The trailing slash is trimmed so that joining paths never produces
// a double slash, which Home Assistant answers with a redirect.
func NewHomeAssistantAutomationsRepository(baseURL, token string, client *http.Client) *HomeAssistantAutomationsRepository {
	if client == nil {
		client = &http.Client{Timeout: defaultTimeout}
	}

	return &HomeAssistantAutomationsRepository{
		baseURL: strings.TrimRight(baseURL, "/"),
		token:   token,
		client:  client,
	}
}

func (r *HomeAssistantAutomationsRepository) List(ctx context.Context) ([]entities.Automation, error) {
	raw, err := r.call(ctx, http.MethodGet, automationsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	// Home Assistant returns a bare array; older builds wrapped it in an
	// object under "automations". Both shapes are accepted.
	switch decoded := raw.(type) {
	case []any:
		return toAutomations(decoded), nil
	case map[string]any:
		if nested, ok := decoded["automations"].([]any); ok {
			return toAutomations(nested), nil
		}
	}

	return []entities.Automation{}, nil
}

func (r *HomeAssistantAutomationsRepository) Get(ctx context.Context, id string) (entities.Automation, error) {
	raw, err := r.call(ctx, http.MethodGet, automationPath(id), nil)
	if err != nil {
		return nil, err
	}

	automation, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("automation %q: unexpected response shape %T", id, raw)
	}

	return automation, nil
}

func (r *HomeAssistantAutomationsRepository) Insert(ctx context.Context, automation entities.Automation) (entities.Automation, error) {
	raw, err := r.call(ctx, http.MethodPost, automationsEndpoint, automation)
	if err != nil {
		return nil, err
	}

	return asAutomation(raw), nil
}

func (r *HomeAssistantAutomationsRepository) Update(ctx context.Context, id string, automation entities.Automation) (entities.Automation, error) {
	raw, err := r.call(ctx, http.MethodPut, automationPath(id), automation)
	if err != nil {
		return nil, err
	}

	return asAutomation(raw), nil
}

func (r *HomeAssistantAutomationsRepository) Delete(ctx context.Context, id string) error {
	_, err := r.call(ctx, http.MethodDelete, automationPath(id), nil)

	return err
}

func (r *HomeAssistantAutomationsRepository) Trigger(ctx context.Context, id string) error {
	_, err := r.call(ctx, http.MethodPost, automationPath(id)+"/trigger", nil)

	return err
}

// call performs one authenticated request and decodes the response.
//
// A non-JSON body (Home Assistant answers several mutations with an empty 200)
// is reported as a status object rather than an error, so callers can tell
// "succeeded, nothing to say" from a genuine failure.
func (r *HomeAssistantAutomationsRepository) call(ctx context.Context, method, endpoint string, payload any) (any, error) {
	var body io.Reader

	if payload != nil {
		encoded, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("encoding request body: %w", err)
		}
		body = bytes.NewReader(encoded)
	}

	request, err := http.NewRequestWithContext(ctx, method, r.baseURL+"/api"+endpoint, body)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+r.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := r.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("calling Home Assistant: %w", err)
	}
	defer response.Body.Close()

	raw, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Home Assistant returned %s for %s %s: %s",
			response.Status, method, endpoint, strings.TrimSpace(string(raw)))
	}

	if !strings.HasPrefix(response.Header.Get("Content-Type"), "application/json") {
		return map[string]any{"status": "success", "status_code": response.StatusCode}, nil
	}

	var decoded any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return decoded, nil
}

func toAutomations(values []any) []entities.Automation {
	automations := make([]entities.Automation, 0, len(values))
	for _, value := range values {
		if automation, ok := value.(map[string]any); ok {
			automations = append(automations, automation)
		}
	}

	return automations
}

func asAutomation(raw any) entities.Automation {
	if automation, ok := raw.(map[string]any); ok {
		return automation
	}

	return entities.Automation{}
}

func automationPath(id string) string {
	return automationsEndpoint + "/" + id
}
