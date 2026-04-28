package reverse

import "fmt"

// Mode controls how line reversal is applied.
type Mode int

const (
	ModeNone    Mode = iota
	ModeReverse      // reverse the order of all collected lines
)

// Reverser buffers lines and emits them in reverse order.
type Reverser struct {
	mode Mode
	buf  []string
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "reverse":
		return ModeReverse, nil
	default:
		return ModeNone, fmt.Errorf("reverse: unknown mode %q (want none|reverse)", s)
	}
}

// New creates a Reverser for the given mode.
func New(mode Mode) *Reverser {
	return &Reverser{mode: mode}
}

// Feed appends a line to the internal buffer.
// When mode is ModeNone the line is returned immediately (pass-through).
// When mode is ModeReverse nil is returned until Flush is called.
func (r *Reverser) Feed(line string) []string {
	if r.mode == ModeNone {
		return []string{line}
	}
	r.buf = append(r.buf, line)
	return nil
}

// Flush returns all buffered lines in reverse order and resets the buffer.
// When mode is ModeNone Flush is a no-op.
func (r *Reverser) Flush() []string {
	if r.mode == ModeNone || len(r.buf) == 0 {
		return nil
	}
	out := make([]string, len(r.buf))
	for i, line := range r.buf {
		out[len(r.buf)-1-i] = line
	}
	r.buf = r.buf[:0]
	return out
}

// Enabled reports whether the reverser will actually reorder lines.
func (r *Reverser) Enabled() bool {
	return r.mode != ModeNone
}
