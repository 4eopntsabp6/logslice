package since_test

import (
	"testing"
	"time"

	"github.com/yourorg/logslice/since"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  time.Duration
	}{
		{"", 0},
		{"1h", time.Hour},
		{"30m", 30 * time.Minute},
		{"2h30m", 2*time.Hour + 30*time.Minute},
	}
	for _, c := range cases {
		got, err := since.ParseMode(c.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	cases := []string{"notaduration", "-1h", "0s"}
	for _, c := range cases {
		_, err := since.ParseMode(c)
		if err == nil {
			t.Errorf("ParseMode(%q) expected error, got nil", c)
		}
	}
}

func TestNew_Disabled(t *testing.T) {
	f := since.New(0, time.Now())
	if f.Enabled() {
		t.Fatal("expected disabled filter")
	}
	if !f.Keep(time.Now().Add(-24 * time.Hour)) {
		t.Fatal("disabled filter should keep all lines")
	}
}

func TestKeep_WithinWindow(t *testing.T) {
	now := time.Now()
	f := since.New(time.Hour, now)
	recent := now.Add(-30 * time.Minute)
	if !f.Keep(recent) {
		t.Error("expected recent line to be kept")
	}
}

func TestKeep_OutsideWindow(t *testing.T) {
	now := time.Now()
	f := since.New(time.Hour, now)
	old := now.Add(-2 * time.Hour)
	if f.Keep(old) {
		t.Error("expected old line to be dropped")
	}
}

func TestKeep_ZeroTimestamp(t *testing.T) {
	f := since.New(time.Hour, time.Now())
	if !f.Keep(time.Time{}) {
		t.Error("zero timestamp should always be kept")
	}
}
