package entities

import "go.uber.org/dig"

// RegisterProviders is a no-op: entities carry no dependencies. It exists so
// every layer registers the same way.
func RegisterProviders(_ *dig.Container) error {
	return nil
}
