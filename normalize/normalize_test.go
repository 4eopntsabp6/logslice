package normalize

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"none", ModeNone},
		{"trim", ModeTrim},
		{"collapse", ModeCollapse},
		{"all", ModeAll},
		{"TRIM", ModeTrim},
		{"ALL", ModeAll},
	}
	for _, tc := range cases {
		got, err := ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("squash")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_InvalidMode(t *testing.T) {
	_, err := New(Mode(99))
	if err == nil {
		t.Error("expected error for invalid mode value")
	}
}

func TestApply_ModeNone(t *testing.T) {
	n, _ := New(ModeNone)
	line := "  hello   world  "
	if got := n.Apply(line); got != line {
		t.Errorf("ModeNone modified line: %q", got)
	}
}

func TestApply_ModeTrim(t *testing.T) {
	n, _ := New(ModeTrim)
	got := n.Apply("  hello   world  ")
	want := "hello   world"
	if got != want {
		t.Errorf("ModeTrim got %q, want %q", got, want)
	}
}

func TestApply_ModeCollapse(t *testing.T) {
	n, _ := New(ModeCollapse)
	got := n.Apply("  hello   world  ")
	want := " hello world "
	if got != want {
		t.Errorf("ModeCollapse got %q, want %q", got, want)
	}
}

func TestApply_ModeAll(t *testing.T) {
	n, _ := New(ModeAll)
	got := n.Apply("  hello   world  ")
	want := "hello world"
	if got != want {
		t.Errorf("ModeAll got %q, want %q", got, want)
	}
}

func TestEnabled(t *testing.T) {
	none, _ := New(ModeNone)
	if none.Enabled() {
		t.Error("ModeNone should not be enabled")
	}
	all, _ := New(ModeAll)
	if !all.Enabled() {
		t.Error("ModeAll should be enabled")
	}
}
