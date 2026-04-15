// Package stats provides runtime statistics collection for logslice operations.
package stats

import (
	"fmt"
	"time"
)

// Stats holds counters and timing information for a slice operation.
type Stats struct {
	LinesRead    int
	LinesMatched int
	LinesDropped int
	BytesRead    int64
	StartTime    time.Time
	EndTime      time.Time
}

// New returns a new Stats instance with StartTime set to now.
func New() *Stats {
	return &Stats{StartTime: time.Now()}
}

// Finish records the end time of the operation.
func (s *Stats) Finish() {
	s.EndTime = time.Now()
}

// Duration returns the elapsed time between Start and Finish.
func (s *Stats) Duration() time.Duration {
	if s.EndTime.IsZero() {
		return time.Since(s.StartTime)
	}
	return s.EndTime.Sub(s.StartTime)
}

// RecordLine updates counters for a single line read.
func (s *Stats) RecordLine(matched bool, bytes int) {
	s.LinesRead++
	s.BytesRead += int64(bytes)
	if matched {
		s.LinesMatched++
	} else {
		s.LinesDropped++
	}
}

// MatchRate returns the percentage of lines matched out of lines read.
func (s *Stats) MatchRate() float64 {
	if s.LinesRead == 0 {
		return 0.0
	}
	return float64(s.LinesMatched) / float64(s.LinesRead) * 100.0
}

// String returns a human-readable summary of the stats.
func (s *Stats) String() string {
	return fmt.Sprintf(
		"read=%d matched=%d dropped=%d bytes=%d duration=%s match_rate=%.1f%%",
		s.LinesRead, s.LinesMatched, s.LinesDropped,
		s.BytesRead, s.Duration().Round(time.Millisecond), s.MatchRate(),
	)
}
