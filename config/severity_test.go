package config

import (
	"testing"
)

func TestDefaultSeverityConfig(t *testing.T) {
	c := DefaultSeverityConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Level != "info" {
		t.Errorf("expected level 'info', got %q", c.Level)
	}
}

func TestSeverityValidate_ValidModes(t *testing.T) {
	for _, tc := range []struct{ mode, level string }{
		{"none", ""},
		{"min", "warn"},
		{"exact", "error"},
	} {
		c := SeverityConfig{Mode: tc.mode, Level: tc.level}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q,%q) unexpected error: %v", tc.mode, tc.level, err)
		}
	}
}

func TestSeverityValidate_InvalidMode(t *testing.T) {
	c := SeverityConfig{Mode: "verbose", Level: "info"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestSeverityValidate_InvalidLevel(t *testing.T) {
	c := SeverityConfig{Mode: "min", Level: "trace"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid level")
	}
}

func TestSeverityValidate_NoneIgnoresLevel(t *testing.T) {
	c := SeverityConfig{Mode: "none", Level: "trace"}
	if err := c.Validate(); err != nil {
		t.Errorf("none mode should ignore invalid level, got: %v", err)
	}
}
