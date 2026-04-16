// Package field provides log field extraction by key from structured log lines.
package field

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ParseMode returns a validated mode string.
func ParseMode(s string) (string, error) {
	switch s {
	case "json", "kv", "none":
		return s, nil
	}
	return "", fmt.Errorf("unknown field mode %q: want json, kv, or none", s)
}

// Extractor pulls a named field value from a log line.
type Extractor struct {
	mode string
	key  string
}

// New creates an Extractor for the given mode and key.
// mode must be one of: json, kv, none.
func New(mode, key string) (*Extractor, error) {
	m, err := ParseMode(mode)
	if err != nil {
		return nil, err
	}
	if m != "none" && key == "" {
		return nil, fmt.Errorf("field key must not be empty when mode is %q", m)
	}
	return &Extractor{mode: m, key: key}, nil
}

// Extract returns the value of the configured key from line.
// Returns empty string if not found or mode is none.
func (e *Extractor) Extract(line string) string {
	switch e.mode {
	case "none":
		return ""
	case "json":
		return extractJSON(line, e.key)
	case "kv":
		return extractKV(line, e.key)
	}
	return ""
}

func extractJSON(line, key string) string {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		return ""
	}
	if v, ok := m[key]; ok {
		return fmt.Sprintf("%v", v)
	}
	return ""
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
