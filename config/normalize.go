package config

import (
	"fmt"

	"github.com/yourorg/logslice/normalize"
)

// NormalizeConfig holds configuration for line normalization.
type NormalizeConfig struct {
	Mode string `mapstructure:"normalize_mode"`
}

// DefaultNormalizeConfig returns a NormalizeConfig with normalization disabled.
func DefaultNormalizeConfig() NormalizeConfig {
	return NormalizeConfig{
		Mode: "none",
	}
}

// Validate checks that the mode string is a recognized normalize mode.
func (c NormalizeConfig) Validate() error {
	if _, err := normalize.ParseMode(c.Mode); err != nil {
		return fmt.Errorf("config: normalize: %w", err)
	}
	return nil
}

// Enabled reports whether normalization is active.
func (c NormalizeConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
