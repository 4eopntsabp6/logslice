package config

import "fmt"

// TruncateConfig holds configuration for line truncation.
type TruncateConfig struct {
	Mode  string `mapstructure:"truncate-mode"`
	Limit int    `mapstructure:"truncate-limit"`
}

// DefaultTruncateConfig returns a TruncateConfig with safe defaults.
func DefaultTruncateConfig() TruncateConfig {
	return TruncateConfig{
		Mode:  "none",
		Limit: 0,
	}
}

// Validate checks that the truncate configuration is coherent.
func (t TruncateConfig) Validate() error {
	switch t.Mode {
	case "none":
		return nil
	case "chars", "bytes":
		if t.Limit <= 0 {
			return fmt.Errorf("truncate limit must be > 0 when mode is %q", t.Mode)
		}
		return nil
	default:
		return fmt.Errorf("invalid truncate mode %q: must be none, chars, or bytes", t.Mode)
	}
}

// Enabled reports whether truncation is active.
func (t TruncateConfig) Enabled() bool {
	return t.Mode != "none"
}
