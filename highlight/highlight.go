// Package highlight provides terminal colour highlighting for matched
// substrings within log lines.
package highlight

import (
	"regexp"
	"strings"
)

// Colour escape codes for terminal output.
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

// Highlighter wraps a compiled regex and a chosen colour code.
type Highlighter struct {
	pattern *regexp.Regexp
	colour  string
	enabled bool
}

// New returns a Highlighter for the given pattern and colour.
// If pattern is empty the highlighter is a no-op.
func New(pattern, colour string, enabled bool) (*Highlighter, error) {
	if pattern == "" {
		return &Highlighter{enabled: false}, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &Highlighter{pattern: re, colour: colour, enabled: enabled}, nil
}

// Apply returns the line with all pattern matches wrapped in the chosen
// colour escape sequence. If highlighting is disabled the original line
// is returned unchanged.
func (h *Highlighter) Apply(line string) string {
	if !h.enabled || h.pattern == nil {
		return line
	}
	return h.pattern.ReplaceAllStringFunc(line, func(match string) string {
		return h.colour + Bold + match + Reset
	})
}

// Strip removes all ANSI escape sequences from a string.
func Strip(s string) string {
	ansi := regexp.MustCompile(`\033\[[0-9;]*m`)
	return ansi.ReplaceAllString(s, "")
}

// ParseColour maps a human-readable colour name to its escape code.
// Unknown names fall back to Cyan.
func ParseColour(name string) string {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "red":
		return Red
	case "yellow":
		return Yellow
	case "cyan":
		return Cyan
	default:
		return Cyan
	}
}
