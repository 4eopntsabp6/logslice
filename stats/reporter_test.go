package stats_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/logslice/stats"
)

func buildStats(matched, dropped int, bytes int64) *stats.Stats {
	s := stats.New()
	for i := 0; i < matched; i++ {
		s.RecordLine(true, int(bytes)/max(matched+dropped, 1))
	}
	for i := 0; i < dropped; i++ {
		s.RecordLine(false, 0)
	}
	s.Finish()
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestReport_ContainsLabels(t *testing.T) {
	var buf bytes.Buffer
	r := stats.NewReporter(&buf)
	s := buildStats(3, 1, 200)

	if err := r.Report(s); err != nil {
		t.Fatalf("Report returned error: %v", err)
	}

	out := buf.String()
	for _, label := range []string{"Lines read:", "Lines matched:", "Lines dropped:", "Bytes read:", "Match rate:", "Duration:"} {
		if !strings.Contains(out, label) {
			t.Errorf("Report output missing label %q\nGot:\n%s", label, out)
		}
	}
}

func TestReport_CorrectValues(t *testing.T) {
	var buf bytes.Buffer
	r := stats.NewReporter(&buf)
	s := buildStats(7, 3, 0)

	if err := r.Report(s); err != nil {
		t.Fatalf("Report returned error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "10") {
		t.Errorf("expected total line count 10 in output, got:\n%s", out)
	}
	if !strings.Contains(out, "70.0%") {
		t.Errorf("expected match rate 70.0%% in output, got:\n%s", out)
	}
}

func TestReportInline_SingleLine(t *testing.T) {
	var buf bytes.Buffer
	r := stats.NewReporter(&buf)
	s := buildStats(2, 2, 80)

	if err := r.ReportInline(s); err != nil {
		t.Fatalf("ReportInline returned error: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 1 {
		t.Errorf("expected 1 line, got %d: %v", len(lines), lines)
	}
}
