package config

import (
	"fmt"
	"strings"
)

// TransformConfig holds configuration for line transformation.
type TransformConfig struct {
	Mode string `mapstructure:"transform"`
}

// DefaultTransformConfig returns a TransformConfig with safe defaults.
func DefaultTransformConfig() TransformConfig {
	return TransformConfig{Mode: "none"}
}

// Validate checks that the transform mode is recognised.
func (c TransformConfig) Validate() error {
	switch strings.ToLower(c.Mode) {
	case "", "none", "upper", "lower", "trim":
		return nil
	default:
		return fmt.Errorf("invalid transform mode: %q", c.Mode)
	}
}

// Enabled reports whether transformation is active.
func (c TransformConfig) Enabled() bool {
	return strings.ToLower(c.Mode) != "none" && c.Mode != ""
}
