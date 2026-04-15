package slicer_test

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/user/logslice/slicer"
)

func buildResults() []slicer.Result {
	return []slicer.Result{
		{Line: "2024-01-15T10:01:00Z DEBUG request received", Timestamp: time.Now(), LineNum: 2},
		{Line: "2024-01-15T10:02:00Z ERROR database timeout", Timestamp: time.Now(), LineNum: 3},
	}
}

func TestWrite_Plain(t *testing.T) {
	var buf bytes.Buffer
	results := buildResults()
	n, err := slicer.Write(&buf, results, slicer.WriteOptions{Format: slicer.FormatPlain})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 2 {
		t.Errorf("expected 2 lines written, got %d", n)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 output lines, got %d", len(lines))
	}
	if lines[0] != results[0].Line {
		t.Errorf("line mismatch: got %q, want %q", lines[0], results[0].Line)
	}
}

func TestWrite_Numbered(t *testing.T) {
	var buf bytes.Buffer
	results := buildResults()
	_, err := slicer.Write(&buf, results, slicer.WriteOptions{Format: slicer.FormatNumbered})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "2\t") {
		t.Errorf("expected line number prefix '2\\t' in output: %q", output)
	}
	if !strings.Contains(output, "3\t") {
		t.Errorf("expected line number prefix '3\\t' in output: %q", output)
	}
}

func TestWrite_Empty(t *testing.T) {
	var buf bytes.Buffer
	n, err := slicer.Write(&buf, nil, slicer.WriteOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected 0 lines written, got %d", n)
	}
}

func TestSummary(t *testing.T) {
	results := buildResults()
	s := slicer.Summary(10, results)
	if !strings.Contains(s, "2 of 10") {
		t.Errorf("unexpected summary: %q", s)
	}

	empty := slicer.Summary(10, nil)
	if !strings.Contains(empty, "No matching") {
		t.Errorf("unexpected empty summary: %q", empty)
	}
}
