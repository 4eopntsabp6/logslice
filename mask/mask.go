// Package mask provides field masking for sensitive log data.
package mask

import (
	"regexp"
	"strings"
)

// Mode controls how masking is applied.
type Mode int

const (
	ModeNone  Mode = iota
	ModeRedact      // replace match with [REDACTED]
	ModePartial     // keep first/last 2 chars, mask middle
)

// Masker applies masking to log lines.
type Masker struct {
	mode    Mode
	pattern *regexp.Regexp
}

// ParseMode parses a mode string into a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "redact":
		return ModeRedact, nil
	case "partial":
		return ModePartial, nil
	default:
		return ModeNone, fmt.Errorf("mask: unknown mode %q", s)
	}
}

// New creates a Masker. pattern is required when mode != ModeNone.
func New(mode Mode, pattern string) (*Masker, error) {
	if mode == ModeNone {
		return &Masker{mode: mode}, nil
	}
	if pattern == "" {
		return nil, fmt.Errorf("mask: pattern required for mode %q", mode)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("mask: invalid pattern: %w", err)
	}
	return &Masker{mode: mode, pattern: re}, nil
}

// Apply masks sensitive content in line according to the configured mode.
func (m *Masker) Apply(line string) string {
	if m.mode == ModeNone || m.pattern == nil {
		return line
	}
	return m.pattern.ReplaceAllStringFunc(line, func(match string) string {
		switch m.mode {
		case ModeRedact:
			return "[REDACTED]"
		case ModePartial:
			if len(match) <= 4 {
				return strings.Repeat("*", len(match))
			}
			return match[:2] + strings.Repeat("*", len(match)-4) + match[len(match)-2:]
		default:
			return match
		}
	})
}
