// Package since provides filtering of log lines by a relative time duration.
package since

import (
	"fmt"
	"time"
)

// ParseMode parses a duration string like "1h", "30m", "2h30m".
func ParseMode(s string) (time.Duration, error) {
	if s == "" {
		return 0, nil
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, fmt.Errorf("since: invalid duration %q: %w", s, err)
	}
	if d <= 0 {
		return 0, fmt.Errorf("since: duration must be positive, got %q", s)
	}
	return d, nil
}

// Filter drops log lines whose timestamp is older than the cutoff.
type Filter struct {
	cutoff time.Time
	enabled bool
}

// New creates a Filter. If duration is zero the filter is disabled.
func New(d time.Duration, now time.Time) *Filter {
	if d == 0 {
		return &Filter{enabled: false}
	}
	return &Filter{cutoff: now.Add(-d), enabled: true}
}

// Enabled reports whether the filter is active.
func (f *Filter) Enabled() bool { return f.enabled }

// Keep returns true if the line's timestamp is within the window.
// If the filter is disabled or ts is zero, the line is always kept.
func (f *Filter) Keep(ts time.Time) bool {
	if !f.enabled || ts.IsZero() {
		return true
	}
	return !ts.Before(f.cutoff)
}
