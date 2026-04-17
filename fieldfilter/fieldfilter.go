// Package fieldfilter provides filtering of log lines based on extracted field values.
package fieldfilter

import (
	"fmt"
	"regexp"
	"strings"
)

// Mode defines how field value matching is performed.
type Mode string

const (
	ModeNone   Mode = "none"
	ModeExact  Mode = "exact"
	ModePrefix Mode = "prefix"
	ModeRegex  Mode = "regex"
)

// Filter matches log lines where a named field satisfies a condition.
type Filter struct {
	mode    Mode
	key     string
	value   string
	pattern *regexp.Regexp
}

// ParseMode parses a mode string into a Mode.
func ParseMode(s string) (Mode, error) {
	switch Mode(strings.ToLower(s)) {
	case ModeNone, ModeExact, ModePrefix, ModeRegex:
		return Mode(strings.ToLower(s)), nil
	}
	return "", fmt.Errorf("fieldfilter: unknown mode %q", s)
}

// New creates a new Filter.
func New(mode Mode, key, value string) (*Filter, error) {
	if mode == ModeNone {
		return &Filter{mode: ModeNone}, nil
	}
	if key == "" {
		return nil, fmt.Errorf("fieldfilter: key required for mode %q", mode)
	}
	if value == "" {
		return nil, fmt.Errorf("fieldfilter: value required for mode %q", mode)
	}
	f := &Filter{mode: mode, key: key, value: value}
	if mode == ModeRegex {
		re, err := regexp.Compile(value)
		if err != nil {
			return nil, fmt.Errorf("fieldfilter: invalid regex: %w", err)
		}
		f.pattern = re
	}
	return f, nil
}

// Match returns true if the line passes the field filter.
func (f *Filter) Match(line string) bool {
	if f.mode == ModeNone {
		return true
	}
	v := extractKV(line, f.key)
	switch f.mode {
	case ModeExact:
		return v == f.value
	case ModePrefix:
		return strings.HasPrefix(v, f.value)
	case ModeRegex:
		return f.pattern.MatchString(v)
	}
	return false
}

func extractKV(line, key string) string {
	prefix := key + "="
	for _, part := range strings.Fields(line) {
		if strings.HasPrefix(part, prefix) {
			return strings.TrimPrefix(part, prefix)
		}
	}
	return ""
}
