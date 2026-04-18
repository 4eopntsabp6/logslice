package config

import (
	"fmt"

	"github.com/user/logslice/severity"
)

// SeverityConfig holds CLI-bound severity filter settings.
type SeverityConfig struct {
	Mode  string `mapstructure:"severity-mode"`
	Level string `mapstructure:"severity-level"`
}

// DefaultSeverityConfig returns a SeverityConfig with sensible defaults.
func DefaultSeverityConfig() SeverityConfig {
	return SeverityConfig{
		Mode:  string(severity.ModeNone),
		Level: "info",
	}
}

// Validate checks that the SeverityConfig fields are valid.
func (c SeverityConfig) Validate() error {
	mode, err := severity.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("severity: %w", err)
	}
	if mode != severity.ModeNone {
		if _, err := severity.ParseLevel(c.Level); err != nil {
			return fmt.Errorf("severity: %w", err)
		}
	}
	return nil
}

// Enabled returns true if severity filtering is active.
func (c SeverityConfig) Enabled() bool {
	return !severity.Mode(c.Mode).IsNone()
}
