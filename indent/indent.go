package indent

import (
	"fmt"
	"strings"
)

// Mode controls how indentation is applied.
type Mode int

const (
	ModeNone  Mode = iota
	ModeSpaces
	ModeTabs
)

// Indenter applies a fixed indentation prefix to log lines.
type Indenter struct {
	mode   Mode
	prefix string
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "spaces":
		return ModeSpaces, nil
	case "tabs":
		return ModeTabs, nil
	default:
		return ModeNone, fmt.Errorf("indent: unknown mode %q", s)
	}
}

// New creates an Indenter. n is the number of spaces or tabs to prepend.
func New(mode Mode, n int) (*Indenter, error) {
	if mode == ModeNone {
		return &Indenter{mode: ModeNone}, nil
	}
	if n <= 0 {
		return nil, fmt.Errorf("indent: n must be greater than zero, got %d", n)
	}
	var unit string
	if mode == ModeTabs {
		unit = "\t"
	} else {
		unit = " "
	}
	return &Indenter{
		mode:   mode,
		prefix: strings.Repeat(unit, n),
	}, nil
}

// Apply prepends the configured prefix to line, or returns it unchanged
// when mode is None.
func (i *Indenter) Apply(line string) string {
	if i.mode == ModeNone {
		return line
	}
	return i.prefix + line
}

// Enabled reports whether indentation is active.
func (i *Indenter) Enabled() bool {
	return i.mode != ModeNone
}
