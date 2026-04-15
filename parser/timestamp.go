package parser

import (
	"fmt"
	"regexp"
	"time"
)

// Common timestamp formats found in structured logs
var timestampFormats = []string{
	time.RFC3339,
	time.RFC3339Nano,
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04:05.000",
	"2006-01-02T15:04:05.000Z",
	"02/Jan/2006:15:04:05 -0700",
	"Jan 2 15:04:05",
	"Jan  2 15:04:05",
}

// timestampRegex matches common ISO-like timestamps at the start of a log line
var timestampRegex = regexp.MustCompile(
	`(\d{4}-\d{2}-\d{2}[T ]\d{2}:\d{2}:\d{2}(?:[.,]\d+)?(?:Z|[+-]\d{2}:?\d{2})?)`,
)

// ExtractTimestamp attempts to parse a timestamp from a log line.
// Returns the parsed time and true if successful, otherwise zero time and false.
func ExtractTimestamp(line string) (time.Time, bool) {
	matches := timestampRegex.FindStringSubmatch(line)
	if len(matches) < 2 {
		return time.Time{}, false
	}

	raw := matches[1]
	for _, format := range timestampFormats {
		if t, err := time.Parse(format, raw); err == nil {
			return t, true
		}
	}

	return time.Time{}, false
}

// ParseTimeArg parses a user-supplied time argument into a time.Time.
// Supports RFC3339 and common shorthand formats.
func ParseTimeArg(s string) (time.Time, error) {
	for _, format := range timestampFormats {
		if t, err := time.Parse(format, s); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unrecognized time format: %q", s)
}
