package config

import "fmt"

// ReverseConfig holds settings for the reverse module.
type ReverseConfig struct {
	Mode string `mapstructure:"reverse_mode"`
}

// DefaultReverseConfig returns a ReverseConfig with safe defaults.
func DefaultReverseConfig() ReverseConfig {
	return ReverseConfig{
		Mode: "none",
	}
}

// Validate checks that the mode value is acceptable.
func (c ReverseConfig) Validate() error {
	switch c.Mode {
	case "", "none", "reverse":
		return nil
	default:
		return fmt.Errorf("reverse: invalid mode %q (want none|reverse)", c.Mode)
	}
}

// Enabled reports whether reversal is active.
func (c ReverseConfig) Enabled() bool {
	return c.Mode == "reverse"
}
