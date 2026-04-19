package config

import (
	"fmt"
	"strings"
)

// ColumnConfig holds settings for column extraction.
type ColumnConfig struct {
	Mode      string
	Delimiter string
	Index     int
}

// DefaultColumnConfig returns a ColumnConfig with safe defaults.
func DefaultColumnConfig() ColumnConfig {
	return ColumnConfig{
		Mode:      "none",
		Delimiter: ",",
		Index:     0,
	}
}

// Validate checks that the ColumnConfig is internally consistent.
func (c ColumnConfig) Validate() error {
	switch strings.ToLower(c.Mode) {
	case "", "none":
		return nil
	case "extract":
		if c.Delimiter == "" {
			return fmt.Errorf("column: delimiter required when mode is extract")
		}
		if c.Index < 0 {
			return fmt.Errorf("column: index must be >= 0")
		}
		return nil
	}
	return fmt.Errorf("column: unknown mode %q", c.Mode)
}

// Enabled reports whether column extraction is active.
func (c ColumnConfig) Enabled() bool {
	return strings.ToLower(c.Mode) == "extract"
}
