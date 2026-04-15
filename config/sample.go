package config

import (
	"fmt"

	"github.com/user/logslice/sample"
)

// SampleConfig holds configuration for log line sampling.
type SampleConfig struct {
	// Mode is the sampling strategy: "none" or "nth".
	Mode string `mapstructure:"sample_mode"`
	// N is the interval when Mode is "nth"; every Nth matched line is kept.
	N int `mapstructure:"sample_n"`
}

// DefaultSampleConfig returns a SampleConfig with sampling disabled.
func DefaultSampleConfig() SampleConfig {
	return SampleConfig{
		Mode: "none",
		N:    1,
	}
}

// Validate checks that the SampleConfig fields are consistent.
func (c *SampleConfig) Validate() error {
	mode, err := sample.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("sample: %w", err)
	}
	if mode == sample.ModeNth && c.N < 1 {
		return fmt.Errorf("sample: interval N must be >= 1 when mode is \"nth\", got %d", c.N)
	}
	return nil
}
