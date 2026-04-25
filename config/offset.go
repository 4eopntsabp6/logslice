package config

import (
	"fmt"

	"github.com/robinovitch61/logslice/offset"
)

// OffsetConfig holds configuration for the offset filter.
type OffsetConfig struct {
	Mode string
	N    int
}

// DefaultOffsetConfig returns an OffsetConfig with offset disabled.
func DefaultOffsetConfig() OffsetConfig {
	return OffsetConfig{
		Mode: "none",
		N:    0,
	}
}

// Validate checks that the OffsetConfig fields are consistent.
func (c OffsetConfig) Validate() error {
	mode, err := offset.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if mode != offset.ModeNone && c.N <= 0 {
		return fmt.Errorf("offset: n must be > 0 when mode is %q", c.Mode)
	}
	return nil
}

// Enabled returns true when the offset filter is active.
func (c OffsetConfig) Enabled() bool {
	mode, err := offset.ParseMode(c.Mode)
	if err != nil {
		return false
	}
	return mode != offset.ModeNone
}
