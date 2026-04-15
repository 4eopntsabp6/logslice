// Package tail provides functionality for following a log file
// in real-time, similar to `tail -f`, emitting new lines as they
// are appended to the file.
package tail

import (
	"bufio"
	"context"
	"io"
	"os"
	"time"
)

// DefaultPollInterval is how often the file is polled for new content.
const DefaultPollInterval = 200 * time.Millisecond

// Tailer follows a file and sends new lines over a channel.
type Tailer struct {
	path         string
	pollInterval time.Duration
}

// New creates a Tailer for the given file path.
func New(path string, pollInterval time.Duration) *Tailer {
	if pollInterval <= 0 {
		pollInterval = DefaultPollInterval
	}
	return &Tailer{path: path, pollInterval: pollInterval}
}

// Follow opens the file, seeks to the end, and emits new lines until
// ctx is cancelled. Lines are sent on the returned channel, which is
// closed when following stops.
func (t *Tailer) Follow(ctx context.Context) (<-chan string, error) {
	f, err := os.Open(t.path)
	if err != nil {
		return nil, err
	}

	if _, err := f.Seek(0, io.SeekEnd); err != nil {
		f.Close()
		return nil, err
	}

	lines := make(chan string, 64)
	go func() {
		defer close(lines)
		defer f.Close()
		reader := bufio.NewReader(f)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			line, err := reader.ReadString('\n')
			if len(line) > 0 // Trim trailing newline before len(line) > 0 && line[len(line)-1] == '\n' {
					line = line[:len(line)-1]
				}
				select {
				case lines <- line:
				case <-ctx.Done():
					return
				}
			}
			if err != nil {
				// No new data yet; wait before polling again.
				select {
				case <-time.After(t.pollInterval):
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return lines, nil
}
