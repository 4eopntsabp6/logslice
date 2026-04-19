package indent

import (
	"strings"
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"", ModeNone},
		{"none", ModeNone},
		{"spaces", ModeSpaces},
		{"tabs", ModeTabs},
		{"SPACES", ModeSpaces},
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
	_, err := ParseMode("indent")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresN(t *testing.T) {
	i, err := New(ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if i.Enabled() {
		t.Error("expected Enabled() == false for ModeNone")
	}
}

func TestNew_InvalidN(t *testing.T) {
	_, err := New(ModeSpaces, 0)
	if err == nil {
		t.Error("expected error for n=0")
	}
}

func TestApply_ModeNone(t *testing.T) {
	i, _ := New(ModeNone, 0)
	line := "hello world"
	if got := i.Apply(line); got != line {
		t.Errorf("Apply() = %q, want %q", got, line)
	}
}

func TestApply_Spaces(t *testing.T) {
	i, err := New(ModeSpaces, 4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := i.Apply("log line")
	if !strings.HasPrefix(got, "    ") {
		t.Errorf("Apply() = %q, want 4-space prefix", got)
	}
}

func TestApply_Tabs(t *testing.T) {
	i, err := New(ModeTabs, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := i.Apply("log line")
	if !strings.HasPrefix(got, "\t\t") {
		t.Errorf("Apply() = %q, want 2-tab prefix", got)
	}
}

func TestEnabled(t *testing.T) {
	i, _ := New(ModeSpaces, 2)
	if !i.Enabled() {
		t.Error("expected Enabled() == true for ModeSpaces")
	}
}
