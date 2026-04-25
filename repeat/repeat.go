// Package repeat provides a filter that drops lines matching a pattern
// more than N times, useful for suppressing noisy repeated log entries.
package repeat

import (
	"fmt"
	"regexp"
)

// Mode controls how the repeat filter behaves.
type Mode int

const (
	ModeNone  Mode = iota // pass all lines through
	ModeLimit             // drop lines matching pattern after N occurrences
)

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "limit":
		return ModeLimit, nil
	default:
		return ModeNone, fmt.Errorf("repeat: unknown mode %q", s)
	}
}

// Filter suppresses lines that match a pattern after they have been seen
// more than Max times.
type Filter struct {
	mode    Mode
	re      *regexp.Regexp
	max     int
	counts  map[string]int
}

// New creates a new Filter. When mode is ModeNone, pattern and max are ignored.
func New(mode Mode, pattern string, max int) (*Filter, error) {
	if mode == ModeNone {
		return &Filter{mode: ModeNone}, nil
	}
	if pattern == "" {
		return nil, fmt.Errorf("repeat: pattern required for mode %q", "limit")
	}
	if max <= 0 {
		return nil, fmt.Errorf("repeat: max must be > 0, got %d", max)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("repeat: invalid pattern: %w", err)
	}
	return &Filter{
		mode:   ModeLimit,
		re:     re,
		max:    max,
		counts: make(map[string]int),
	}, nil
}

// Keep returns true if the line should be kept.
func (f *Filter) Keep(line string) bool {
	if f.mode == ModeNone {
		return true
	}
	match := f.re.FindString(line)
	if match == "" {
		return true
	}
	f.counts[match]++
	return f.counts[match] <= f.max
}
