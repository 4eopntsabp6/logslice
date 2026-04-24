package config

import (
	"fmt"
	"strings"
)

// MergeConfig holds configuration for the merge stage.
type MergeConfig struct {
	Mode      string `mapstructure:"merge-mode"`
	Pattern   string `mapstructure:"merge-pattern"`
	Separator string `mapstructure:"merge-separator"`
}

// DefaultMergeConfig returns a MergeConfig with safe defaults.
func DefaultMergeConfig() MergeConfig {
	return MergeConfig{
		Mode:      "none",
		Separator: " ",
	}
}

// Validate checks that the MergeConfig is self-consistent.
func (c MergeConfig) Validate() error {
	switch strings.ToLower(c.Mode) {
	case "none", "":
		return nil
	case "continuation":
		if c.Pattern == "" {
			return fmt.Errorf("merge: pattern is required when mode is %q", c.Mode)
		}
		return nil
	default:
		return fmt.Errorf("merge: unknown mode %q", c.Mode)
	}
}

// Enabled returns true when merging is active.
func (c MergeConfig) Enabled() bool {
	return strings.ToLower(c.Mode) != "none" && c.Mode != ""
}
