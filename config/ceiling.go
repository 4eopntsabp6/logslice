package config

import (
	"fmt"

	"github.com/yourorg/logslice/ceiling"
)

// CeilingConfig holds configuration for the ceiling filter.
type CeilingConfig struct {
	Mode    string `mapstructure:"ceiling-mode"`
	Pattern string `mapstructure:"ceiling-pattern"`
	Max     int    `mapstructure:"ceiling-max"`
}

// DefaultCeilingConfig returns a CeilingConfig with safe defaults.
func DefaultCeilingConfig() CeilingConfig {
	return CeilingConfig{
		Mode:    "none",
		Pattern: "",
		Max:     0,
	}
}

// Validate checks that the CeilingConfig fields are consistent.
func (c CeilingConfig) Validate() error {
	mode, err := ceiling.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if mode == ceiling.ModeCap && c.Max < 1 {
		return fmt.Errorf("ceiling: max must be >= 1 when mode is \"cap\"")
	}
	return nil
}

// Enabled returns true when the ceiling filter is active.
func (c CeilingConfig) Enabled() bool {
	mode, err := ceiling.ParseMode(c.Mode)
	if err != nil {
		return false
	}
	return mode != ceiling.ModeNone
}
