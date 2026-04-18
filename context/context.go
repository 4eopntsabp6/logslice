// Package context provides line context extraction — capturing N lines
// before and/or after each matched line.
package context

// Mode controls how context lines are captured.
type Mode int

const (
	ModeNone Mode = iota
	ModeBefore
	ModeAfter
	ModeBoth
)

// Buffer holds a sliding window of recent lines and pending after-lines.
type Buffer struct {
	mode    Mode
	before  int
	after   int
	ring    []string
	head    int
	count   int
	pending int // lines of after-context remaining to emit
}

// New creates a Buffer. before/after are ignored when mode is ModeNone.
func New(mode Mode, before, after int) (*Buffer, error) {
	if mode != ModeNone && before < 0 {
		return nil, fmt.Errorf("context: before must be >= 0")
	}
	if mode != ModeNone && after < 0 {
		return nil, fmt.Errorf("context: after must be >= 0")
	}
	b := &Buffer{mode: mode, before: before, after: after}
	if mode == ModeBefore || mode == ModeBoth {
		b.ring = make([]string, before)
	}
	return b, nil
}

// Feed records a line into the before-ring. Returns true if the line should
// be emitted as after-context (i.e. we are in a post-match window).
func (b *Buffer) Feed(line string) bool {
	if b.mode == ModeNone {
		return false
	}
	if b.ring != nil && b.before > 0 {
		b.ring[b.head] = line
		b.head = (b.head + 1) % b.before
		if b.count < b.before {
			b.count++
		}
	}
	if b.pending > 0 {
		b.pending--
		return true
	}
	return false
}

// Before returns buffered pre-match lines in order.
func (b *Buffer) Before() []string {
	if b.before == 0 || b.count == 0 {
		return nil
	}
	out := make([]string, b.count)
	start := (b.head - b.count + b.before) % b.before
	for i := 0; i < b.count; i++ {
		out[i] = b.ring[(start+i)%b.before]
	}
	return out
}

// OnMatch signals a match, arming the after-context counter.
func (b *Buffer) OnMatch() {
	if b.mode == ModeAfter || b.mode == ModeBoth {
		b.pending = b.after
	}
	// reset before ring so same lines aren't re-emitted
	b.count = 0
}

import "fmt"
