package repositories

import (
	"context"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
)

// AutomationsRepository abstracts access to Home Assistant's automations.
//
// Enabling and disabling an automation are deliberately absent: Home Assistant
// exposes no endpoint for either, so they are read-modify-write compositions of
// Get and Update and belong to the caller, not to this contract.
type AutomationsRepository interface {
	List(ctx context.Context) ([]entities.Automation, error)
	Get(ctx context.Context, id string) (entities.Automation, error)
	Insert(ctx context.Context, automation entities.Automation) (entities.Automation, error)
	Update(ctx context.Context, id string, automation entities.Automation) (entities.Automation, error)
	Delete(ctx context.Context, id string) error
	Trigger(ctx context.Context, id string) error
}
