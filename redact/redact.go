// Package redact provides line-level redaction by replacing sensitive
// pattern matches with a fixed placeholder string.
package redact

import (
	"fmt"
	"regexp"
)

// Mode controls whether redaction is active.
type Mode int

const (
	ModeNone   Mode = iota
	ModePattern     // replace all matches of a regex with a placeholder
)

const defaultPlaceholder = "[REDACTED]"

// Redactor replaces sensitive content in log lines.
type Redactor struct {
	mode        Mode
	pattern     *regexp.Regexp
	placeholder string
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "pattern":
		return ModePattern, nil
	default:
		return ModeNone, fmt.Errorf("redact: unknown mode %q", s)
	}
}

// New creates a Redactor. When mode is ModeNone, pattern and placeholder are
// ignored and all lines are passed through unchanged.
func New(mode Mode, pattern, placeholder string) (*Redactor, error) {
	if mode == ModeNone {
		return &Redactor{mode: ModeNone}, nil
	}
	if pattern == "" {
		return nil, fmt.Errorf("redact: pattern required for mode %q", "pattern")
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("redact: invalid pattern: %w", err)
	}
	if placeholder == "" {
		placeholder = defaultPlaceholder
	}
	return &Redactor{mode: mode, pattern: re, placeholder: placeholder}, nil
}

// Apply returns the line with sensitive content replaced. If the Redactor is
// disabled (ModeNone) the original line is returned unchanged.
func (r *Redactor) Apply(line string) string {
	if r.mode == ModeNone {
		return line
	}
	return r.pattern.ReplaceAllString(line, r.placeholder)
}

// Enabled reports whether the Redactor will modify lines.
func (r *Redactor) Enabled() bool {
	return r.mode != ModeNone
}
