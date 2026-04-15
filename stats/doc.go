// Package stats provides lightweight runtime statistics for logslice operations.
//
// It tracks the number of lines read, matched, and dropped during a slice run,
// as well as byte throughput and wall-clock duration.
//
// Basic usage:
//
//	s := stats.New()
//	for _, line := range lines {
//		matched := filter.Match(line)
//		s.RecordLine(matched, len(line))
//	}
//	s.Finish()
//
//	r := stats.NewReporter(os.Stderr)
//	r.Report(s)
//
// The Reporter supports both tabular (Report) and inline (ReportInline) output
// formats, suitable for verbose mode or piped output respectively.
package stats
