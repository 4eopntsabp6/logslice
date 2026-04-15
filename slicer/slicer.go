// Package slicer provides functionality for extracting log lines
// from a reader within a specified time window.
package slicer

import (
	"bufio"
	"io"
	"time"

	"github.com/user/logslice/parser"
)

// Options configures the behavior of the log slicer.
type Options struct {
	From      time.Time
	To        time.Time
	Inclusive bool
}

// Result holds a matched log line and its parsed timestamp.
type Result struct {
	Line      string
	Timestamp time.Time
	LineNum   int
}

// Slice reads lines from r and returns those whose timestamps fall
// within the window defined by opts. Lines without a parseable
// timestamp are skipped unless both From and To are zero.
func Slice(r io.Reader, opts Options) ([]Result, error) {
	var results []Result
	scanner := bufio.NewScanner(r)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		ts, ok := parser.ExtractTimestamp(line)
		if !ok {
			if opts.From.IsZero() && opts.To.IsZero() {
				results = append(results, Result{Line: line, LineNum: lineNum})
			}
			continue
		}

		if inWindow(ts, opts) {
			results = append(results, Result{
				Line:      line,
				Timestamp: ts,
				LineNum:   lineNum,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

// inWindow reports whether ts falls within the time window defined by opts.
func inWindow(ts time.Time, opts Options) bool {
	if !opts.From.IsZero() {
		if opts.Inclusive {
			if ts.Before(opts.From) {
				return false
			}
		} else {
			if !ts.After(opts.From) {
				return false
			}
		}
	}
	if !opts.To.IsZero() {
		if opts.Inclusive {
			if ts.After(opts.To) {
				return false
			}
		} else {
			if !ts.Before(opts.To) {
				return false
			}
		}
	}
	return true
}
