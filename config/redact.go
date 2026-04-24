package config

import (
	"fmt"
)

// RedactConfig controls line-level redaction of sensitive content.
type RedactConfig struct {
	Mode        string `mapstructure:"redact_mode"`
	Pattern     string `mapstructure:"redact_pattern"`
	Placeholder string `mapstructure:"redact_placeholder"`
}

// DefaultRedactConfig returns a RedactConfig with redaction disabled.
func DefaultRedactConfig() RedactConfig {
	return RedactConfig{
		Mode:        "none",
		Placeholder: "[REDACTED]",
	}
}

// Validate checks that the RedactConfig fields are consistent.
func (c *RedactConfig) Validate() error {
	switch c.Mode {
	case "", "none":
		return nil
	case "pattern":
		if c.Pattern == "" {
			return fmt.Errorf("redact: pattern required when mode is %q", c.Mode)
		}
		return nil
	default:
		return fmt.Errorf("redact: unknown mode %q", c.Mode)
	}
}

// Enabled reports whether redaction is active.
func (c *RedactConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
