package config

import "testing"

func TestDefaultNormalizeConfig(t *testing.T) {
	c := DefaultNormalizeConfig()
	if c.Mode != "none" {
		t.Errorf("expected default mode 'none', got %q", c.Mode)
	}
	if err := c.Validate(); err != nil {
		t.Errorf("default config should be valid: %v", err)
	}
}

func TestNormalizeValidate_ValidModes(t *testing.T) {
	for _, mode := range []string{"none", "trim", "collapse", "all"} {
		c := NormalizeConfig{Mode: mode}
		if err := c.Validate(); err != nil {
			t.Errorf("mode %q should be valid: %v", mode, err)
		}
	}
}

func TestNormalizeValidate_InvalidMode(t *testing.T) {
	c := NormalizeConfig{Mode: "squash"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestNormalizeEnabled_None(t *testing.T) {
	c := NormalizeConfig{Mode: "none"}
	if c.Enabled() {
		t.Error("mode 'none' should not be enabled")
	}
}

func TestNormalizeEnabled_Active(t *testing.T) {
	for _, mode := range []string{"trim", "collapse", "all"} {
		c := NormalizeConfig{Mode: mode}
		if !c.Enabled() {
			t.Errorf("mode %q should be enabled", mode)
		}
	}
}
