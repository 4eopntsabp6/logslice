// Package input provides utilities for reading log input from files or stdin.
package input

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// LineReader wraps a bufio.Scanner and tracks line numbers.
type LineReader struct {
	scanner    *bufio.Scanner
	lineNum    int
	closer     io.Closer
}

// NewFileReader opens a file and returns a LineReader for it.
// The caller is responsible for calling Close when done.
func NewFileReader(path string) (*LineReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("input: cannot open file %q: %w", path, err)
	}
	return &LineReader{
		scanner: bufio.NewScanner(f),
		closer:  f,
	}, nil
}

// NewStdinReader returns a LineReader that reads from os.Stdin.
func NewStdinReader() *LineReader {
	return &LineReader{
		scanner: bufio.NewScanner(os.Stdin),
		closer:  io.NopCloser(os.Stdin),
	}
}

// ReadLine advances to the next line and returns it.
// Returns ("", false) when there are no more lines.
func (r *LineReader) ReadLine() (string, bool) {
	if !r.scanner.Scan() {
		return "", false
	}
	r.lineNum++
	return r.scanner.Text(), true
}

// LineNum returns the current 1-based line number.
func (r *LineReader) LineNum() int {
	return r.lineNum
}

// Err returns any scanner error that occurred during reading.
func (r *LineReader) Err() error {
	return r.scanner.Err()
}

// Close releases any underlying resources.
func (r *LineReader) Close() error {
	if r.closer != nil {
		return r.closer.Close()
	}
	return nil
}
