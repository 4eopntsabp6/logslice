package offset

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
		{"skip", ModeSkip},
		{"start", ModeStart},
	}
	for _, tc := range cases {
		got, err := ParseMode(tc.input)
		if err != nil {
			t.Fatalf("ParseMode(%q): unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("bad")
	if err == nil {
		t.Fatal("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresN(t *testing.T) {
	f, err := New(ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Keep("line") {
		t.Error("ModeNone should keep all lines")
	}
}

func TestNew_InvalidN(t *testing.T) {
	_, err := New(ModeSkip, 0)
	if err == nil {
		t.Fatal("expected error for n=0")
	}
}

func TestKeep_ModeSkip(t *testing.T) {
	f, _ := New(ModeSkip, 3)
	results := make([]bool, 6)
	for i := range results {
		results[i] = f.Keep("line")
	}
	// first 3 skipped, rest kept
	expected := []bool{false, false, false, true, true, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("line %d: got %v, want %v", i+1, got, expected[i])
		}
	}
}

func TestKeep_ModeStart(t *testing.T) {
	f, _ := New(ModeStart, 3)
	results := make([]bool, 5)
	for i := range results {
		results[i] = f.Keep("line")
	}
	// lines 1,2 dropped; from line 3 onward kept
	expected := []bool{false, false, true, true, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("line %d: got %v, want %v", i+1, got, expected[i])
		}
	}
}
