// Package output provides formatting and output destination support
// for logslice results.
package output

import (
	"fmt"
	"io"
	"strings"
)

// Format represents an output format type.
type Format string

const (
	// FormatPlain outputs lines as-is.
	FormatPlain Format = "plain"
	// FormatJSON wraps each line in a simple JSON object.
	FormatJSON Format = "json"
	// FormatCSV outputs lines with index and content as CSV.
	FormatCSV Format = "csv"
)

// Formatter writes log lines to a destination in a specific format.
type Formatter struct {
	format Format
	w      io.Writer
}

// NewFormatter creates a new Formatter for the given format and writer.
// Returns an error if the format is unrecognised.
func NewFormatter(format Format, w io.Writer) (*Formatter, error) {
	switch format {
	case FormatPlain, FormatJSON, FormatCSV:
		return &Formatter{format: format, w: w}, nil
	default:
		return nil, fmt.Errorf("unsupported format %q: must be plain, json, or csv", format)
	}
}

// WriteLine writes a single log line at position idx (1-based) using the
// configured format.
func (f *Formatter) WriteLine(idx int, line string) error {
	var out string
	switch f.format {
	case FormatJSON:
		escaped := strings.ReplaceAll(line, `"`, `\"`)
		out = fmt.Sprintf(`{"index":%d,"line":%q}`, idx, escaped)
	case FormatCSV:
		escaped := strings.ReplaceAll(line, `"`, `""`)
		out = fmt.Sprintf(`%d,"%s"`, idx, escaped)
	default:
		out = line
	}
	_, err := fmt.Fprintln(f.w, out)
	return err
}

// ParseFormat converts a string to a Format, returning an error if invalid.
func ParseFormat(s string) (Format, error) {
	switch Format(strings.ToLower(s)) {
	case FormatPlain:
		return FormatPlain, nil
	case FormatJSON:
		return FormatJSON, nil
	case FormatCSV:
		return FormatCSV, nil
	default:
		return "", fmt.Errorf("unknown format %q", s)
	}
}
