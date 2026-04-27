package config

import (
	"fmt"
	"time"

	"github.com/robinovitch61/logslice/window"
)

// WindowConfig holds configuration for the sliding window filter.
type WindowConfig struct {
	Mode     string
	Duration time.Duration
}

// DefaultWindowConfig returns a WindowConfig with the window disabled.
func DefaultWindowConfig() WindowConfig {
	return WindowConfig{
		Mode:     "none",
		Duration: 0,
	}
}

// Validate checks that the WindowConfig is self-consistent.
func (c WindowConfig) Validate() error {
	m, err := window.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if m == window.ModeSliding && c.Duration <= 0 {
		return fmt.Errorf("window: sliding mode requires a positive duration")
	}
	return nil
}

// Enabled returns true when the window filter is active.
func (c WindowConfig) Enabled() bool {
	m, err := window.ParseMode(c.Mode)
	if err != nil {
		return false
	}
	return m != window.ModeNone
}
