// Package merge implements continuation-line merging for logslice.
//
// Some log formats (e.g. Java stack traces, Python tracebacks) emit a single
// logical event across multiple physical lines. The first line is the "root"
line and subsequent lines are "continuation" lines distinguished by a common
// prefix pattern such as leading whitespace.
//
// Usage:
//
//	m, err := merge.New(merge.ModeContinuation, `^\s+`, " ")
//	for _, line := range lines {
//		if out, ok := m.Feed(line); ok {
//			process(out)
//		}
//	}
//	if out, ok := m.Flush(); ok {
//		process(out)
//	}
package merge
