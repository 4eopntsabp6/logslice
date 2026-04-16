package transform_test

import (
	"testing"

	"github.com/user/logslice/transform"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  transform.Mode
	}{
		{"", transform.ModeNone},
		{"none", transform.ModeNone},
		{"upper", transform.ModeUpper},
		{"lower", transform.ModeLower},
		{"trim", transform.ModeTrimSpace},
	}
	for _, c := range cases {
		got, err := transform.ParseMode(c.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := transform.ParseMode("reverse")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestApply_ModeNone(t *testing.T) {
	tr := transform.New(transform.ModeNone)
	if got := tr.Apply("Hello World"); got != "Hello World" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestApply_Upper(t *testing.T) {
	tr := transform.New(transform.ModeUpper)
	if got := tr.Apply("hello"); got != "HELLO" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestApply_Lower(t *testing.T) {
	tr := transform.New(transform.ModeLower)
	if got := tr.Apply("HELLO"); got != "hello" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestApply_Trim(t *testing.T) {
	tr := transform.New(transform.ModeTrimSpace)
	if got := tr.Apply("  hello  "); got != "hello" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestEnabled(t *testing.T) {
	if transform.New(transform.ModeNone).Enabled() {
		t.Error("ModeNone should not be enabled")
	}
	if !transform.New(transform.ModeUpper).Enabled() {
		t.Error("ModeUpper should be enabled")
	}
}
