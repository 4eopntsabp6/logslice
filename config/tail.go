package config

import (
	"errors"
	"time"
)

// TailConfig holds configuration for real-time file following.
type TailConfig struct {
	// Enabled indicates whether tail/follow mode is active.
	Enabled bool

	// PollInterval is how frequently the file is checked for new content.
	PollInterval time.Duration

	// MaxLines, when > 0, stops following after this many lines are emitted.
	MaxLines int
}

// DefaultTailConfig returns a TailConfig with sensible defaults.
func DefaultTailConfig() TailConfig {
	return TailConfig{
		Enabled:      false,
		PollInterval: 200 * time.Millisecond,
		MaxLines:     0,
	}
}

// Validate checks that the TailConfig values are valid.
func (tc TailConfig) Validate() error {
	if !tc.Enabled {
		return nil
	}
	if tc.PollInterval <= 0 {
		return errors.New("tail: poll interval must be greater than zero")
	}
	if tc.MaxLines < 0 {
		return errors.New("tail: max-lines must be >= 0")
	}
	return nil
}
