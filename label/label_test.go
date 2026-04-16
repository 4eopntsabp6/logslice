package label

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"none", ModeNone},
		{"", ModeNone},
		{"prefix", ModePrefix},
		{"append", ModeAppend},
		{"PREFIX", ModePrefix},
	}
	for _, c := range cases {
		got, err := ParseMode(c.input)
		if err != nil || got != c.want {
			t.Errorf("ParseMode(%q) = %v, %v; want %v, nil", c.input, got, err, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("wrap")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoPattern(t *testing.T) {
	l, err := New(ModeNone, "", "")
	if err != nil || l == nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNew_MissingLabel(t *testing.T) {
	_, err := New(ModePrefix, "", "error")
	if err == nil {
		t.Error("expected error for empty label")
	}
}

func TestNew_MissingPattern(t *testing.T) {
	_, err := New(ModeAppend, "WARN", "")
	if err == nil {
		t.Error("expected error for empty pattern")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := New(ModePrefix, "TAG", "[invalid")
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}

func TestApply_ModeNone(t *testing.T) {
	l, _ := New(ModeNone, "", "")
	got := l.Apply("hello world")
	if got != "hello world" {
		t.Errorf("expected unchanged line, got %q", got)
	}
}

func TestApply_Prefix_Match(t *testing.T) {
	l, _ := New(ModePrefix, "ERROR", "(?i)error")
	got := l.Apply("error occurred")
	if got != "[ERROR] error occurred" {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestApply_Append_Match(t *testing.T) {
	l, _ := New(ModeAppend, "WARN", "(?i)warn")
	got := l.Apply("warning: low disk")
	if got != "warning: low disk [WARN]" {
		t.Errorf("unexpected output: %q", got)
	}
}

func TestApply_NoMatch(t *testing.T) {
	l, _ := New(ModePrefix, "ERROR", "(?i)error")
	got := l.Apply("all good")
	if got != "all good" {
		t.Errorf("expected unchanged line, got %q", got)
	}
}
