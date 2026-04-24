package config

import "testing"

func TestDefaultEntropyConfig(t *testing.T) {
	c := DefaultEntropyConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Threshold != 0 {
		t.Errorf("expected threshold 0, got %f", c.Threshold)
	}
}

func TestEntropyValidate_ValidModes(t *testing.T) {
	for _, m := range []string{"none", "above", "below", ""} {
		c := EntropyConfig{Mode: m, Threshold: 2.5}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q) unexpected error: %v", m, err)
		}
	}
}

func TestEntropyValidate_InvalidMode(t *testing.T) {
	c := EntropyConfig{Mode: "sideways", Threshold: 1.0}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestEntropyValidate_NegativeThreshold(t *testing.T) {
	c := EntropyConfig{Mode: "above", Threshold: -1}
	if err := c.Validate(); err == nil {
		t.Error("expected error for negative threshold with active mode")
	}
}

func TestEntropyValidate_NoneIgnoresNegative(t *testing.T) {
	c := EntropyConfig{Mode: "none", Threshold: -5}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error for none mode: %v", err)
	}
}

func TestEntropyEnabled(t *testing.T) {
	if DefaultEntropyConfig().Enabled() {
		t.Error("default config should not be enabled")
	}
	c := EntropyConfig{Mode: "above", Threshold: 3.0}
	if !c.Enabled() {
		t.Error("above mode should be enabled")
	}
}
