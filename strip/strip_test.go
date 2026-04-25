package strip_test

import (
	"testing"

	"github.com/user/logslice/strip"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  strip.Mode
	}{
		{"", strip.ModeNone},
		{"none", strip.ModeNone},
		{"ansi", strip.ModeANSI},
		{"whitespace", strip.ModeWhitespace},
		{"both", strip.ModeBoth},
		{"ANSI", strip.ModeANSI},
		{"Both", strip.ModeBoth},
	}
	for _, tc := range cases {
		got, err := strip.ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := strip.ParseMode("unknown")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestApply_ModeNone(t *testing.T) {
	s, _ := strip.New(strip.ModeNone)
	line := "  \x1b[31mhello\x1b[0m  "
	if got := s.Apply(line); got != line {
		t.Errorf("ModeNone modified line: %q", got)
	}
}

func TestApply_ModeANSI(t *testing.T) {
	s, _ := strip.New(strip.ModeANSI)
	got := s.Apply("\x1b[31merror\x1b[0m: bad input")
	want := "error: bad input"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestApply_ModeWhitespace(t *testing.T) {
	s, _ := strip.New(strip.ModeWhitespace)
	got := s.Apply("   hello world   ")
	want := "hello world"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestApply_ModeBoth(t *testing.T) {
	s, _ := strip.New(strip.ModeBoth)
	got := s.Apply("  \x1b[32mok\x1b[0m  ")
	want := "ok"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestEnabled(t *testing.T) {
	none, _ := strip.New(strip.ModeNone)
	if none.Enabled() {
		t.Error("ModeNone should not be enabled")
	}
	ansi, _ := strip.New(strip.ModeANSI)
	if !ansi.Enabled() {
		t.Error("ModeANSI should be enabled")
	}
}
