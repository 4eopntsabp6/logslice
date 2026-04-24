// Package entropy provides line filtering based on Shannon entropy scoring.
// Lines with entropy outside the configured range are dropped or flagged.
package entropy

import (
	"fmt"
	"math"
)

// Mode controls how entropy filtering is applied.
type Mode int

const (
	ModeNone  Mode = iota
	ModeAbove      // keep lines whose entropy is above the threshold
	ModeBelow      // keep lines whose entropy is below the threshold
)

// Filter scores each line and decides whether to keep it.
type Filter struct {
	mode      Mode
	threshold float64
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "above":
		return ModeAbove, nil
	case "below":
		return ModeBelow, nil
	default:
		return ModeNone, fmt.Errorf("entropy: unknown mode %q (want none|above|below)", s)
	}
}

// New creates a Filter. threshold must be >= 0 when mode is not ModeNone.
func New(mode Mode, threshold float64) (*Filter, error) {
	if mode != ModeNone && threshold < 0 {
		return nil, fmt.Errorf("entropy: threshold must be >= 0, got %f", threshold)
	}
	return &Filter{mode: mode, threshold: threshold}, nil
}

// Keep returns true if the line passes the entropy filter.
func (f *Filter) Keep(line string) bool {
	if f.mode == ModeNone {
		return true
	}
	s := score(line)
	switch f.mode {
	case ModeAbove:
		return s >= f.threshold
	case ModeBelow:
		return s <= f.threshold
	}
	return true
}

// Score returns the Shannon entropy of s in bits.
func Score(s string) float64 { return score(s) }

func score(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	n := float64(len([]rune(s)))
	var h float64
	for _, c := range freq {
		p := float64(c) / n
		h -= p * math.Log2(p)
	}
	return h
}
