// Package bookmark provides line bookmarking by index or regex match.
package bookmark

import (
	"fmt"
	"regexp"
)

// Mode controls how bookmarks are applied.
type Mode int

const (
	ModeNone  Mode = iota
	ModeIndex      // match by line number (1-based)
	ModeRegex      // match by regex
)

// Bookmark marks matching lines with a prefix tag.
type Bookmark struct {
	mode    Mode
	indices map[int]bool
	re      *regexp.Regexp
	tag     string
}

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch s {
	case "", "none":
		return ModeNone, nil
	case "index":
		return ModeIndex, nil
	case "regex":
		return ModeRegex, nil
	}
	return ModeNone, fmt.Errorf("bookmark: unknown mode %q", s)
}

// New creates a Bookmark. tag is the prefix applied to matched lines.
func New(mode Mode, tag string, indices []int, pattern string) (*Bookmark, error) {
	if mode == ModeNone {
		return &Bookmark{mode: ModeNone}, nil
	}
	if tag == "" {
		return nil, fmt.Errorf("bookmark: tag must not be empty")
	}
	b := &Bookmark{mode: mode, tag: tag}
	switch mode {
	case ModeIndex:
		if len(indices) == 0 {
			return nil, fmt.Errorf("bookmark: at least one index required")
		}
		b.indices = make(map[int]bool, len(indices))
		for _, i := range indices {
			if i < 1 {
				return nil, fmt.Errorf("bookmark: index must be >= 1, got %d", i)
			}
			b.indices[i] = true
		}
	case ModeRegex:
		if pattern == "" {
			return nil, fmt.Errorf("bookmark: pattern required for regex mode")
		}
		re, err := regexp.Compile(pattern)
		if err != nil {
			return nil, fmt.Errorf("bookmark: invalid pattern: %w", err)
		}
		b.re = re
	}
	return b, nil
}

// Apply returns the line with the bookmark tag prepended if it matches,
// or the original line otherwise. lineNum is 1-based.
func (b *Bookmark) Apply(line string, lineNum int) string {
	if b.mode == ModeNone {
		return line
	}
	if b.matches(line, lineNum) {
		return b.tag + line
	}
	return line
}

func (b *Bookmark) matches(line string, lineNum int) bool {
	switch b.mode {
	case ModeIndex:
		return b.indices[lineNum]
	case ModeRegex:
		return b.re.MatchString(line)
	}
	return false
}
