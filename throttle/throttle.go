// Package throttle provides line-output throttling by inserting a delay between emitted lines.
package throttle

import (
	"fmt"
	"time"
)

// Mode controls whether throttling is active.
type Mode int

const (
	ModeNone  Mode = iota
	ModeDelay      // insert a fixed delay between each line
)

// Throttler holds throttle configuration.
type Throttler struct {
	mode  Mode
	delay time.Duration
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "delay":
		return ModeDelay, nil
	default:
		return ModeNone, fmt.Errorf("unknown throttle mode %q: want none|delay", s)
	}
}

// New creates a Throttler. delay is ignored when mode is ModeNone.
func New(mode Mode, delay time.Duration) (*Throttler, error) {
	if mode == ModeDelay && delay <= 0 {
		return nil, fmt.Errorf("throttle: delay must be positive, got %s", delay)
	}
	return &Throttler{mode: mode, delay: delay}, nil
}

// Wait blocks for the configured delay if throttling is enabled.
func (t *Throttler) Wait() {
	if t.mode == ModeDelay {
		time.Sleep(t.delay)
	}
}

// Enabled reports whether throttling is active.
func (t *Throttler) Enabled() bool {
	return t.mode != ModeNone
}
