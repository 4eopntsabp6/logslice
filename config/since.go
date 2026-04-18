package config

import (
	"fmt"
	"time"
)

// SinceConfig holds configuration for the since (relative time) filter.
type SinceConfig struct {
	// Duration is a human-readable duration string e.g. "1h", "30m".
	Duration string
}

// DefaultSinceConfig returns a SinceConfig with the filter disabled.
func DefaultSinceConfig() SinceConfig {
	return SinceConfig{Duration: ""}
}

// Validate checks that the duration is parseable and positive.
func (c SinceConfig) Validate() error {
	if c.Duration == "" {
		return nil
	}
	d, err := time.ParseDuration(c.Duration)
	if err != nil {
		return fmt.Errorf("since: invalid duration %q: %w", c.Duration, err)
	}
	if d <= 0 {
		return fmt.Errorf("since: duration must be positive, got %q", c.Duration)
	}
	return nil
}

// Enabled reports whether the since filter is active.
func (c SinceConfig) Enabled() bool {
	return c.Duration != ""
}
