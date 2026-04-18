// Package linenum provides line number tracking and filtering for log slices.
package linenum

import "fmt"

// Mode controls how line number filtering is applied.
type Mode int

const (
	ModeNone  Mode = iota
	ModeRange      // include only lines within [From, To]
	ModeList       // include only lines matching explicit list
)

// Filter holds line number filter configuration.
type Filter struct {
	mode Mode
	from int
	to   int
	set  map[int]struct{}
}

// ParseMode parses a mode string into a Mode value.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "range":
		return ModeRange, nil
	case "list":
		return ModeList, nil
	}
	return ModeNone, fmt.Errorf("linenum: unknown mode %q", s)
}

// New creates a new Filter. For ModeRange provide from/to; for ModeList provide lines.
func New(mode Mode, from, to int, lines []int) (*Filter, error) {
	if mode == ModeRange {
		if from < 1 {
			return nil, fmt.Errorf("linenum: from must be >= 1")
		}
		if to > 0 && to < from {
			return nil, fmt.Errorf("linenum: to must be >= from")
		}
	}
	set := make(map[int]struct{}, len(lines))
	for _, l := range lines {
		set[l] = struct{}{}
	}
	return &Filter{mode: mode, from: from, to: to, set: set}, nil
}

// Keep returns true if the given 1-based line number should be kept.
func (f *Filter) Keep(n int) bool {
	switch f.mode {
	case ModeRange:
		if n < f.from {
			return false
		}
		if f.to > 0 && n > f.to {
			return false
		}
		return true
	case ModeList:
		_, ok := f.set[n]
		return ok
	}
	return true
}
