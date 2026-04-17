package ratelimit_test

import (
	"testing"

	"github.com/user/logslice/ratelimit"
)

func TestParseMode_Valid(t *testing.T) {
	for _, tc := range []string{"none", "lines"} {
		if _, err := ratelimit.ParseMode(tc); err != nil {
			t.Errorf("expected valid mode for %q, got %v", tc, err)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	if _, err := ratelimit.ParseMode("burst"); err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneIgnoresLimit(t *testing.T) {
	l, err := ratelimit.New(ratelimit.ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 0; i < 100; i++ {
		if !l.Allow() {
			t.Fatal("ModeNone should allow all lines")
		}
	}
}

func TestNew_InvalidLimit(t *testing.T) {
	if _, err := ratelimit.New(ratelimit.ModeLines, 0); err == nil {
		t.Error("expected error for limit=0")
	}
	if _, err := ratelimit.New(ratelimit.ModeLines, -1); err == nil {
		t.Error("expected error for negative limit")
	}
}

func TestAllow_RespectsLimit(t *testing.T) {
	l, err := ratelimit.New(ratelimit.ModeLines, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	allowed := 0
	for i := 0; i < 10; i++ {
		if l.Allow() {
			allowed++
		}
	}
	if allowed != 3 {
		t.Errorf("expected 3 allowed, got %d", allowed)
	}
}
