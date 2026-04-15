package filter

import (
	"bufio"
	"io"
	"regexp"
)

// RegexFilter holds a compiled regex pattern used to match log lines.
type RegexFilter struct {
	pattern *regexp.Regexp
}

// NewRegexFilter compiles the given pattern and returns a RegexFilter.
// Returns an error if the pattern is invalid.
func NewRegexFilter(pattern string) (*RegexFilter, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexFilter{pattern: re}, nil
}

// Match returns true if the line matches the compiled regex pattern.
func (f *RegexFilter) Match(line string) bool {
	return f.pattern.MatchString(line)
}

// Apply reads lines from r and writes matching lines to w.
// Returns the number of lines written and any error encountered.
func (f *RegexFilter) Apply(r io.Reader, w io.Writer) (int, error) {
	scanner := bufio.NewScanner(r)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if f.Match(line) {
			if _, err := io.WriteString(w, line+"\n"); err != nil {
				return count, err
			}
			count++
		}
	}
	return count, scanner.Err()
}
