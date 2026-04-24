package entropy

import (
	"math"
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"", ModeNone},
		{"none", ModeNone},
		{"above", ModeAbove},
		{"below", ModeBelow},
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
	_, err := ParseMode("random")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresThreshold(t *testing.T) {
	f, err := New(ModeNone, -99)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Keep("hello") {
		t.Error("ModeNone should keep all lines")
	}
}

func TestNew_InvalidThreshold(t *testing.T) {
	_, err := New(ModeAbove, -1)
	if err == nil {
		t.Error("expected error for negative threshold")
	}
}

func TestKeep_Above(t *testing.T) {
	f, _ := New(ModeAbove, 3.0)
	// high entropy string
	if !f.Keep("aB3!xZ9@qW") {
		t.Error("high-entropy line should be kept")
	}
	// low entropy string (all same char)
	if f.Keep("aaaaaaaaaa") {
		t.Error("low-entropy line should be dropped")
	}
}

func TestKeep_Below(t *testing.T) {
	f, _ := New(ModeBelow, 1.5)
	if !f.Keep("aaaaaab") {
		t.Error("low-entropy line should be kept")
	}
	if f.Keep("aB3!xZ9@qW2#mN") {
		t.Error("high-entropy line should be dropped")
	}
}

func TestScore_Empty(t *testing.T) {
	if Score("") != 0 {
		t.Error("empty string should have entropy 0")
	}
}

func TestScore_SingleChar(t *testing.T) {
	if Score("aaaa") != 0 {
		t.Error("single-character string should have entropy 0")
	}
}

func TestScore_TwoEqual(t *testing.T) {
	got := Score("aabb")
	if math.Abs(got-1.0) > 1e-9 {
		t.Errorf("expected entropy 1.0, got %f", got)
	}
}
