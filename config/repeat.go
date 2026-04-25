package config

import (
	"fmt"

	"github.com/yourorg/logslice/repeat"
)

// RepeatConfig holds configuration for the repeat suppression filter.
type RepeatConfig struct {
	Mode    string `mapstructure:"repeat-mode"`
	Pattern string `mapstructure:"repeat-pattern"`
	Max     int    `mapstructure:"repeat-max"`
}

// DefaultRepeatConfig returns a RepeatConfig with safe defaults.
func DefaultRepeatConfig() RepeatConfig {
	return RepeatConfig{
		Mode:    "none",
		Pattern: "",
		Max:     0,
	}
}

// Validate checks that the RepeatConfig fields are consistent.
func (c RepeatConfig) Validate() error {
	mode, err := repeat.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("repeat: %w", err)
	}
	if mode == repeat.ModeNone {
		return nil
	}
	if c.Pattern == "" {
		return fmt.Errorf("repeat: pattern is required when mode is %q", c.Mode)
	}
	if c.Max <= 0 {
		return fmt.Errorf("repeat: max must be > 0 when mode is %q", c.Mode)
	}
	return nil
}

// Enabled returns true when the repeat filter is active.
func (c RepeatConfig) Enabled() bool {
	mode, err := repeat.ParseMode(c.Mode)
	if err != nil {
		return false
	}
	return mode != repeat.ModeNone
}
