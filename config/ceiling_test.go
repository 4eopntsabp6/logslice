package config

import (
	"testing"
)

func TestDefaultCeilingConfig(t *testing.T) {
	c := DefaultCeilingConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Max != 0 {
		t.Errorf("expected max 0, got %d", c.Max)
	}
	if c.Pattern != "" {
		t.Errorf("expected empty pattern, got %q", c.Pattern)
	}
}

func TestCeilingValidate_Disabled(t *testing.T) {
	c := DefaultCeilingConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCeilingValidate_InvalidMode(t *testing.T) {
	c := DefaultCeilingConfig()
	c.Mode = "bogus"
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestCeilingValidate_InvalidMax(t *testing.T) {
	c := DefaultCeilingConfig()
	c.Mode = "cap"
	c.Max = 0
	if err := c.Validate(); err == nil {
		t.Error("expected error for zero max")
	}
}

func TestCeilingValidate_ValidCap(t *testing.T) {
	c := DefaultCeilingConfig()
	c.Mode = "cap"
	c.Max = 100
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCeilingValidate_ValidCapWithPattern(t *testing.T) {
	c := DefaultCeilingConfig()
	c.Mode = "cap"
	c.Max = 50
	c.Pattern = `ERROR`
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestCeilingEnabled_None(t *testing.T) {
	c := DefaultCeilingConfig()
	if c.Enabled() {
		t.Error("expected Enabled() false for mode none")
	}
}

func TestCeilingEnabled_Cap(t *testing.T) {
	c := DefaultCeilingConfig()
	c.Mode = "cap"
	c.Max = 10
	if !c.Enabled() {
		t.Error("expected Enabled() true for mode cap")
	}
}
