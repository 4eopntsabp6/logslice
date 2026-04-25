package config

import (
	"fmt"
	"strings"
)

// StripConfig controls ANSI and whitespace stripping behaviour.
type StripConfig struct {
	Mode string `mapstructure:"strip_mode"`
}

// DefaultStripConfig returns a StripConfig with stripping disabled.
func DefaultStripConfig() StripConfig {
	return StripConfig{
		Mode: "none",
	}
}

// Validate checks that the Mode value is one of the accepted options.
func (c *StripConfig) Validate() error {
	switch strings.ToLower(c.Mode) {
	case "", "none", "ansi", "whitespace", "both":
		return nil
	}
	return fmt.Errorf("config: invalid strip mode %q (want none|ansi|whitespace|both)", c.Mode)
}

// Enabled reports whether the mode is anything other than "none".
func (c *StripConfig) Enabled() bool {
	return strings.ToLower(c.Mode) != "none" && c.Mode != ""
}
