package config

import (
	"fmt"
	"strings"
)

// HighlightConfig holds the user-supplied highlighting preferences.
type HighlightConfig struct {
	// Pattern is the regex used to select substrings to colour.
	Pattern string
	// Colour is the human-readable colour name (red, yellow, cyan).
	Colour string
	// Enabled controls whether highlighting is active at all.
	Enabled bool
}

// ValidColours lists the accepted colour names.
var ValidColours = []string{"red", "yellow", "cyan"}

// Validate checks that the highlight configuration is self-consistent.
func (h HighlightConfig) Validate() error {
	if !h.Enabled {
		return nil
	}
	if h.Pattern == "" {
		return fmt.Errorf("highlight: --highlight-pattern must be set when highlighting is enabled")
	}
	norm := strings.ToLower(strings.TrimSpace(h.Colour))
	for _, c := range ValidColours {
		if norm == c {
			return nil
		}
	}
	return fmt.Errorf("highlight: unsupported colour %q, choose from %s",
		h.Colour, strings.Join(ValidColours, ", "))
}
