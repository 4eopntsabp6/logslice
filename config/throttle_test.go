package config_test

import (
	"testing"
	"time"

	"github.com/user/logslice/config"
)

func TestDefaultThrottleConfig(t *testing.T) {
	c := config.DefaultThrottleConfig()
	if c.Mode != "none" {
		t.Errorf("default mode = %q, want none", c.Mode)
	}
	if c.Delay != 0 {
		t.Errorf("default delay = %v, want 0", c.Delay)
	}
}

func TestThrottleValidate_ValidModes(t *testing.T) {
	cases := []struct {
		mode  string
		delay time.Duration
	}{
		{"none", 0},
		{"", 0},
		{"delay", 50 * time.Millisecond},
	}
	for _, c := range cases {
		cfg := config.ThrottleConfig{Mode: c.mode, Delay: c.delay}
		if err := cfg.Validate(); err != nil {
			t.Errorf("Validate(%q) unexpected error: %v", c.mode, err)
		}
	}
}

func TestThrottleValidate_InvalidMode(t *testing.T) {
	cfg := config.ThrottleConfig{Mode: "burst"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestThrottleValidate_DelayModeZeroDuration(t *testing.T) {
	cfg := config.ThrottleConfig{Mode: "delay", Delay: 0}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for delay mode with zero duration")
	}
}

func TestThrottleEnabled(t *testing.T) {
	if (config.ThrottleConfig{Mode: "none"}).Enabled() {
		t.Error("none should not be enabled")
	}
	if (config.ThrottleConfig{Mode: ""}).Enabled() {
		t.Error("empty should not be enabled")
	}
	if !(config.ThrottleConfig{Mode: "delay", Delay: time.Millisecond}).Enabled() {
		t.Error("delay should be enabled")
	}
}
