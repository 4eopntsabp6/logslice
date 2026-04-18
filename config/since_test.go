package config_test

import (
	"testing"

	"github.com/yourorg/logslice/config"
)

func TestDefaultSinceConfig(t *testing.T) {
	c := config.DefaultSinceConfig()
	if c.Duration != "" {
		t.Errorf("expected empty duration, got %q", c.Duration)
	}
	if c.Enabled() {
		t.Error("default config should be disabled")
	}
}

func TestSinceValidate_Empty(t *testing.T) {
	c := config.SinceConfig{Duration: ""}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSinceValidate_ValidDurations(t *testing.T) {
	for _, d := range []string{"1h", "30m", "2h15m", "90s"} {
		c := config.SinceConfig{Duration: d}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q) unexpected error: %v", d, err)
		}
		if !c.Enabled() {
			t.Errorf("Enabled(%q) expected true", d)
		}
	}
}

func TestSinceValidate_InvalidDuration(t *testing.T) {
	c := config.SinceConfig{Duration: "notvalid"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid duration")
	}
}

func TestSinceValidate_NegativeDuration(t *testing.T) {
	c := config.SinceConfig{Duration: "-1h"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for negative duration")
	}
}

func TestSinceValidate_ZeroDuration(t *testing.T) {
	c := config.SinceConfig{Duration: "0s"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero duration")
	}
}
