package config

import "fmt"

// BurstConfig controls the burst detection filter, which identifies and
// optionally suppresses lines that arrive in rapid bursts exceeding a
// configured threshold within a rolling time window.
type BurstConfig struct {
	// Mode is the burst detection mode: "none", "suppress", or "flag".
	Mode string `mapstructure:"burst_mode"`

	// Threshold is the maximum number of lines allowed within Window before
	// burst detection triggers.
	Threshold int `mapstructure:"burst_threshold"`

	// Window is the rolling duration string (e.g. "1s", "500ms") within
	// which Threshold is evaluated.
	Window string `mapstructure:"burst_window"`
}

// DefaultBurstConfig returns a BurstConfig with burst detection disabled.
func DefaultBurstConfig() BurstConfig {
	return BurstConfig{
		Mode:      "none",
		Threshold: 0,
		Window:    "",
	}
}

// Validate checks that the BurstConfig fields are consistent and valid.
func (c BurstConfig) Validate() error {
	switch c.Mode {
	case "none":
		return nil
	case "suppress", "flag":
		if c.Threshold <= 0 {
			return fmt.Errorf("burst: threshold must be greater than zero for mode %q", c.Mode)
		}
		if c.Window == "" {
			return fmt.Errorf("burst: window duration must be set for mode %q", c.Mode)
		}
		return nil
	default:
		return fmt.Errorf("burst: unknown mode %q (want none, suppress, flag)", c.Mode)
	}
}

// Enabled reports whether burst detection is active.
func (c BurstConfig) Enabled() bool {
	return c.Mode != "none"
}
