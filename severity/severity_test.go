package severity

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	for _, tc := range []struct{ in string; want Mode }{
		{"none", ModeNone},
		{"min", ModeMin},
		{"exact", ModeExact},
	} {
		got, err := ParseMode(tc.in)
		if err != nil || got != tc.want {
			t.Errorf("ParseMode(%q) = %v, %v", tc.in, got, err)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("verbose")
	if err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestParseLevel_Valid(t *testing.T) {
	for _, tc := range []struct{ in string; want Level }{
		{"debug", LevelDebug},
		{"INFO", LevelInfo},
		{"warn", LevelWarn},
		{"error", LevelError},
		{"fatal", LevelFatal},
	} {
		got, err := ParseLevel(tc.in)
		if err != nil || got != tc.want {
			t.Errorf("ParseLevel(%q) = %v, %v", tc.in, got, err)
		}
	}
}

func TestParseLevel_Invalid(t *testing.T) {
	_, err := ParseLevel("trace")
	if err == nil {
		t.Error("expected error for unknown level")
	}
}

func TestNew_ModeNone(t *testing.T) {
	f, err := New(ModeNone, "")
	if err != nil || f == nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !f.Allow("anything goes here") {
		t.Error("ModeNone should allow all lines")
	}
}

func TestAllow_ModeMin(t *testing.T) {
	f, _ := New(ModeMin, "warn")
	if f.Allow("DEBUG something") {
		t.Error("debug should be filtered below warn")
	}
	if f.Allow("INFO something") {
		t.Error("info should be filtered below warn")
	}
	if !f.Allow("WARN something") {
		t.Error("warn should pass")
	}
	if !f.Allow("ERROR something") {
		t.Error("error should pass")
	}
	if !f.Allow("FATAL something") {
		t.Error("fatal should pass")
	}
}

func TestAllow_ModeExact(t *testing.T) {
	f, _ := New(ModeExact, "error")
	if !f.Allow("ERROR: disk full") {
		t.Error("error line should pass exact filter")
	}
	if f.Allow("WARN: low memory") {
		t.Error("warn should not pass exact error filter")
	}
}

func TestNew_InvalidLevel(t *testing.T) {
	_, err := New(ModeMin, "trace")
	if err == nil {
		t.Error("expected error for invalid level")
	}
}
