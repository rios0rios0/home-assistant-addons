package repositories

import (
	"fmt"
	"os"

	"go.uber.org/dig"

	domain "github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/repositories"
)

// RegisterProviders binds the Home Assistant repository to its domain contract.
//
// Neither environment variable has a default. HA_URL previously fell back to
// "http://homeassistant.local:8123", which hardcoded an unencrypted endpoint
// and silently aimed requests — bearer token included — at a guessed host. The
// add-on always supplies both from config.yaml.
func RegisterProviders(container *dig.Container) error {
	return container.Provide(func() (domain.AutomationsRepository, error) {
		baseURL := os.Getenv("HA_URL")
		if baseURL == "" {
			return nil, fmt.Errorf("HA_URL environment variable must be set")
		}

		token := os.Getenv("HA_TOKEN")
		if token == "" {
			return nil, fmt.Errorf("HA_TOKEN environment variable must be set")
		}

		return NewHomeAssistantAutomationsRepository(baseURL, token, nil), nil
	})
}
