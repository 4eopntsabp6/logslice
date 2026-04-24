package redact_test

import (
	"testing"

	"github.com/yourorg/logslice/redact"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  redact.Mode
	}{
		{"", redact.ModeNone},
		{"none", redact.ModeNone},
		{"pattern", redact.ModePattern},
	}
	for _, tc := range cases {
		got, err := redact.ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := redact.ParseMode("scrub")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoPattern(t *testing.T) {
	r, err := redact.New(redact.ModeNone, "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Enabled() {
		t.Error("expected Enabled() == false for ModeNone")
	}
}

func TestNew_MissingPattern(t *testing.T) {
	_, err := redact.New(redact.ModePattern, "", "")
	if err == nil {
		t.Error("expected error when pattern is empty")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := redact.New(redact.ModePattern, "[invalid", "")
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}

func TestApply_ModeNone(t *testing.T) {
	r, _ := redact.New(redact.ModeNone, "", "")
	line := "password=secret123"
	if got := r.Apply(line); got != line {
		t.Errorf("Apply() = %q, want %q", got, line)
	}
}

func TestApply_RedactsMatch(t *testing.T) {
	r, err := redact.New(redact.ModePattern, `password=\S+`, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := r.Apply("user=admin password=secret123 host=localhost")
	want := "user=admin [REDACTED] host=localhost"
	if got != want {
		t.Errorf("Apply() = %q, want %q", got, want)
	}
}

func TestApply_CustomPlaceholder(t *testing.T) {
	r, err := redact.New(redact.ModePattern, `\d{4}-\d{4}-\d{4}-\d{4}`, "****")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := r.Apply("card: 1234-5678-9012-3456")
	want := "card: ****"
	if got != want {
		t.Errorf("Apply() = %q, want %q", got, want)
	}
}

func TestApply_NoMatch(t *testing.T) {
	r, _ := redact.New(redact.ModePattern, `token=\S+`, "")
	line := "level=info msg=started"
	if got := r.Apply(line); got != line {
		t.Errorf("Apply() changed non-matching line: %q", got)
	}
}
