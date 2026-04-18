package config

import (
	"fmt"
)

// ContextConfig holds settings for pre/post match line context.
type ContextConfig struct {
	Enabled bool
	Mode    string // none, before, after, both
	Before  int
	After   int
}

// DefaultContextConfig returns a disabled context config.
func DefaultContextConfig() ContextConfig {
	return ContextConfig{
		Enabled: false,
		Mode:    "none",
		Before:  0,
		After:   0,
	}
}

// Validate checks the ContextConfig for consistency.
func (c *ContextConfig) Validate() error {
	if !c.Enabled {
		return nil
	}
	switch c.Mode {
	case "none", "before", "after", "both":
	default:
		return fmt.Errorf("context: invalid mode %q", c.Mode)
	}
	if c.Before < 0 {
		return fmt.Errorf("context: before must be >= 0")
	}
	if c.After < 0 {
		return fmt.Errorf("context: after must be >= 0")
	}
	if c.Mode == "before" && c.Before == 0 {
		return fmt.Errorf("context: mode 'before' requires before > 0")
	}
	if c.Mode == "after" && c.After == 0 {
		return fmt.Errorf("context: mode 'after' requires after > 0")
	}
	return nil
}

// ContextEnabled returns true when context capture is active.
func (c *ContextConfig) ContextEnabled() bool {
	return c.Enabled && c.Mode != "none"
}
