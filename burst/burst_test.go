package burst

import (
	"testing"
	"time"
)

func TestParseMode_Valid(t *testing.T) {
	for _, s := range []string{"none", "window"} {
		if _, err := ParseMode(s); err != nil {
			t.Errorf("expected valid mode for %q, got %v", s, err)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	if _, err := ParseMode("burst"); err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresParams(t *testing.T) {
	f, err := New(ModeNone, 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Allow(time.Now()) {
		t.Error("ModeNone should always allow")
	}
}

func TestNew_InvalidThreshold(t *testing.T) {
	if _, err := New(ModeWindow, 0, time.Second); err == nil {
		t.Error("expected error for zero threshold")
	}
}

func TestNew_InvalidWindow(t *testing.T) {
	if _, err := New(ModeWindow, 5, 0); err == nil {
		t.Error("expected error for zero window")
	}
}

func TestAllow_BurstDetected(t *testing.T) {
	f, err := New(ModeWindow, 3, time.Second)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	now := time.Now()
	if f.Allow(now) {
		t.Error("should not allow before threshold")
	}
	if f.Allow(now.Add(100 * time.Millisecond)) {
		t.Error("should not allow before threshold")
	}
	if !f.Allow(now.Add(200 * time.Millisecond)) {
		t.Error("should allow at threshold")
	}
}

func TestAllow_OldLinesExpire(t *testing.T) {
	f, err := New(ModeWindow, 2, 500*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	now := time.Now()
	f.Allow(now)
	// second line is outside window relative to third
	late := now.Add(600 * time.Millisecond)
	if f.Allow(late) {
		t.Error("old line should have expired; burst should not be active")
	}
}
