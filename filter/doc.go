// Package filter provides log line filtering utilities for logslice.
//
// Currently supported filters:
//
//   - RegexFilter: filters log lines using a compiled regular expression.
//     Useful for extracting lines matching specific log levels (e.g. ERROR, WARN),
//     request IDs, user identifiers, or any arbitrary pattern.
//
// Example usage:
//
//	f, err := filter.NewRegexFilter(`ERROR|WARN`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	count, err := f.Apply(inputReader, outputWriter)
//
// Filters are designed to compose with the time-window slicing provided by
// the parser package, allowing callers to first narrow a log file by time
// range and then further refine results by pattern matching.
package filter
