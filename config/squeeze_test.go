package config

import (
	"testing"
)

func TestDefaultSqueezeConfig(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	if cfg.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", cfg.Mode)
	}
}

func TestSqueezeValidate_Disabled(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for disabled config, got %v", err)
	}
}

func TestSqueezeValidate_InvalidMode(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	cfg.Mode = "badmode"
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid mode, got nil")
	}
}

func TestSqueezeValidate_ValidBlank(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	cfg.Mode = "blank"
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for blank mode, got %v", err)
	}
}

func TestSqueezeValidate_ValidWhitespace(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	cfg.Mode = "whitespace"
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for whitespace mode, got %v", err)
	}
}

func TestSqueezeEnabled_None(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	if cfg.Enabled() {
		t.Error("expected Enabled() to be false for mode 'none'")
	}
}

func TestSqueezeEnabled_Active(t *testing.T) {
	cfg := DefaultSqueezeConfig()
	cfg.Mode = "blank"
	if !cfg.Enabled() {
		t.Error("expected Enabled() to be true for mode 'blank'")
	}
}
