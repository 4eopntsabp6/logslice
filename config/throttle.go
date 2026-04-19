package config

import (
	"fmt"
	"time"
)

// ThrottleConfig controls per-line output throttling.
type ThrottleConfig struct {
	Mode  string
	Delay time.Duration
}

// DefaultThrottleConfig returns a ThrottleConfig with throttling disabled.
func DefaultThrottleConfig() ThrottleConfig {
	return ThrottleConfig{
		Mode:  "none",
		Delay: 0,
	}
}

// Validate checks that the throttle configuration is consistent.
func (c ThrottleConfig) Validate() error {
	switch c.Mode {
	case "", "none":
		return nil
	case "delay":
		if c.Delay <= 0 {
			return fmt.Errorf("throttle: delay mode requires a positive delay duration")
		}
		return nil
	default:
		return fmt.Errorf("throttle: unknown mode %q: want none|delay", c.Mode)
	}
}

// Enabled reports whether throttling is active.
func (c ThrottleConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
