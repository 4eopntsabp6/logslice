package config

import (
	"testing"
)

func TestDefaultCountConfig(t *testing.T) {
	cfg := DefaultCountConfig()
	if cfg.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", cfg.Mode)
	}
	if cfg.N != 0 {
		t.Errorf("expected N 0, got %d", cfg.N)
	}
}

func TestCountValidate_Disabled(t *testing.T) {
	cfg := DefaultCountConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for disabled config, got %v", err)
	}
}

func TestCountValidate_InvalidMode(t *testing.T) {
	cfg := DefaultCountConfig()
	cfg.Mode = "badmode"
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid mode, got nil")
	}
}

func TestCountValidate_ZeroN(t *testing.T) {
	cfg := DefaultCountConfig()
	cfg.Mode = "first"
	cfg.N = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero N with active mode, got nil")
	}
}

func TestCountValidate_ValidFirst(t *testing.T) {
	cfg := DefaultCountConfig()
	cfg.Mode = "first"
	cfg.N = 10
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCountValidate_ValidLast(t *testing.T) {
	cfg := DefaultCountConfig()
	cfg.Mode = "last"
	cfg.N = 5
	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestCountEnabled_None(t *testing.T) {
	cfg := DefaultCountConfig()
	if cfg.Enabled() {
		t.Error("expected Enabled() false for mode 'none'")
	}
}

func TestCountEnabled_Active(t *testing.T) {
	cfg := DefaultCountConfig()
	cfg.Mode = "first"
	cfg.N = 1
	if !cfg.Enabled() {
		t.Error("expected Enabled() true for active mode")
	}
}
