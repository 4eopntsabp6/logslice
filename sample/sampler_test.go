package sample

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
		{"nth", ModeNth},
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
		t.Error("expected error for unknown mode, got nil")
	}
}

func TestNew_InvalidInterval(t *testing.T) {
	_, err := New(ModeNth, 0)
	if err == nil {
		t.Error("expected error for n=0, got nil")
	}
}

func TestNew_ModeNoneIgnoresN(t *testing.T) {
	s, err := New(ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if s == nil {
		t.Fatal("expected non-nil sampler")
	}
}

func TestKeep_ModeNone_AllowsAll(t *testing.T) {
	s, _ := New(ModeNone, 1)
	for i := 0; i < 10; i++ {
		if !s.Keep() {
			t.Errorf("ModeNone: Keep() returned false at iteration %d", i)
		}
	}
}

func TestKeep_ModeNth_EveryThird(t *testing.T) {
	s, _ := New(ModeNth, 3)
	results := make([]bool, 9)
	for i := range results {
		results[i] = s.Keep()
	}
	// positions 2,5,8 (0-indexed) should be true
	expected := []bool{false, false, true, false, false, true, false, false, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("index %d: got %v, want %v", i, got, expected[i])
		}
	}
}

func TestKeep_ModeNth_One(t *testing.T) {
	s, _ := New(ModeNth, 1)
	for i := 0; i < 5; i++ {
		if !s.Keep() {
			t.Errorf("n=1: Keep() should always return true, failed at %d", i)
		}
	}
}

func TestReset(t *testing.T) {
	s, _ := New(ModeNth, 3)
	s.Keep() // 1
	s.Keep() // 2
	s.Reset()
	// after reset counter is 0, so next two calls should be false
	if s.Keep() {
		t.Error("expected false after reset on first call")
	}
	if s.Keep() {
		t.Error("expected false after reset on second call")
	}
	if !s.Keep() {
		t.Error("expected true on third call after reset")
	}
}
