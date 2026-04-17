package config

import (
	"fmt"

	"github.com/user/logslice/fieldfilter"
)

// FieldFilterConfig holds configuration for field-based filtering.
type FieldFilterConfig struct {
	Mode  string
	Key   string
	Value string
}

// DefaultFieldFilterConfig returns a disabled field filter config.
func DefaultFieldFilterConfig() FieldFilterConfig {
	return FieldFilterConfig{Mode: "none"}
}

// Validate checks the FieldFilterConfig for correctness.
func (c *FieldFilterConfig) Validate() error {
	m, err := fieldfilter.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("fieldfilter: %w", err)
	}
	if m == fieldfilter.ModeNone {
		return nil
	}
	if c.Key == "" {
		return fmt.Errorf("fieldfilter: key is required when mode is %q", c.Mode)
	}
	if c.Value == "" {
		return fmt.Errorf("fieldfilter: value is required when mode is %q", c.Mode)
	}
	return nil
}

// Enabled returns true when filtering is active.
func (c *FieldFilterConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
