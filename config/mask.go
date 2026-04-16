package config

import "fmt"

// MaskConfig controls log field masking behaviour.
type MaskConfig struct {
	Enabled bool
	Mode    string
	Pattern string
}

// DefaultMaskConfig returns a MaskConfig with masking disabled.
func DefaultMaskConfig() MaskConfig {
	return MaskConfig{
		Enabled: false,
		Mode:    "none",
		Pattern: "",
	}
}

// Validate checks that the MaskConfig is consistent.
func (c *MaskConfig) Validate() error {
	if !c.Enabled {
		return nil
	}
	switch c.Mode {
	case "redact", "partial":
	case "", "none":
		return fmt.Errorf("mask: mode must be set when masking is enabled")
	default:
		return fmt.Errorf("mask: unknown mode %q", c.Mode)
	}
	if c.Pattern == "" {
		return fmt.Errorf("mask: pattern is required when masking is enabled")
	}
	return nil
}

// MaskEnabled reports whether active masking is configured.
func (c *MaskConfig) MaskEnabled() bool {
	return c.Enabled && c.Mode != "none" && c.Mode != ""
}
