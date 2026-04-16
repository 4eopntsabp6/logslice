package config_test

import (
	"testing"

	"github.com/user/logslice/config"
)

func TestDefaultTransformConfig(t *testing.T) {
	c := config.DefaultTransformConfig()
	if c.Mode != "none" {
		t.Errorf("expected default mode 'none', got %q", c.Mode)
	}
}

func TestTransformValidate_ValidModes(t *testing.T) {
	for _, m := range []string{"", "none", "upper", "lower", "trim"} {
		c := config.TransformConfig{Mode: m}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q) unexpected error: %v", m, err)
		}
	}
}

func TestTransformValidate_InvalidMode(t *testing.T) {
	c := config.TransformConfig{Mode: "reverse"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestTransformEnabled_None(t *testing.T) {
	for _, m := range []string{"", "none"} {
		c := config.TransformConfig{Mode: m}
		if c.Enabled() {
			t.Errorf("mode %q should not be enabled", m)
		}
	}
}

func TestTransformEnabled_Active(t *testing.T) {
	for _, m := range []string{"upper", "lower", "trim"} {
		c := config.TransformConfig{Mode: m}
		if !c.Enabled() {
			t.Errorf("mode %q should be enabled", m)
		}
	}
}
