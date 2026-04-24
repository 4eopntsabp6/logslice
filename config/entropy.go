package config

import (
	"fmt"

	"github.com/yourorg/logslice/entropy"
)

// EntropyConfig holds settings for the entropy filter stage.
type EntropyConfig struct {
	Mode      string  `mapstructure:"entropy-mode"`
	Threshold float64 `mapstructure:"entropy-threshold"`
}

// DefaultEntropyConfig returns an EntropyConfig with filtering disabled.
func DefaultEntropyConfig() EntropyConfig {
	return EntropyConfig{
		Mode:      "none",
		Threshold: 0,
	}
}

// Validate checks that Mode is recognised and Threshold is non-negative when
// the filter is active.
func (c EntropyConfig) Validate() error {
	mode, err := entropy.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if mode != 0 && c.Threshold < 0 {
		return fmt.Errorf("entropy: threshold must be >= 0")
	}
	return nil
}

// Enabled returns true when the entropy filter is active.
func (c EntropyConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
