package bookmark

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
		{"index", ModeIndex},
		{"regex", ModeRegex},
	}
	for _, c := range cases {
		got, err := ParseMode(c.input)
		if err != nil || got != c.want {
			t.Errorf("ParseMode(%q) = %v, %v; want %v, nil", c.input, got, err, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("unknown")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNew_ModeNone(t *testing.T) {
	b, err := New(ModeNone, "", nil, "")
	if err != nil || b == nil {
		t.Fatalf("unexpected: %v", err)
	}
	if got := b.Apply("hello", 1); got != "hello" {
		t.Errorf("expected passthrough, got %q", got)
	}
}

func TestNew_ModeIndex_MissingTag(t *testing.T) {
	_, err := New(ModeIndex, "", []int{1}, "")
	if err == nil {
		t.Fatal("expected error for missing tag")
	}
}

func TestNew_ModeIndex_NoIndices(t *testing.T) {
	_, err := New(ModeIndex, ">>>", nil, "")
	if err == nil {
		t.Fatal("expected error for no indices")
	}
}

func TestNew_ModeIndex_InvalidIndex(t *testing.T) {
	_, err := New(ModeIndex, ">>>", []int{0}, "")
	if err == nil {
		t.Fatal("expected error for index < 1")
	}
}

func TestApply_ModeIndex(t *testing.T) {
	b, err := New(ModeIndex, "[*] ", []int{2, 4}, "")
	if err != nil {
		t.Fatal(err)
	}
	if got := b.Apply("line one", 1); got != "line one" {
		t.Errorf("expected no bookmark, got %q", got)
	}
	if got := b.Apply("line two", 2); got != "[*] line two" {
		t.Errorf("expected bookmark, got %q", got)
	}
}

func TestNew_ModeRegex_MissingPattern(t *testing.T) {
	_, err := New(ModeRegex, ">>>", nil, "")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNew_ModeRegex_InvalidPattern(t *testing.T) {
	_, err := New(ModeRegex, ">>>", nil, "[invalid")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestApply_ModeRegex(t *testing.T) {
	b, err := New(ModeRegex, "[!] ", nil, `ERROR`)
	if err != nil {
		t.Fatal(err)
	}
	if got := b.Apply("INFO all good", 1); got != "INFO all good" {
		t.Errorf("expected no bookmark, got %q", got)
	}
	if got := b.Apply("ERROR something failed", 2); got != "[!] ERROR something failed" {
		t.Errorf("expected bookmark, got %q", got)
	}
}
