package filter

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewRegexFilter_Valid(t *testing.T) {
	_, err := NewRegexFilter(`ERROR|WARN`)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestNewRegexFilter_Invalid(t *testing.T) {
	_, err := NewRegexFilter(`[invalid`)
	if err == nil {
		t.Fatal("expected error for invalid pattern, got nil")
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		pattern string
		line    string
		want    bool
	}{
		{`ERROR`, "2024-01-01 ERROR something failed", true},
		{`ERROR`, "2024-01-01 INFO all good", false},
		{`ERROR|WARN`, "2024-01-01 WARN disk space low", true},
		{`user_id=\d+`, "request user_id=42 processed", true},
		{`user_id=\d+`, "request processed", false},
	}
	for _, tt := range tests {
		f, err := NewRegexFilter(tt.pattern)
		if err != nil {
			t.Fatalf("failed to compile pattern %q: %v", tt.pattern, err)
		}
		got := f.Match(tt.line)
		if got != tt.want {
			t.Errorf("Match(%q) with pattern %q = %v, want %v", tt.line, tt.pattern, got, tt.want)
		}
	}
}

func TestApply(t *testing.T) {
	input := strings.Join([]string{
		"2024-01-01 INFO service started",
		"2024-01-01 ERROR connection refused",
		"2024-01-01 WARN retrying request",
		"2024-01-01 INFO request completed",
		"2024-01-01 ERROR timeout exceeded",
	}, "\n")

	f, err := NewRegexFilter(`ERROR`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var out bytes.Buffer
	count, err := f.Apply(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("Apply returned error: %v", err)
	}
	if count != 2 {
		t.Errorf("expected 2 matched lines, got %d", count)
	}
	result := out.String()
	if !strings.Contains(result, "connection refused") {
		t.Error("expected output to contain 'connection refused'")
	}
	if !strings.Contains(result, "timeout exceeded") {
		t.Error("expected output to contain 'timeout exceeded'")
	}
	if strings.Contains(result, "INFO") {
		t.Error("expected output to not contain INFO lines")
	}
}
