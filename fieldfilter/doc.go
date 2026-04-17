// Package fieldfilter provides key=value field-based filtering for log lines.
//
// It supports four modes:
//
//   - none:   all lines pass through (filtering disabled)
//   - exact:  the extracted field value must exactly match the target
//   - prefix: the extracted field value must begin with the target
//   - regex:  the extracted field value must match the compiled regex
//
// Fields are extracted from lines formatted as space-separated key=value pairs,
// e.g.: ts=2024-01-01T00:00:00Z level=error msg=something
//
// Example usage:
//
//	f, err := fieldfilter.New(fieldfilter.ModeExact, "level", "error")
//	if err != nil { ... }
//	if f.Match(line) { ... }
package fieldfilter
