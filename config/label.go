package config

import (
	"fmt"

	"github.com/user/logslice/label"
)

// LabelConfig holds configuration for line labelling.
type LabelConfig struct {
	Mode    string
	Label   string
	Pattern string
}

// DefaultLabelConfig returns a LabelConfig with labelling disabled.
func DefaultLabelConfig() LabelConfig {
	return LabelConfig{Mode: "none"}
}

// Validate checks that the LabelConfig fields are consistent.
func (c LabelConfig) Validate() error {
	m, err := label.ParseMode(c.Mode)
	if err != nil {
		return fmt.Errorf("label: %w", err)
	}
	if m == label.ModeNone {
		return nil
	}
	if c.Label == "" {
		return fmt.Errorf("label: label text required when mode is %q", c.Mode)
	}
	if c.Pattern == "" {
		return fmt.Errorf("label: pattern required when mode is %q", c.Mode)
	}
	return nil
}

// Enabled returns true when labelling is active.
func (c LabelConfig) Enabled() bool {
	m, _ := label.ParseMode(c.Mode)
	return m != label.ModeNone
}
