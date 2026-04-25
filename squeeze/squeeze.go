// Package squeeze provides a filter that collapses consecutive blank lines
// into a single blank line, reducing visual noise in log output.
package squeeze

import (
	"fmt"
	"strings"
)

// Mode controls how blank-line squeezing behaves.
type Mode int

const (
	ModeNone       Mode = iota // pass all lines unchanged
	ModeBlank                  // collapse consecutive blank lines into one
	ModeWhitespace             // treat whitespace-only lines as blank
)

// Filter holds state for the squeeze operation.
type Filter struct {
	mode         Mode
	lastWasBlank bool
}

// ParseMode converts a string to a Mode value.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "blank":
		return ModeBlank, nil
	case "whitespace":
		return ModeWhitespace, nil
	default:
		return ModeNone, fmt.Errorf("squeeze: unknown mode %q (want none|blank|whitespace)", s)
	}
}

// New creates a Filter for the given mode.
func New(mode Mode) (*Filter, error) {
	return &Filter{mode: mode}, nil
}

// Keep returns true if the line should be emitted.
// It must be called for every line in order so that consecutive-blank
// detection works correctly.
func (f *Filter) Keep(line string) bool {
	if f.mode == ModeNone {
		return true
	}

	isBlank := line == ""
	if f.mode == ModeWhitespace {
		isBlank = strings.TrimSpace(line) == ""
	}

	if isBlank {
		if f.lastWasBlank {
			return false
		}
		f.lastWasBlank = true
		return true
	}

	f.lastWasBlank = false
	return true
}
