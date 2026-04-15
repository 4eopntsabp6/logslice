package config

import (
	"errors"
	"time"
)

// Config holds all runtime configuration for a logslice run.
type Config struct {
	// Input source: file path or empty for stdin
	InputPath string

	// Output destination: file path, "stdout", or "stderr"
	OutputDest string

	// Output format: "plain", "json", or "csv"
	OutputFormat string

	// Time window filters (zero value means unbounded)
	From time.Time
	To   time.Time

	// Optional regex pattern to further filter lines
	Pattern string

	// If true, prefix each output line with its original line number
	Numbered bool

	// If true, print a summary after processing
	Summary bool
}

// Validate checks that the Config fields are consistent and returns
// a descriptive error if anything is invalid.
func (c *Config) Validate() error {
	if c.OutputFormat == "" {
		c.OutputFormat = "plain"
	}

	validFormats := map[string]bool{"plain": true, "json": true, "csv": true}
	if !validFormats[c.OutputFormat] {
		return errors.New("invalid output format: must be plain, json, or csv")
	}

	if c.OutputDest == "" {
		c.OutputDest = "stdout"
	}

	if !c.From.IsZero() && !c.To.IsZero() && c.To.Before(c.From) {
		return errors.New("--to must not be before --from")
	}

	return nil
}

// HasTimeWindow reports whether at least one time bound is set.
func (c *Config) HasTimeWindow() bool {
	return !c.From.IsZero() || !c.To.IsZero()
}
