package slicer_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/logslice/slicer"
)

const sampleLogs = `2024-01-15T10:00:00Z INFO  service started
2024-01-15T10:01:00Z DEBUG request received id=42
2024-01-15T10:02:00Z ERROR database timeout retries=3
2024-01-15T10:03:00Z INFO  request completed id=42
2024-01-15T10:04:00Z WARN  memory usage high pct=85
`

func mustParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func TestSlice_NoWindow(t *testing.T) {
	r := strings.NewReader(sampleLogs)
	results, err := slicer.Slice(r, slicer.Options{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 5 {
		t.Errorf("expected 5 results, got %d", len(results))
	}
}

func TestSlice_FromOnly(t *testing.T) {
	r := strings.NewReader(sampleLogs)
	opts := slicer.Options{
		From:      mustParse("2024-01-15T10:02:00Z"),
		Inclusive: true,
	}
	results, err := slicer.Slice(r, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
}

func TestSlice_ToOnly(t *testing.T) {
	r := strings.NewReader(sampleLogs)
	opts := slicer.Options{
		To:        mustParse("2024-01-15T10:02:00Z"),
		Inclusive: true,
	}
	results, err := slicer.Slice(r, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
}

func TestSlice_FromTo(t *testing.T) {
	r := strings.NewReader(sampleLogs)
	opts := slicer.Options{
		From:      mustParse("2024-01-15T10:01:00Z"),
		To:        mustParse("2024-01-15T10:03:00Z"),
		Inclusive: true,
	}
	results, err := slicer.Slice(r, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
	if results[0].LineNum != 2 {
		t.Errorf("expected first match at line 2, got %d", results[0].LineNum)
	}
}

func TestSlice_NoMatch(t *testing.T) {
	r := strings.NewReader(sampleLogs)
	opts := slicer.Options{
		From: mustParse("2024-01-15T11:00:00Z"),
		To:   mustParse("2024-01-15T12:00:00Z"),
	}
	results, err := slicer.Slice(r, opts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}
