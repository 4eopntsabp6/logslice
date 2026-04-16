package transform

import (
	"fmt"
	"strings"
)

// Mode defines the transformation to apply to each log line.
type Mode int

const (
	ModeNone    Mode = iota
	ModeUpper
	ModeLower
	ModeTrimSpace
)

// Transformer applies a text transformation to log lines.
type Transformer struct {
	mode Mode
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "upper":
		return ModeUpper, nil
	case "lower":
		return ModeLower, nil
	case "trim":
		return ModeTrimSpace, nil
	default:
		return ModeNone, fmt.Errorf("unknown transform mode: %q", s)
	}
}

// New creates a new Transformer for the given mode.
func New(mode Mode) *Transformer {
	return &Transformer{mode: mode}
}

// Apply transforms the given line according to the configured mode.
func (t *Transformer) Apply(line string) string {
	switch t.mode {
	case ModeUpper:
		return strings.ToUpper(line)
	case ModeLower:
		return strings.ToLower(line)
	case ModeTrimSpace:
		return strings.TrimSpace(line)
	default:
		return line
	}
}

// Enabled reports whether any transformation is active.
func (t *Transformer) Enabled() bool {
	return t.mode != ModeNone
}
