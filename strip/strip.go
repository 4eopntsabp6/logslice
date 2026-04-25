// Package strip provides a filter that removes ANSI escape sequences
// and/or leading/trailing whitespace from log lines.
package strip

import (
	"fmt"
	"regexp"
	"strings"
)

// Mode controls what stripping behaviour is applied.
type Mode int

const (
	ModeNone      Mode = iota
	ModeANSI           // remove ANSI colour/control codes
	ModeWhitespace     // trim leading and trailing whitespace
	ModeBoth           // remove ANSI codes then trim whitespace
)

var ansiRe = regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

// ParseMode converts a string into a Mode value.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "ansi":
		return ModeANSI, nil
	case "whitespace":
		return ModeWhitespace, nil
	case "both":
		return ModeBoth, nil
	}
	return ModeNone, fmt.Errorf("strip: unknown mode %q", s)
}

// Stripper applies the configured stripping to each line.
type Stripper struct {
	mode Mode
}

// New creates a Stripper. ModeNone is valid and results in a no-op.
func New(mode Mode) (*Stripper, error) {
	return &Stripper{mode: mode}, nil
}

// Apply returns the line after applying the configured stripping.
func (s *Stripper) Apply(line string) string {
	switch s.mode {
	case ModeANSI:
		return ansiRe.ReplaceAllString(line, "")
	case ModeWhitespace:
		return strings.TrimSpace(line)
	case ModeBoth:
		return strings.TrimSpace(ansiRe.ReplaceAllString(line, ""))
	}
	return line
}

// Enabled reports whether any stripping will be performed.
func (s *Stripper) Enabled() bool {
	return s.mode != ModeNone
}
