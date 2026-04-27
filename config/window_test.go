package config_test

import (
	"testing"
	"time"

	"github.com/robinovitch61/logslice/config"
)

func TestDefaultWindowConfig(t *testing.T) {
	c := config.DefaultWindowConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Duration != 0 {
		t.Errorf("expected zero duration, got %v", c.Duration)
	}
}

func TestWindowValidate_Disabled(t *testing.T) {
	c := config.DefaultWindowConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWindowValidate_InvalidMode(t *testing.T) {
	c := config.WindowConfig{Mode: "tumbling", Duration: time.Second}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestWindowValidate_SlidingZeroDuration(t *testing.T) {
	c := config.WindowConfig{Mode: "sliding", Duration: 0}
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero duration in sliding mode")
	}
}

func TestWindowValidate_SlidingValidDuration(t *testing.T) {
	c := config.WindowConfig{Mode: "sliding", Duration: 10 * time.Second}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestWindowEnabled_None(t *testing.T) {
	c := config.DefaultWindowConfig()
	if c.Enabled() {
		t.Error("expected Enabled() == false for mode 'none'")
	}
}

func TestWindowEnabled_Sliding(t *testing.T) {
	c := config.WindowConfig{Mode: "sliding", Duration: 5 * time.Second}
	if !c.Enabled() {
		t.Error("expected Enabled() == true for mode 'sliding'")
	}
}
