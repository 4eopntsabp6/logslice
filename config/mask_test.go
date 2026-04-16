package config

import "testing"

func TestDefaultMaskConfig(t *testing.T) {
	c := DefaultMaskConfig()
	if c.Enabled {
		t.Error("expected disabled by default")
	}
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
}

func TestMaskValidate_Disabled(t *testing.T) {
	c := DefaultMaskConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMaskValidate_EnabledNoMode(t *testing.T) {
	c := MaskConfig{Enabled: true, Mode: "none", Pattern: `\d+`}
	if err := c.Validate(); err == nil {
		t.Error("expected error for mode=none when enabled")
	}
}

func TestMaskValidate_EnabledNoPattern(t *testing.T) {
	c := MaskConfig{Enabled: true, Mode: "redact", Pattern: ""}
	if err := c.Validate(); err == nil {
		t.Error("expected error for missing pattern")
	}
}

func TestMaskValidate_InvalidMode(t *testing.T) {
	c := MaskConfig{Enabled: true, Mode: "scramble", Pattern: `\d+`}
	if err := c.Validate(); err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestMaskValidate_ValidCombinations(t *testing.T) {
	for _, mode := range []string{"redact", "partial"} {
		c := MaskConfig{Enabled: true, Mode: mode, Pattern: `\d+`}
		if err := c.Validate(); err != nil {
			t.Errorf("mode %q: unexpected error: %v", mode, err)
		}
	}
}

func TestMaskEnabled(t *testing.T) {
	c := MaskConfig{Enabled: true, Mode: "redact", Pattern: `\d+`}
	if !c.MaskEnabled() {
		t.Error("expected MaskEnabled true")
	}
	c2 := DefaultMaskConfig()
	if c2.MaskEnabled() {
		t.Error("expected MaskEnabled false for default")
	}
}
