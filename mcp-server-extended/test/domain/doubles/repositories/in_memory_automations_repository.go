package repositories

import (
	"context"
	"errors"
	"sort"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/internal/domain/entities"
)

// InMemoryAutomationsRepository is an in-memory double.
//
// It keeps real state so read-modify-write flows such as enable/disable are
// exercised end to end, rather than asserting on call sequences the way a mock
// would.
type InMemoryAutomationsRepository struct {
	automations map[string]entities.Automation
	Triggered   []string
	err         error
}

func NewInMemoryAutomationsRepository(automations ...entities.Automation) *InMemoryAutomationsRepository {
	stored := make(map[string]entities.Automation, len(automations))
	for _, automation := range automations {
		id, _ := automation["id"].(string)
		stored[id] = automation
	}

	return &InMemoryAutomationsRepository{automations: stored}
}

// WithOnError makes every operation fail, for exercising the error paths.
func (r *InMemoryAutomationsRepository) WithOnError(err error) *InMemoryAutomationsRepository {
	r.err = err

	return r
}

func (r *InMemoryAutomationsRepository) List(_ context.Context) ([]entities.Automation, error) {
	if r.err != nil {
		return nil, r.err
	}

	// Ordered by id so assertions are deterministic; Go map order is not.
	ids := make([]string, 0, len(r.automations))
	for id := range r.automations {
		ids = append(ids, id)
	}
	sort.Strings(ids)

	listed := make([]entities.Automation, 0, len(ids))
	for _, id := range ids {
		listed = append(listed, r.automations[id])
	}

	return listed, nil
}

func (r *InMemoryAutomationsRepository) Get(_ context.Context, id string) (entities.Automation, error) {
	if r.err != nil {
		return nil, r.err
	}

	automation, ok := r.automations[id]
	if !ok {
		return nil, errors.New("automation not found: " + id)
	}

	return automation, nil
}

func (r *InMemoryAutomationsRepository) Insert(_ context.Context, automation entities.Automation) (entities.Automation, error) {
	if r.err != nil {
		return nil, r.err
	}

	id, _ := automation["id"].(string)
	if id == "" {
		id = "generated"
		automation["id"] = id
	}
	r.automations[id] = automation

	return automation, nil
}

func (r *InMemoryAutomationsRepository) Update(_ context.Context, id string, automation entities.Automation) (entities.Automation, error) {
	if r.err != nil {
		return nil, r.err
	}

	r.automations[id] = automation

	return automation, nil
}

func (r *InMemoryAutomationsRepository) Delete(_ context.Context, id string) error {
	if r.err != nil {
		return r.err
	}

	delete(r.automations, id)

	return nil
}

func (r *InMemoryAutomationsRepository) Trigger(_ context.Context, id string) error {
	if r.err != nil {
		return r.err
	}

	r.Triggered = append(r.Triggered, id)

	return nil
}
