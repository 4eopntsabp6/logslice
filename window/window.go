// Package window provides a sliding window filter that keeps only lines
// whose timestamps fall within a rolling duration relative to the first
// matched timestamp.
package window

import (
	"fmt"
	"time"

	"github.com/robinovitch61/logslice/parser"
)

// Mode controls how the sliding window filter behaves.
type Mode int

const (
	// ModeNone disables the sliding window filter.
	ModeNone Mode = iota
	// ModeSliding keeps lines within a rolling duration from the first timestamp.
	ModeSliding
)

// Filter holds state for the sliding window filter.
type Filter struct {
	mode     Mode
	duration time.Duration
	anchor   *time.Time
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "sliding":
		return ModeSliding, nil
	default:
		return ModeNone, fmt.Errorf("window: unknown mode %q", s)
	}
}

// New creates a Filter. For ModeNone, duration is ignored.
func New(mode Mode, duration time.Duration) (*Filter, error) {
	if mode == ModeSliding && duration <= 0 {
		return nil, fmt.Errorf("window: sliding mode requires a positive duration")
	}
	return &Filter{mode: mode, duration: duration}, nil
}

// Keep returns true if the line should be kept.
// It extracts a timestamp from line; if none is found the line is kept
// when the anchor has not yet been set, and dropped once the window is
// established.
func (f *Filter) Keep(line string) bool {
	if f.mode == ModeNone {
		return true
	}

	t, ok := parser.ExtractTimestamp(line)
	if !ok {
		return f.anchor == nil
	}

	if f.anchor == nil {
		f.anchor = &t
		return true
	}

	return !t.Before(*f.anchor) && t.Before(f.anchor.Add(f.duration))
}

// Reset clears the anchor so the window restarts on the next timestamped line.
func (f *Filter) Reset() {
	f.anchor = nil
}
