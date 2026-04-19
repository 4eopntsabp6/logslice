// Package prefix provides line prefixing functionality for logslice output.
package prefix

import (
	"fmt"
	"strings"
)

// Mode controls how lines are prefixed.
type Mode int

const (
	ModeNone      Mode = iota
	ModeText           // prepend a fixed string
	ModeLineNumber     // prepend the current line number
)

// Prefixer applies a prefix to each line.
type Prefixer struct {
	mode    Mode
	text    string
	counter int
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "text":
		return ModeText, nil
	case "linenum":
		return ModeLineNumber, nil
	default:
		return ModeNone, fmt.Errorf("unknown prefix mode %q: want none, text, linenum", s)
	}
}

// New creates a Prefixer. text is only used when mode is ModeText.
func New(mode Mode, text string) (*Prefixer, error) {
	if mode == ModeText && text == "" {
		return nil, fmt.Errorf("prefix mode 'text' requires a non-empty prefix string")
	}
	return &Prefixer{mode: mode, text: text}, nil
}

// Apply prepends the configured prefix to line and returns the result.
// If mode is ModeNone the original line is returned unchanged.
func (p *Prefixer) Apply(line string) string {
	switch p.mode {
	case ModeText:
		return p.text + line
	case ModeLineNumber:
		p.counter++
		return fmt.Sprintf("%d: %s", p.counter, line)
	default:
		return line
	}
}

// Reset resets the internal line counter (used for ModeLineNumber).
func (p *Prefixer) Reset() {
	p.counter = 0
}
