package entities

// Automation is a Home Assistant automation.
//
// It is carried as a decoded object rather than a fixed struct on purpose: the
// interesting parts of an automation (triggers, conditions, actions) are
// user-authored and vary per automation, and this add-on passes them through to
// the caller untouched. Only the few fields it reasons about have accessors.
//
// The accessors return `any` because Home Assistant does not commit to a type
// for them — an id may arrive as a string or a number, and an absent field must
// stay absent rather than become a zero value.
type Automation map[string]any

func (a Automation) ID() any {
	return a["id"]
}

func (a Automation) Alias() any {
	return a["alias"]
}

func (a Automation) Description() any {
	return a["description"]
}

// IsEnabled reports whether the automation is enabled.
//
// It defaults to true when the field is absent: Home Assistant omits `enabled`
// for automations that have never been toggled, and those are running.
func (a Automation) IsEnabled() any {
	enabled, ok := a["enabled"]
	if !ok {
		return true
	}

	return enabled
}

// SetEnabled marks the automation enabled or disabled.
//
// Home Assistant has no dedicated enable/disable endpoint for automations, so
// toggling one means reading it, flipping this field, and writing it back.
func (a Automation) SetEnabled(enabled bool) {
	a["enabled"] = enabled
}
