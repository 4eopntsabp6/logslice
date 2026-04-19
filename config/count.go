package config

import "fmt"

// CountConfig controls line count limiting and skipping.
type CountConfig struct {
	Mode  string
	N     int
}

// DefaultCountConfig returns a CountConfig with no counting active.
func DefaultCountConfig() CountConfig {
	return CountConfig{Mode: "none", N: 0}
}

// Validate checks the CountConfig for logical consistency.
func (c CountConfig) Validate() error {
	switch c.Mode {
	case "", "none":
		return nil
	case "limit", "skip":
		if c.N <= 0 {
			return fmt.Errorf("count: n must be > 0 for mode %q", c.Mode)
		}
		return nil
	}
	return fmt.Errorf("count: unknown mode %q", c.Mode)
}

// Enabled returns true if counting/skipping is active.
func (c CountConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
