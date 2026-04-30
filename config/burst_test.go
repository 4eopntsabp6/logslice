package config

import (
	"testing"
)

func TestDefaultBurstConfig(t *testing.T) {
	c := DefaultBurstConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Threshold != 0 {
		t.Errorf("expected threshold 0, got %d", c.Threshold)
	}
	if c.Window != "" {
		t.Errorf("expected empty window, got %q", c.Window)
	}
}

func TestBurstValidate_Disabled(t *testing.T) {
	c := DefaultBurstConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBurstValidate_InvalidMode(t *testing.T) {
	c := DefaultBurstConfig()
	c.Mode = "unknown"
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestBurstValidate_LimitModeNoThreshold(t *testing.T) {
	c := DefaultBurstConfig()
	c.Mode = "limit"
	c.Window = "1s"
	c.Threshold = 0
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero threshold")
	}
}

func TestBurstValidate_LimitModeNoWindow(t *testing.T) {
	c := DefaultBurstConfig()
	c.Mode = "limit"
	c.Threshold = 5
	c.Window = ""
	if err := c.Validate(); err == nil {
		t.Error("expected error for missing window")
	}
}

func TestBurstValidate_LimitModeValid(t *testing.T) {
	c := DefaultBurstConfig()
	c.Mode = "limit"
	c.Threshold = 10
	c.Window = "5s"
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBurstEnabled_None(t *testing.T) {
	c := DefaultBurstConfig()
	if c.Enabled() {
		t.Error("expected Enabled() false for mode none")
	}
}

func TestBurstEnabled_Limit(t *testing.T) {
	c := DefaultBurstConfig()
	c.Mode = "limit"
	c.Threshold = 5
	c.Window = "1s"
	if !c.Enabled() {
		t.Error("expected Enabled() true for mode limit")
	}
}
