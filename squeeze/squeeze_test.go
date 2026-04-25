package squeeze

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"", ModeNone},
		{"none", ModeNone},
		{"blank", ModeBlank},
		{"whitespace", ModeWhitespace},
		{"BLANK", ModeBlank},
	}
	for _, tc := range cases {
		got, err := ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("squash")
	if err == nil {
		t.Error("expected error for unknown mode, got nil")
	}
}

func TestKeep_ModeNone_AllowsAll(t *testing.T) {
	f, _ := New(ModeNone)
	lines := []string{"hello", "", "", "world", ""}
	for _, l := range lines {
		if !f.Keep(l) {
			t.Errorf("ModeNone: expected Keep(%q) = true", l)
		}
	}
}

func TestKeep_ModeBlank_CollapsesConsecutive(t *testing.T) {
	f, _ := New(ModeBlank)
	input := []string{"a", "", "", "", "b", "", "c"}
	want := []bool{true, true, false, false, true, true, true}
	for i, l := range input {
		got := f.Keep(l)
		if got != want[i] {
			t.Errorf("line %d %q: Keep = %v, want %v", i, l, got, want[i])
		}
	}
}

func TestKeep_ModeBlank_NonBlankResetsState(t *testing.T) {
	f, _ := New(ModeBlank)
	sequence := []struct {
		line string
		want bool
	}{
		{"", true},
		{"x", true},
		{"", true},
		{"", false},
	}
	for _, tc := range sequence {
		if got := f.Keep(tc.line); got != tc.want {
			t.Errorf("Keep(%q) = %v, want %v", tc.line, got, tc.want)
		}
	}
}

func TestKeep_ModeWhitespace_TreatsSpacesAsBlank(t *testing.T) {
	f, _ := New(ModeWhitespace)
	input := []string{"log", "   ", "\t", "more"}
	want := []bool{true, true, false, true}
	for i, l := range input {
		got := f.Keep(l)
		if got != want[i] {
			t.Errorf("line %d %q: Keep = %v, want %v", i, l, got, want[i])
		}
	}
}

func TestKeep_ModeWhitespace_EmptyLineIsBlank(t *testing.T) {
	f, _ := New(ModeWhitespace)
	if !f.Keep("") {
		t.Error("first empty line should be kept")
	}
	if f.Keep("") {
		t.Error("second consecutive empty line should be suppressed")
	}
}
