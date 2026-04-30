package config

import (
	"testing"
)

func TestDefaultStripConfig(t *testing.T) {
	c := DefaultStripConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
}

func TestStripValidate_Disabled(t *testing.T) {
	c := DefaultStripConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStripValidate_InvalidMode(t *testing.T) {
	c := DefaultStripConfig()
	c.Mode = "bogus"
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestStripValidate_ValidANSI(t *testing.T) {
	c := DefaultStripConfig()
	c.Mode = "ansi"
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStripValidate_ValidWhitespace(t *testing.T) {
	c := DefaultStripConfig()
	c.Mode = "whitespace"
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStripValidate_ValidBoth(t *testing.T) {
	c := DefaultStripConfig()
	c.Mode = "both"
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestStripEnabled_None(t *testing.T) {
	c := DefaultStripConfig()
	if c.Enabled() {
		t.Error("expected Enabled() false for mode none")
	}
}

func TestStripEnabled_ANSI(t *testing.T) {
	c := DefaultStripConfig()
	c.Mode = "ansi"
	if !c.Enabled() {
		t.Error("expected Enabled() true for mode ansi")
	}
}
