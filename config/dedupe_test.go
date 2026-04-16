package config

import "testing"

func TestDefaultDedupeConfig(t *testing.T) {
	cfg := DefaultDedupeConfig()
	if cfg.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", cfg.Mode)
	}
	if cfg.Field != "" {
		t.Errorf("expected empty field, got %q", cfg.Field)
	}
}

func TestDedupeValidate_ValidModes(t *testing.T) {
	for _, mode := range []string{"none", "consecutive", "global"} {
		cfg := DedupeConfig{Mode: mode}
		if err := cfg.Validate(); err != nil {
			t.Errorf("mode %q: unexpected error: %v", mode, err)
		}
	}
}

func TestDedupeValidate_InvalidMode(t *testing.T) {
	cfg := DedupeConfig{Mode: "fuzzy"}
	if err := cfg.Validate(); err == nil {
		t.Error("expected error for invalid mode, got nil")
	}
}

func TestDedupeEnabled_None(t *testing.T) {
	cfg := DedupeConfig{Mode: "none"}
	if cfg.Enabled() {
		t.Error("expected Enabled() false for mode 'none'")
	}
}

func TestDedupeEnabled_Consecutive(t *testing.T) {
	cfg := DedupeConfig{Mode: "consecutive"}
	if !cfg.Enabled() {
		t.Error("expected Enabled() true for mode 'consecutive'")
	}
}

func TestDedupeEnabled_Global(t *testing.T) {
	cfg := DedupeConfig{Mode: "global"}
	if !cfg.Enabled() {
		t.Error("expected Enabled() true for mode 'global'")
	}
}
