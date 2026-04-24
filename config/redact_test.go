package config_test

import (
	"testing"

	"github.com/yourorg/logslice/config"
)

func TestDefaultRedactConfig(t *testing.T) {
	c := config.DefaultRedactConfig()
	if c.Mode != "none" {
		t.Errorf("default Mode = %q, want \"none\"", c.Mode)
	}
	if c.Placeholder != "[REDACTED]" {
		t.Errorf("default Placeholder = %q, want \"[REDACTED]\"", c.Placeholder)
	}
}

func TestRedactValidate_Disabled(t *testing.T) {
	cases := []string{"", "none"}
	for _, m := range cases {
		c := config.RedactConfig{Mode: m}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q) unexpected error: %v", m, err)
		}
	}
}

func TestRedactValidate_InvalidMode(t *testing.T) {
	c := config.RedactConfig{Mode: "scrub"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestRedactValidate_PatternModeNoPattern(t *testing.T) {
	c := config.RedactConfig{Mode: "pattern", Pattern: ""}
	if err := c.Validate(); err == nil {
		t.Error("expected error when pattern is empty")
	}
}

func TestRedactValidate_ValidPattern(t *testing.T) {
	c := config.RedactConfig{Mode: "pattern", Pattern: `password=\S+`}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRedactEnabled_None(t *testing.T) {
	c := config.RedactConfig{Mode: "none"}
	if c.Enabled() {
		t.Error("expected Enabled() == false for mode none")
	}
}

func TestRedactEnabled_Pattern(t *testing.T) {
	c := config.RedactConfig{Mode: "pattern", Pattern: `token=\S+`}
	if !c.Enabled() {
		t.Error("expected Enabled() == true for mode pattern")
	}
}
