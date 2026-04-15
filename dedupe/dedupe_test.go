package dedupe_test

import (
	"testing"

	"github.com/yourorg/logslice/dedupe"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  dedupe.Mode
	}{
		{"none", dedupe.ModeNone},
		{"", dedupe.ModeNone},
		{"consecutive", dedupe.ModeConsecutive},
		{"global", dedupe.ModeGlobal},
	}
	for _, tc := range cases {
		got, ok := dedupe.ParseMode(tc.input)
		if !ok {
			t.Errorf("ParseMode(%q) returned ok=false", tc.input)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, ok := dedupe.ParseMode("unknown")
	if ok {
		t.Error("expected ok=false for unknown mode")
	}
}

func TestFilter_ModeNone_AllowsAll(t *testing.T) {
	f := dedupe.New(dedupe.ModeNone)
	lines := []string{"a", "a", "b", "a"}
	for _, l := range lines {
		if !f.Allow(l) {
			t.Errorf("ModeNone: expected Allow(%q)=true", l)
		}
	}
	if f.Dropped != 0 {
		t.Errorf("ModeNone: expected Dropped=0, got %d", f.Dropped)
	}
}

func TestFilter_ModeConsecutive(t *testing.T) {
	f := dedupe.New(dedupe.ModeConsecutive)
	results := []bool{
		f.Allow("a"), // true
		f.Allow("a"), // false — duplicate
		f.Allow("b"), // true
		f.Allow("b"), // false — duplicate
		f.Allow("a"), // true — not consecutive
	}
	expected := []bool{true, false, true, false, true}
	for i, got := range results {
		if got != expected[i] {
			t.Errorf("step %d: got %v, want %v", i, got, expected[i])
		}
	}
	if f.Dropped != 2 {
		t.Errorf("expected Dropped=2, got %d", f.Dropped)
	}
}

func TestFilter_ModeGlobal(t *testing.T) {
	f := dedupe.New(dedupe.ModeGlobal)
	if !f.Allow("a") {
		t.Error("first 'a' should be allowed")
	}
	if f.Allow("a") {
		t.Error("second 'a' should be dropped")
	}
	if !f.Allow("b") {
		t.Error("first 'b' should be allowed")
	}
	if f.Allow("a") {
		t.Error("third 'a' should still be dropped")
	}
	if f.Dropped != 2 {
		t.Errorf("expected Dropped=2, got %d", f.Dropped)
	}
}

func TestFilter_Reset(t *testing.T) {
	f := dedupe.New(dedupe.ModeGlobal)
	f.Allow("x")
	f.Allow("x") // dropped
	f.Reset()
	if f.Dropped != 0 {
		t.Errorf("after Reset, expected Dropped=0, got %d", f.Dropped)
	}
	if !f.Allow("x") {
		t.Error("after Reset, 'x' should be allowed again")
	}
}
