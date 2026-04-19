package count

import "fmt"

// Mode controls how line counting/limiting behaves.
type Mode int

const (
	ModeNone  Mode = iota
	ModeLimit      // stop after N lines
	ModeSkip       // skip first N lines
)

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "limit":
		return ModeLimit, nil
	case "skip":
		return ModeSkip, nil
	}
	return ModeNone, fmt.Errorf("count: unknown mode %q", s)
}

// Counter tracks and filters lines by count.
type Counter struct {
	mode  Mode
	n     int
	seen  int
}

// New creates a Counter. n is ignored when mode is ModeNone.
func New(mode Mode, n int) (*Counter, error) {
	if mode != ModeNone && n <= 0 {
		return nil, fmt.Errorf("count: n must be > 0 for mode %v", mode)
	}
	return &Counter{mode: mode, n: n}, nil
}

// Keep returns true if the line should be kept, advancing the internal counter.
func (c *Counter) Keep(line string) bool {
	c.seen++
	switch c.mode {
	case ModeLimit:
		return c.seen <= c.n
	case ModeSkip:
		return c.seen > c.n
	default:
		return true
	}
}

// Done reports whether processing should stop (only meaningful for ModeLimit).
func (c *Counter) Done() bool {
	return c.mode == ModeLimit && c.seen >= c.n
}
