package output

import (
	"fmt"
	"io"
	"os"
)

// Destination describes where output should be written.
type Destination string

const (
	// DestStdout writes to standard output.
	DestStdout Destination = "stdout"
	// DestFile writes to a named file.
	DestFile Destination = "file"
)

// OpenDestination returns an io.WriteCloser for the given destination.
// When dest is DestStdout the path argument is ignored.
// The caller is responsible for closing the returned writer.
func OpenDestination(dest Destination, path string) (io.WriteCloser, error) {
	switch dest {
	case DestStdout:
		return nopCloser{os.Stdout}, nil
	case DestFile:
		if path == "" {
			return nil, fmt.Errorf("output path must not be empty for file destination")
		}
		f, err := os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("opening output file %q: %w", path, err)
		}
		return f, nil
	default:
		return nil, fmt.Errorf("unknown destination %q", dest)
	}
}

// nopCloser wraps an io.Writer that needs no closing (e.g. os.Stdout).
type nopCloser struct{ io.Writer }

func (nopCloser) Close() error { return nil }

// ParseDestination converts a string to a Destination.
func ParseDestination(s string) (Destination, error) {
	switch Destination(s) {
	case DestStdout:
		return DestStdout, nil
	case DestFile:
		return DestFile, nil
	default:
		return "", fmt.Errorf("unknown destination %q: must be stdout or file", s)
	}
}
