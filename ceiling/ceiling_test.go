package ceiling_test

import (
	"regexp"
	"testing"

	"github.com/yourorg/logslice/ceiling"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		in   string
		want ceiling.Mode
	}{
		{"", ceiling.ModeNone},
		{"none", ceiling.ModeNone},
		{"cap", ceiling.ModeCap},
	}
	for _, c := range cases {
		got, err := ceiling.ParseMode(c.in)
		if err != nil {
			t.Fatalf("ParseMode(%q) unexpected error: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ceiling.ParseMode("unknown")
	if err == nil {
		t.Fatal("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresParams(t *testing.T) {
	f, err := ceiling.New(ceiling.ModeNone, nil, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Keep("anything") {
		t.Error("ModeNone should keep all lines")
	}
}

func TestNew_InvalidMax(t *testing.T) {
	_, err := ceiling.New(ceiling.ModeCap, nil, 0)
	if err == nil {
		t.Fatal("expected error for max=0")
	}
}

func TestKeep_CapWithNoPattern(t *testing.T) {
	f, err := ceiling.New(ceiling.ModeCap, nil, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := []string{"a", "b", "c", "d", "e"}
	expected := []bool{true, true, true, false, false}
	for i, line := range lines {
		if got := f.Keep(line); got != expected[i] {
			t.Errorf("line %d Keep(%q) = %v, want %v", i, line, got, expected[i])
		}
	}
}

func TestKeep_CapWithPattern(t *testing.T) {
	re := regexp.MustCompile(`ERROR`)
	f, err := ceiling.New(ceiling.ModeCap, re, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	cases := []struct {
		line string
		want bool
	}{
		{"INFO hello", false},
		{"ERROR first", true},
		{"DEBUG skip", false},
		{"ERROR second", true},
		{"ERROR third", false}, // cap reached
	}
	for _, c := range cases {
		if got := f.Keep(c.line); got != c.want {
			t.Errorf("Keep(%q) = %v, want %v", c.line, got, c.want)
		}
	}
}
