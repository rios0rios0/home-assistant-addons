package main

import (
	"go.uber.org/dig"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal"
	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers"
)

// injectController builds the container and resolves the root type.
func injectController() (*controllers.AutomationsController, error) {
	container := dig.New()
	if err := internal.RegisterProviders(container); err != nil {
		return nil, err
	}

	var controller *controllers.AutomationsController
	if err := container.Invoke(func(c *controllers.AutomationsController) {
		controller = c
	}); err != nil {
		return nil, err
	}

	return controller, nil
}
