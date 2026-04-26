// Package normalize provides line normalization by collapsing internal
// whitespace, trimming leading/trailing space, or both.
package normalize

import (
	"fmt"
	"regexp"
	"strings"
)

// Mode controls how normalization is applied.
type Mode int

const (
	ModeNone      Mode = iota // no normalization
	ModeTrim                  // trim leading and trailing whitespace
	ModeCollapse             // collapse internal runs of whitespace to a single space
	ModeAll                  // trim + collapse
)

var modeNames = map[string]Mode{
	"none":     ModeNone,
	"trim":     ModeTrim,
	"collapse": ModeCollapse,
	"all":      ModeAll,
}

var multiSpace = regexp.MustCompile(`\s+`)

// ParseMode converts a string into a Mode.
func ParseMode(s string) (Mode, error) {
	if m, ok := modeNames[strings.ToLower(s)]; ok {
		return m, nil
	}
	return ModeNone, fmt.Errorf("normalize: unknown mode %q (want none|trim|collapse|all)", s)
}

// Normalizer applies whitespace normalization to log lines.
type Normalizer struct {
	mode Mode
}

// New creates a Normalizer. Returns an error if mode is invalid.
func New(mode Mode) (*Normalizer, error) {
	if mode < ModeNone || mode > ModeAll {
		return nil, fmt.Errorf("normalize: invalid mode %d", mode)
	}
	return &Normalizer{mode: mode}, nil
}

// Apply returns the normalized form of line according to the configured mode.
func (n *Normalizer) Apply(line string) string {
	switch n.mode {
	case ModeTrim:
		return strings.TrimSpace(line)
	case ModeCollapse:
		return multiSpace.ReplaceAllString(line, " ")
	case ModeAll:
		return strings.TrimSpace(multiSpace.ReplaceAllString(line, " "))
	default:
		return line
	}
}

// Enabled reports whether normalization is active.
func (n *Normalizer) Enabled() bool {
	return n.mode != ModeNone
}
