//go:build unit

package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rios0rios0/home-assistant-addons/mcp-server-extended/test/domain/builders"
)

func TestAutomation(t *testing.T) {
	t.Parallel()

	t.Run("should return the listed fields when they are present", func(t *testing.T) {
		// given
		automation := builders.NewAutomationBuilder().
			WithID("1").
			WithAlias("Morning lights").
			WithDescription("turns the lights on").
			WithEnabled(false).
			Build()

		// when & then
		assert.Equal(t, "1", automation.ID())
		assert.Equal(t, "Morning lights", automation.Alias())
		assert.Equal(t, "turns the lights on", automation.Description())
		assert.Equal(t, false, automation.IsEnabled())
	})

	t.Run("should default enabled to true when the field is absent", func(t *testing.T) {
		// given
		// Home Assistant omits `enabled` for automations that have never been
		// toggled, and those are running — reporting them as disabled is wrong.
		automation := builders.NewAutomationBuilder().WithID("1").Build()

		// when
		enabled := automation.IsEnabled()

		// then
		assert.Equal(t, true, enabled)
	})

	t.Run("should report absent fields as nil so they serialise as null", func(t *testing.T) {
		// given
		automation := builders.NewAutomationBuilder().Build()

		// when & then
		assert.Nil(t, automation.ID())
		assert.Nil(t, automation.Alias())
		assert.Nil(t, automation.Description())
	})

	t.Run("should write the enabled field when set", func(t *testing.T) {
		// given
		automation := builders.NewAutomationBuilder().WithID("1").Build()

		// when
		automation.SetEnabled(true)

		// then
		assert.Equal(t, true, automation.IsEnabled())

		// when
		automation.SetEnabled(false)

		// then
		assert.Equal(t, false, automation.IsEnabled())
	})
}
