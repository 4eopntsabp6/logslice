package stats

import (
	"fmt"
	"io"
	"text/tabwriter"
)

// Reporter writes formatted stats output to a writer.
type Reporter struct {
	w io.Writer
}

// NewReporter creates a Reporter that writes to w.
func NewReporter(w io.Writer) *Reporter {
	return &Reporter{w: w}
}

// Report writes a tabular summary of s to the reporter's writer.
func (r *Reporter) Report(s *Stats) error {
	tw := tabwriter.NewWriter(r.w, 0, 0, 2, ' ', 0)

	lines := []struct {
		label string
		value string
	}{
		{"Lines read:", fmt.Sprintf("%d", s.LinesRead)},
		{"Lines matched:", fmt.Sprintf("%d", s.LinesMatched)},
		{"Lines dropped:", fmt.Sprintf("%d", s.LinesDropped)},
		{"Bytes read:", fmt.Sprintf("%d", s.BytesRead)},
		{"Match rate:", fmt.Sprintf("%.1f%%", s.MatchRate())},
		{"Duration:", s.Duration().Round(1e6).String()},
	}

	for _, l := range lines {
		if _, err := fmt.Fprintf(tw, "  %s\t%s\n", l.label, l.value); err != nil {
			return err
		}
	}

	return tw.Flush()
}

// ReportInline writes a single-line summary to the reporter's writer.
func (r *Reporter) ReportInline(s *Stats) error {
	_, err := fmt.Fprintln(r.w, s.String())
	return err
}
