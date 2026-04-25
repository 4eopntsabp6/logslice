package repeat_test

import (
	"testing"

	"github.com/yourorg/logslice/repeat"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  repeat.Mode
	}{
		{"", repeat.ModeNone},
		{"none", repeat.ModeNone},
		{"limit", repeat.ModeLimit},
	}
	for _, tc := range cases {
		got, err := repeat.ParseMode(tc.input)
		if err != nil {
			t.Fatalf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := repeat.ParseMode("bogus")
	if err == nil {
		t.Fatal("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoPattern(t *testing.T) {
	f, err := repeat.New(repeat.ModeNone, "", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Keep("anything") {
		t.Error("ModeNone should keep all lines")
	}
}

func TestNew_MissingPattern(t *testing.T) {
	_, err := repeat.New(repeat.ModeLimit, "", 3)
	if err == nil {
		t.Fatal("expected error for missing pattern")
	}
}

func TestNew_InvalidMax(t *testing.T) {
	_, err := repeat.New(repeat.ModeLimit, `ERROR`, 0)
	if err == nil {
		t.Fatal("expected error for max <= 0")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := repeat.New(repeat.ModeLimit, `[invalid`, 2)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestKeep_AllowsUpToMax(t *testing.T) {
	f, err := repeat.New(repeat.ModeLimit, `ERROR`, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	line := "2024-01-01 ERROR something went wrong"
	for i := 1; i <= 3; i++ {
		if !f.Keep(line) {
			t.Errorf("expected Keep=true on occurrence %d", i)
		}
	}
	if f.Keep(line) {
		t.Error("expected Keep=false after max occurrences exceeded")
	}
}

func TestKeep_NonMatchingLinesAlwaysPass(t *testing.T) {
	f, err := repeat.New(repeat.ModeLimit, `ERROR`, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 0; i < 10; i++ {
		if !f.Keep("INFO all good") {
			t.Error("non-matching line should always pass")
		}
	}
}

func TestKeep_TracksMatchesSeparately(t *testing.T) {
	f, err := repeat.New(repeat.ModeLimit, `(ERROR|WARN)`, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// ERROR and WARN are tracked independently
	for i := 0; i < 2; i++ {
		if !f.Keep("ERROR foo") {
			t.Errorf("ERROR: expected keep on %d", i)
		}
		if !f.Keep("WARN bar") {
			t.Errorf("WARN: expected keep on %d", i)
		}
	}
	if f.Keep("ERROR foo") {
		t.Error("ERROR should be suppressed after max")
	}
	if f.Keep("WARN bar") {
		t.Error("WARN should be suppressed after max")
	}
}
