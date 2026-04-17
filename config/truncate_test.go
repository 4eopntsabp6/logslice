package config_test

import (
	"testing"

	"github.com/yourorg/logslice/config"
)

func TestDefaultTruncateConfig(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	if cfg.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", cfg.Mode)
	}
	if cfg.Limit != 0 {
		t.Errorf("expected limit 0, got %d", cfg.Limit)
	}
}

func TestTruncateValidate_Disabled(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTruncateValidate_InvalidMode(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	cfg.Mode = "bad"
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestTruncateValidate_ZeroLimit(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	cfg.Mode = "chars"
	cfg.Limit = 0
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for zero limit with active mode")
	}
}

func TestTruncateValidate_ValidChars(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	cfg.Mode = "chars"
	cfg.Limit = 80
	if err := cfg.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTruncateEnabled_None(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	if cfg.Enabled() {
		t.Error("expected Enabled() false for mode none")
	}
}

func TestTruncateEnabled_Active(t *testing.T) {
	cfg := config.DefaultTruncateConfig()
	cfg.Mode = "words"
	cfg.Limit = 10
	if !cfg.Enabled() {
		t.Error("expected Enabled() true for active mode")
	}
}
