package slicer

import (
	"fmt"
	"io"
)

// WriteFormat controls how results are written to the output.
type WriteFormat int

const (
	// FormatPlain writes one log line per output line.
	FormatPlain WriteFormat = iota
	// FormatNumbered prefixes each line with its original line number.
	FormatNumbered
)

// WriteOptions configures the Write function.
type WriteOptions struct {
	Format WriteFormat
}

// Write outputs results to w according to wopts.
// It returns the number of lines written and any error encountered.
func Write(w io.Writer, results []Result, wopts WriteOptions) (int, error) {
	written := 0
	for _, r := range results {
		var err error
		switch wopts.Format {
		case FormatNumbered:
			_, err = fmt.Fprintf(w, "%d\t%s\n", r.LineNum, r.Line)
		default:
			_, err = fmt.Fprintln(w, r.Line)
		}
		if err != nil {
			return written, fmt.Errorf("write error at line %d: %w", r.LineNum, err)
		}
		written++
	}
	return written, nil
}

// Summary returns a human-readable summary string for a slice operation.
func Summary(total int, results []Result) string {
	if len(results) == 0 {
		return fmt.Sprintf("No matching lines found (scanned %d lines)", total)
	}
	return fmt.Sprintf(
		"Matched %d of %d lines (lines %d–%d)",
		len(results),
		total,
		results[0].LineNum,
		results[len(results)-1].LineNum,
	)
}
