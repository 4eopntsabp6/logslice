package merge

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		in   string
		want Mode
	}{
		{"none", ModeNone},
		{"", ModeNone},
		{"continuation", ModeContinuation},
		{"CONTINUATION", ModeContinuation},
	}
	for _, c := range cases {
		got, err := ParseMode(c.in)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("unknown")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoPattern(t *testing.T) {
	m, err := New(ModeNone, "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	out, ok := m.Feed("hello")
	if !ok || out != "hello" {
		t.Errorf("expected pass-through, got (%q, %v)", out, ok)
	}
}

func TestNew_MissingPattern(t *testing.T) {
	_, err := New(ModeContinuation, "", " ")
	if err == nil {
		t.Error("expected error for missing pattern")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := New(ModeContinuation, "[", " ")
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}

func TestFeed_MergesContinuationLines(t *testing.T) {
	m, err := New(ModeContinuation, `^\s+`, " ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// First root line — buffered, nothing emitted yet.
	_, ok := m.Feed("ERROR root cause")
	if ok {
		t.Error("expected no output for first root line")
	}

	// Continuation lines — still buffered.
	m.Feed("  at foo.go:10")
	m.Feed("  at bar.go:20")

	// Second root line triggers flush of first block.
	out, ok := m.Feed("INFO next event")
	if !ok {
		t.Error("expected merged output")
	}
	want := "ERROR root cause   at foo.go:10   at bar.go:20"
	if out != want {
		t.Errorf("got %q, want %q", out, want)
	}
}

func TestFlush_EmitsRemaining(t *testing.T) {
	m, _ := New(ModeContinuation, `^\s+`, "|")
	m.Feed("root")
	m.Feed("  cont")

	out, ok := m.Flush()
	if !ok {
		t.Fatal("expected flush to return a line")
	}
	if out != "root|  cont" {
		t.Errorf("got %q", out)
	}

	// Second flush should be empty.
	_, ok = m.Flush()
	if ok {
		t.Error("expected empty flush after drain")
	}
}
