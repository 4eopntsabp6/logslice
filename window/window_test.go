package window_test

import (
	"testing"
	"time"

	"github.com/robinovitch61/logslice/window"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  window.Mode
	}{
		{"", window.ModeNone},
		{"none", window.ModeNone},
		{"sliding", window.ModeSliding},
	}
	for _, tc := range cases {
		got, err := window.ParseMode(tc.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", tc.input, err)
		}
		if got != tc.want {
			t.Errorf("ParseMode(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := window.ParseMode("rolling")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresDuration(t *testing.T) {
	f, err := window.New(window.ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f == nil {
		t.Fatal("expected non-nil filter")
	}
}

func TestNew_SlidingRequiresPositiveDuration(t *testing.T) {
	_, err := window.New(window.ModeSliding, 0)
	if err == nil {
		t.Error("expected error for zero duration")
	}
	_, err = window.New(window.ModeSliding, -time.Second)
	if err == nil {
		t.Error("expected error for negative duration")
	}
}

func TestKeep_ModeNone_AllowsAll(t *testing.T) {
	f, _ := window.New(window.ModeNone, 0)
	for _, line := range []string{"anything", "", "2024-01-01T00:00:00Z info msg"} {
		if !f.Keep(line) {
			t.Errorf("ModeNone should keep all lines, dropped: %q", line)
		}
	}
}

func TestKeep_SlidingWindow(t *testing.T) {
	f, _ := window.New(window.ModeSliding, 5*time.Second)

	// First timestamped line sets the anchor.
	if !f.Keep("2024-01-01T00:00:00Z first") {
		t.Error("anchor line should be kept")
	}
	if !f.Keep("2024-01-01T00:00:04Z within") {
		t.Error("line within window should be kept")
	}
	if f.Keep("2024-01-01T00:00:05Z boundary") {
		t.Error("line at boundary (exclusive) should be dropped")
	}
	if f.Keep("2024-01-01T00:00:10Z outside") {
		t.Error("line outside window should be dropped")
	}
}

func TestKeep_Reset(t *testing.T) {
	f, _ := window.New(window.ModeSliding, 2*time.Second)
	f.Keep("2024-01-01T00:00:00Z anchor")
	f.Reset()
	// After reset the next timestamped line becomes the new anchor.
	if !f.Keep("2024-01-01T00:01:00Z new-anchor") {
		t.Error("post-reset anchor should be kept")
	}
	if !f.Keep("2024-01-01T00:01:01Z still-in") {
		t.Error("line within new window should be kept")
	}
}
