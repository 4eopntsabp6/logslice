// Package sample provides line sampling for large log streams,
// allowing every Nth line to be emitted rather than all matched lines.
package sample

import "fmt"

// Mode represents the sampling strategy.
type Mode int

const (
	// ModeNone disables sampling; all lines pass through.
	ModeNone Mode = iota
	// ModeNth emits every Nth line.
	ModeNth
)

// Sampler holds sampling state.
type Sampler struct {
	mode    Mode
	n       int
	counter int
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "nth":
		return ModeNth, nil
	default:
		return ModeNone, fmt.Errorf("unknown sample mode %q: want none|nth", s)
	}
}

// New creates a Sampler. n is the interval for ModeNth (must be >= 1).
// For ModeNone, n is ignored.
func New(mode Mode, n int) (*Sampler, error) {
	if mode == ModeNth && n < 1 {
		return nil, fmt.Errorf("sample interval must be >= 1, got %d", n)
	}
	return &Sampler{mode: mode, n: n}, nil
}

// Keep returns true if the current line should be kept.
// It must be called once per candidate line.
func (s *Sampler) Keep() bool {
	if s.mode == ModeNone {
		return true
	}
	s.counter++
	if s.counter >= s.n {
		s.counter = 0
		return true
	}
	return false
}

// Reset resets the internal counter.
func (s *Sampler) Reset() {
	s.counter = 0
}
