package controllers

import (
	"go.uber.org/dig"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/infrastructure/controllers/mappers"
)

func RegisterProviders(container *dig.Container) error {
	if err := container.Provide(mappers.NewListAutomationsResponseMapper); err != nil {
		return err
	}

	return container.Provide(NewAutomationsController)
}
