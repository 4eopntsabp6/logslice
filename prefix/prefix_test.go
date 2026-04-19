package prefix

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
		{"text", ModeText},
		{"linenum", ModeLineNumber},
		{"TEXT", ModeText},
	}
	for _, c := range cases {
		got, err := ParseMode(c.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("bold")
	if err == nil {
		t.Error("expected error for unknown mode, got nil")
	}
}

func TestNew_ModeNoneNoText(t *testing.T) {
	p, err := New(ModeNone, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := p.Apply("hello"); got != "hello" {
		t.Errorf("Apply = %q, want %q", got, "hello")
	}
}

func TestNew_ModeText_MissingText(t *testing.T) {
	_, err := New(ModeText, "")
	if err == nil {
		t.Error("expected error for empty text in ModeText")
	}
}

func TestApply_ModeText(t *testing.T) {
	p, err := New(ModeText, "[INFO] ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := p.Apply("something happened")
	want := "[INFO] something happened"
	if got != want {
		t.Errorf("Apply = %q, want %q", got, want)
	}
}

func TestApply_ModeLineNumber(t *testing.T) {
	p, err := New(ModeLineNumber, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cases := []struct {
		line string
		want string
	}{
		{"first", "1: first"},
		{"second", "2: second"},
		{"third", "3: third"},
	}
	for _, c := range cases {
		if got := p.Apply(c.line); got != c.want {
			t.Errorf("Apply(%q) = %q, want %q", c.line, got, c.want)
		}
	}
}

func TestReset_ResetsCounter(t *testing.T) {
	p, _ := New(ModeLineNumber, "")
	p.Apply("a")
	p.Apply("b")
	p.Reset()
	got := p.Apply("c")
	if got != "1: c" {
		t.Errorf("after Reset Apply = %q, want %q", got, "1: c")
	}
}
