package throttle_test

import (
	"testing"
	"time"

	"github.com/user/logslice/throttle"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  throttle.Mode
	}{
		{"", throttle.ModeNone},
		{"none", throttle.ModeNone},
		{"delay", throttle.ModeDelay},
	}
	for _, c := range cases {
		got, err := throttle.ParseMode(c.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := throttle.ParseMode("burst")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresDelay(t *testing.T) {
	_, err := throttle.New(throttle.ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNew_InvalidDelay(t *testing.T) {
	_, err := throttle.New(throttle.ModeDelay, 0)
	if err == nil {
		t.Error("expected error for zero delay")
	}
}

func TestNew_NegativeDelay(t *testing.T) {
	_, err := throttle.New(throttle.ModeDelay, -1*time.Millisecond)
	if err == nil {
		t.Error("expected error for negative delay")
	}
}

func TestEnabled(t *testing.T) {
	none, _ := throttle.New(throttle.ModeNone, 0)
	if none.Enabled() {
		t.Error("ModeNone should not be enabled")
	}
	d, _ := throttle.New(throttle.ModeDelay, 10*time.Millisecond)
	if !d.Enabled() {
		t.Error("ModeDelay should be enabled")
	}
}

func TestWait_ModeNone_DoesNotBlock(t *testing.T) {
	th, _ := throttle.New(throttle.ModeNone, 0)
	start := time.Now()
	th.Wait()
	if time.Since(start) > 10*time.Millisecond {
		t.Error("ModeNone Wait should return immediately")
	}
}
