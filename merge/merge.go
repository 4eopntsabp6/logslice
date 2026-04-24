// Package merge provides line-merging for continuation lines in log output.
// A continuation line is one that matches a given prefix pattern (e.g. a stack
// trace line that begins with whitespace or a known marker).
package merge

import (
	"fmt"
	"regexp"
	"strings"
)

// Mode controls how the merger behaves.
type Mode int

const (
	ModeNone        Mode = iota // pass-through; no merging
	ModeContinuation            // merge lines matching a continuation pattern
)

// Merger accumulates continuation lines and emits merged blocks.
type Merger struct {
	mode      Mode
	pattern   *regexp.Regexp
	separator string
	pending   []string
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "none", "":
		return ModeNone, nil
	case "continuation":
		return ModeContinuation, nil
	default:
		return ModeNone, fmt.Errorf("merge: unknown mode %q", s)
	}
}

// New creates a Merger. pattern is the regex that identifies a continuation
// line. separator is the string used to join merged lines (e.g. " ").
func New(mode Mode, pattern, separator string) (*Merger, error) {
	if mode == ModeNone {
		return &Merger{mode: ModeNone}, nil
	}
	if pattern == "" {
		return nil, fmt.Errorf("merge: pattern required for mode %q", mode)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("merge: invalid pattern: %w", err)
	}
	if separator == "" {
		separator = " "
	}
	return &Merger{mode: mode, pattern: re, separator: separator}, nil
}

// Feed accepts a line and returns a merged line if one is ready, or ("", false)
// when the line has been buffered as a continuation.
func (m *Merger) Feed(line string) (string, bool) {
	if m.mode == ModeNone {
		return line, true
	}
	if m.pattern.MatchString(line) {
		m.pending = append(m.pending, line)
		return "", false
	}
	// New root line — flush any pending block first.
	if len(m.pending) > 0 {
		merged := strings.Join(m.pending, m.separator)
		m.pending = []string{line}
		return merged, true
	}
	m.pending = []string{line}
	return "", false
}

// Flush returns any remaining buffered lines as a single merged string.
// Call after the input is exhausted.
func (m *Merger) Flush() (string, bool) {
	if len(m.pending) == 0 {
		return "", false
	}
	merged := strings.Join(m.pending, m.separator)
	m.pending = nil
	return merged, true
}
