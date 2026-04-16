package label

import (
	"fmt"
	"regexp"
	"strings"
)

// Mode controls how labels are applied to matched lines.
type Mode string

const (
	ModeNone   Mode = "none"
	ModePrefix Mode = "prefix"
	ModeAppend Mode = "append"
)

// Labeler tags lines matching a pattern with a label string.
type Labeler struct {
	mode    Mode
	label   string
	pattern *regexp.Regexp
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "none", "":
		return ModeNone, nil
	case "prefix":
		return ModePrefix, nil
	case "append":
		return ModeAppend, nil
	}
	return "", fmt.Errorf("unknown label mode: %q", s)
}

// New creates a Labeler. Pattern may be empty when mode is none.
func New(mode Mode, label, pattern string) (*Labeler, error) {
	if mode == ModeNone {
		return &Labeler{mode: mode}, nil
	}
	if label == "" {
		return nil, fmt.Errorf("label text must not be empty")
	}
	if pattern == "" {
		return nil, fmt.Errorf("pattern must not be empty when mode is %q", mode)
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("invalid pattern: %w", err)
	}
	return &Labeler{mode: mode, label: label, pattern: re}, nil
}

// Apply returns the line with a label attached if the pattern matches.
func (l *Labeler) Apply(line string) string {
	if l.mode == ModeNone {
		return line
	}
	if !l.pattern.MatchString(line) {
		return line
	}
	switch l.mode {
	case ModePrefix:
		return fmt.Sprintf("[%s] %s", l.label, line)
	case ModeAppend:
		return fmt.Sprintf("%s [%s]", line, l.label)
	}
	return line
}
