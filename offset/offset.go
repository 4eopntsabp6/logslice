// Package offset provides line-offset filtering: skip the first N lines
// or emit only lines starting from a given offset.
package offset

import "fmt"

// Mode controls how the offset filter operates.
type Mode int

const (
	ModeNone  Mode = iota // pass all lines through
	ModeSkip              // skip the first N lines
	ModeStart             // emit lines from line N onward (1-based)
)

// Filter holds offset state.
type Filter struct {
	mode    Mode
	n       int
	counter int
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "skip":
		return ModeSkip, nil
	case "start":
		return ModeStart, nil
	default:
		return ModeNone, fmt.Errorf("offset: unknown mode %q", s)
	}
}

// New creates a Filter. For ModeNone n is ignored.
func New(mode Mode, n int) (*Filter, error) {
	if mode != ModeNone && n <= 0 {
		return nil, fmt.Errorf("offset: n must be > 0, got %d", n)
	}
	return &Filter{mode: mode, n: n}, nil
}

// Keep returns true if the line should be kept.
func (f *Filter) Keep(_ string) bool {
	f.counter++
	switch f.mode {
	case ModeNone:
		return true
	case ModeSkip:
		return f.counter > f.n
	case ModeStart:
		return f.counter >= f.n
	default:
		return true
	}
}
