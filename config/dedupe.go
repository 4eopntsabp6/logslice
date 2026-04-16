package config

import "fmt"

// DedupeConfig holds configuration for line deduplication.
type DedupeConfig struct {
	Mode  string `mapstructure:"dedupe-mode"`
	Field string `mapstructure:"dedupe-field"`
}

// DefaultDedupeConfig returns a DedupeConfig with safe defaults.
func DefaultDedupeConfig() DedupeConfig {
	return DedupeConfig{
		Mode:  "none",
		Field: "",
	}
}

// Validate checks that the dedupe configuration is coherent.
func (d DedupeConfig) Validate() error {
	switch d.Mode {
	case "none", "consecutive", "global":
		// valid
	default:
		return fmt.Errorf("invalid dedupe mode %q: must be none, consecutive, or global", d.Mode)
	}
	return nil
}

// Enabled reports whether deduplication is active.
func (d DedupeConfig) Enabled() bool {
	return d.Mode != "none"
}
