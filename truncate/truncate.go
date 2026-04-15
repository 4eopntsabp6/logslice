// Package truncate provides utilities for truncating long log lines
// to a configurable maximum byte or rune length before output.
package truncate

import (
	"fmt"
	"unicode/utf8"
)

// Mode controls how truncation is applied.
type Mode int

const (
	// ModeNone disables truncation.
	ModeNone Mode = iota
	// ModeBytes truncates at a maximum number of bytes.
	ModeBytes
	// ModeRunes truncates at a maximum number of Unicode code points.
	ModeRunes
)

// DefaultSuffix is appended to truncated lines.
const DefaultSuffix = "..."

// Truncator applies line truncation according to the configured mode and limit.
type Truncator struct {
	mode   Mode
	limit  int
	suffix string
}

// New creates a Truncator. limit must be > 0 when mode is not ModeNone.
// suffix is appended to truncated lines; pass "" to use DefaultSuffix.
func New(mode Mode, limit int, suffix string) (*Truncator, error) {
	if mode != ModeNone && limit <= 0 {
		return nil, fmt.Errorf("truncate: limit must be > 0, got %d", limit)
	}
	if suffix == "" {
		suffix = DefaultSuffix
	}
	return &Truncator{mode: mode, limit: limit, suffix: suffix}, nil
}

// Apply returns the (possibly truncated) version of line.
func (t *Truncator) Apply(line string) string {
	switch t.mode {
	case ModeBytes:
		return t.applyBytes(line)
	case ModeRunes:
		return t.applyRunes(line)
	default:
		return line
	}
}

func (t *Truncator) applyBytes(line string) string {
	if len(line) <= t.limit {
		return line
	}
	cut := t.limit
	for cut > 0 && !utf8.RuneStart(line[cut]) {
		cut--
	}
	return line[:cut] + t.suffix
}

func (t *Truncator) applyRunes(line string) string {
	count := 0
	for i := range line {
		if count == t.limit {
			return line[:i] + t.suffix
		}
		count++
	}
	return line
}

// ParseMode converts a string to a Mode. Recognised values: "none", "bytes", "runes".
func ParseMode(s string) (Mode, error) {
	switch s {
	case "none", "":
		return ModeNone, nil
	case "bytes":
		return ModeBytes, nil
	case "runes":
		return ModeRunes, nil
	default:
		return ModeNone, fmt.Errorf("truncate: unknown mode %q", s)
	}
}
