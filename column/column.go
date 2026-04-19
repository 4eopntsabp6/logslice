// Package column provides line splitting and column extraction by delimiter.
package column

import (
	"fmt"
	"strings"
)

// Mode controls column extraction behaviour.
type Mode int

const (
	ModeNone Mode = iota
	ModeExtract
)

// Extractor extracts a specific column from each line.
type Extractor struct {
	mode      Mode
	delimiter string
	index     int
}

// ParseMode parses a mode string into a Mode value.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(s) {
	case "", "none":
		return ModeNone, nil
	case "extract":
		return ModeExtract, nil
	}
	return ModeNone, fmt.Errorf("column: unknown mode %q", s)
}

// New creates a new Extractor. delimiter and index are ignored when mode is ModeNone.
func New(mode Mode, delimiter string, index int) (*Extractor, error) {
	if mode == ModeNone {
		return &Extractor{mode: ModeNone}, nil
	}
	if delimiter == "" {
		return nil, fmt.Errorf("column: delimiter must not be empty")
	}
	if index < 0 {
		return nil, fmt.Errorf("column: index must be >= 0, got %d", index)
	}
	return &Extractor{mode: mode, delimiter: delimiter, index: index}, nil
}

// Apply returns the extracted column, or the original line if mode is None
// or the column index is out of range.
func (e *Extractor) Apply(line string) string {
	if e.mode == ModeNone {
		return line
	}
	parts := strings.Split(line, e.delimiter)
	if e.index >= len(parts) {
		return line
	}
	return strings.TrimSpace(parts[e.index])
}
