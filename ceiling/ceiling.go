// Package ceiling provides a filter that drops lines once a maximum
// match count has been reached. Unlike count (which filters by position),
// ceiling tracks how many lines have matched a pattern and stops emitting
// once the cap is hit.
package ceiling

import (
	"fmt"
	"regexp"
)

// Mode controls how the ceiling filter behaves.
type Mode int

const (
	ModeNone  Mode = iota // pass all lines through
	ModeCap               // stop after N matched lines
)

// ParseMode converts a string into a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "cap":
		return ModeCap, nil
	default:
		return ModeNone, fmt.Errorf("ceiling: unknown mode %q", s)
	}
}

// Filter holds the ceiling state.
type Filter struct {
	mode    Mode
	pattern *regexp.Regexp
	max     int
	matched int
}

// New creates a Filter. For ModeCap, pattern may be nil (matches every line)
// and max must be >= 1.
func New(mode Mode, pattern *regexp.Regexp, max int) (*Filter, error) {
	if mode == ModeNone {
		return &Filter{mode: ModeNone}, nil
	}
	if max < 1 {
		return nil, fmt.Errorf("ceiling: max must be >= 1, got %d", max)
	}
	return &Filter{mode: mode, pattern: pattern, max: max}, nil
}

// Keep returns true if the line should be emitted. Once the cap is reached
// all subsequent lines are dropped.
func (f *Filter) Keep(line string) bool {
	if f.mode == ModeNone {
		return true
	}
	if f.matched >= f.max {
		return false
	}
	matches := f.pattern == nil || f.pattern.MatchString(line)
	if matches {
		f.matched++
	}
	return matches
}
