// Package dedupe provides line deduplication for log output,
// optionally suppressing consecutive or globally repeated lines.
package dedupe

// Mode controls how deduplication is applied.
type Mode int

const (
	// ModeNone disables deduplication.
	ModeNone Mode = iota
	// ModeConsecutive suppresses lines that repeat back-to-back.
	ModeConsecutive
	// ModeGlobal suppresses any line already seen anywhere in the stream.
	ModeGlobal
)

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, bool) {
	switch s {
	case "none", "":
		return ModeNone, true
	case "consecutive":
		return ModeConsecutive, true
	case "global":
		return ModeGlobal, true
	}
	return ModeNone, false
}

// Filter holds state for deduplication across a stream of lines.
type Filter struct {
	mode     Mode
	last     string
	seen     map[string]struct{}
	Dropped  int
}

// New creates a Filter for the given mode.
func New(mode Mode) *Filter {
	f := &Filter{mode: mode}
	if mode == ModeGlobal {
		f.seen = make(map[string]struct{})
	}
	return f
}

// Allow returns true if the line should be passed through.
// It updates internal state as a side effect.
func (f *Filter) Allow(line string) bool {
	switch f.mode {
	case ModeNone:
		return true
	case ModeConsecutive:
		if line == f.last {
			f.Dropped++
			return false
		}
		f.last = line
		return true
	case ModeGlobal:
		if _, exists := f.seen[line]; exists {
			f.Dropped++
			return false
		}
		f.seen[line] = struct{}{}
		return true
	}
	return true
}

// Reset clears all internal state, allowing the filter to be reused.
func (f *Filter) Reset() {
	f.last = ""
	f.Dropped = 0
	if f.mode == ModeGlobal {
		f.seen = make(map[string]struct{})
	}
}
