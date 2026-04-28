package reverse

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
		{"reverse", ModeReverse},
	}
	for _, tc := range cases {
		got, err := ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q): unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("backwards")
	if err == nil {
		t.Error("expected error for unknown mode, got nil")
	}
}

func TestFeed_ModeNone_PassThrough(t *testing.T) {
	r := New(ModeNone)
	lines := []string{"alpha", "beta", "gamma"}
	for _, l := range lines {
		out := r.Feed(l)
		if len(out) != 1 || out[0] != l {
			t.Errorf("Feed(%q) = %v, want [%q]", l, out, l)
		}
	}
	if got := r.Flush(); got != nil {
		t.Errorf("Flush (none) = %v, want nil", got)
	}
}

func TestFeed_ModeReverse_BuffersLines(t *testing.T) {
	r := New(ModeReverse)
	for _, l := range []string{"a", "b", "c"} {
		if out := r.Feed(l); out != nil {
			t.Errorf("Feed(%q) should return nil during buffering, got %v", l, out)
		}
	}
}

func TestFlush_ReversesOrder(t *testing.T) {
	r := New(ModeReverse)
	input := []string{"first", "second", "third"}
	for _, l := range input {
		r.Feed(l)
	}
	got := r.Flush()
	want := []string{"third", "second", "first"}
	if len(got) != len(want) {
		t.Fatalf("Flush len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Flush[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFlush_ResetsBuffer(t *testing.T) {
	r := New(ModeReverse)
	r.Feed("x")
	r.Flush()
	if got := r.Flush(); got != nil {
		t.Errorf("second Flush should return nil, got %v", got)
	}
}

func TestEnabled(t *testing.T) {
	if New(ModeNone).Enabled() {
		t.Error("ModeNone should not be enabled")
	}
	if !New(ModeReverse).Enabled() {
		t.Error("ModeReverse should be enabled")
	}
}
