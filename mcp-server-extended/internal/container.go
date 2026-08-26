package internal

import (
	"go.uber.org/dig"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
	domainRepositories "github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/repositories"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers"
	infraRepositories "github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/repositories"
)

// RegisterProviders registers every layer bottom-up, so Dig can resolve each
// constructor's parameters from providers registered before it.
func RegisterProviders(container *dig.Container) error {
	if err := infraRepositories.RegisterProviders(container); err != nil {
		return err
	}
	if err := domainRepositories.RegisterProviders(container); err != nil {
		return err
	}
	if err := entities.RegisterProviders(container); err != nil {
		return err
	}

	return controllers.RegisterProviders(container)
}
