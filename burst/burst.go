// Package burst provides a filter that detects and emits bursts of log lines
// exceeding a threshold count within a sliding time window.
package burst

import (
	"fmt"
	"time"
)

// Mode controls burst filter behaviour.
type Mode string

const (
	ModeNone   Mode = "none"
	ModeWindow Mode = "window"
)

// ParseMode parses a mode string into a Mode value.
func ParseMode(s string) (Mode, error) {
	switch Mode(s) {
	case ModeNone, ModeWindow:
		return Mode(s), nil
	}
	return "", fmt.Errorf("burst: unknown mode %q", s)
}

// Filter tracks line timestamps and suppresses output unless a burst is detected.
type Filter struct {
	mode      Mode
	threshold int
	window    time.Duration
	times     []time.Time
}

// New creates a new burst Filter. threshold is the minimum number of lines
// within window to constitute a burst. ModeNone disables filtering.
func New(mode Mode, threshold int, window time.Duration) (*Filter, error) {
	if mode == ModeNone {
		return &Filter{mode: mode}, nil
	}
	if threshold < 1 {
		return nil, fmt.Errorf("burst: threshold must be >= 1, got %d", threshold)
	}
	if window <= 0 {
		return nil, fmt.Errorf("burst: window must be positive, got %s", window)
	}
	return &Filter{mode: mode, threshold: threshold, window: window}, nil
}

// Allow records the line's timestamp and returns true if a burst is active
// (i.e. threshold lines seen within the window). Always returns true for ModeNone.
func (f *Filter) Allow(t time.Time) bool {
	if f.mode == ModeNone {
		return true
	}
	cutoff := t.Add(-f.window)
	kept := f.times[:0]
	for _, ts := range f.times {
		if ts.After(cutoff) {
			kept = append(kept, ts)
		}
	}
	kept = append(kept, t)
	f.times = kept
	return len(f.times) >= f.threshold
}
