package repositories

import "go.uber.org/dig"

// RegisterProviders is a no-op: this layer holds contracts only, and their
// implementations register in the infrastructure layer.
func RegisterProviders(_ *dig.Container) error {
	return nil
}
