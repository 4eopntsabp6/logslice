package config

import "fmt"

// TruncateConfig controls line truncation behaviour.
type TruncateConfig struct {
	Mode  string // none | chars | words | bytes
	Limit int
}

// DefaultTruncateConfig returns a TruncateConfig with truncation disabled.
func DefaultTruncateConfig() TruncateConfig {
	return TruncateConfig{
		Mode:  "none",
		Limit: 0,
	}
}

var validTruncateModes = map[string]bool{
	"none":  true,
	"chars": true,
	"words": true,
	"bytes": true,
}

// Validate returns an error if the TruncateConfig is misconfigured.
func (c TruncateConfig) Validate() error {
	if !validTruncateModes[c.Mode] {
		return fmt.Errorf("truncate: invalid mode %q (valid: none, chars, words, bytes)", c.Mode)
	}
	if c.Mode != "none" && c.Limit <= 0 {
		return fmt.Errorf("truncate: limit must be > 0 when mode is %q", c.Mode)
	}
	return nil
}

// Enabled reports whether truncation is active.
func (c TruncateConfig) Enabled() bool {
	return c.Mode != "none"
}
