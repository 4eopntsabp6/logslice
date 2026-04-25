package config

import (
	"fmt"

	"github.com/logslice/logslice/squeeze"
)

// SqueezeConfig holds configuration for blank-line squeezing.
type SqueezeConfig struct {
	Mode string `mapstructure:"squeeze_mode"`
}

// DefaultSqueezeConfig returns a SqueezeConfig with squeezing disabled.
func DefaultSqueezeConfig() SqueezeConfig {
	return SqueezeConfig{
		Mode: "none",
	}
}

// Validate checks that the SqueezeConfig contains a recognised mode.
func (c SqueezeConfig) Validate() error {
	_, err := squeeze.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("config: squeeze: %w", err)
	}
	return nil
}

// Enabled returns true when squeezing is active (mode != none).
func (c SqueezeConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
