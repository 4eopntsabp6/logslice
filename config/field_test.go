package config

import (
	"testing"
)

func TestDefaultFieldConfig(t *testing.T) {
	cfg := DefaultFieldConfig()
	if cfg.Mode != "none" {
		t.Fatalf("expected mode none, got %q", cfg.Mode)
	}
	if cfg.Enabled() {
		t.Fatal("expected Enabled() false for default config")
	}
}

func TestFieldValidate_Disabled(t *testing.T) {
	cfg := DefaultFieldConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFieldValidate_InvalidMode(t *testing.T) {
	cfg := FieldConfig{Mode: "xml", Key: "id"}
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func TestFieldValidate_EnabledNoKey(t *testing.T) {
	for _, mode := range []string{"json", "kv"} {
		cfg := FieldConfig{Mode: mode}
		if err := cfg.Validate(); err == nil {
			t.Fatalf("expected error for mode %q with no key", mode)
		}
	}
}

func TestFieldValidate_ValidCombinations(t *testing.T) {
	for _, mode := range []string{"json", "kv"} {
		cfg := FieldConfig{Mode: mode, Key: "level"}
		if err := cfg.Validate(); err != nil {
			t.Fatalf("unexpected error for mode %q: %v", mode, err)
		}
		if !cfg.Enabled() {
			t.Fatalf("expected Enabled() true for mode %q", mode)
		}
	}
}

func TestFieldEnabled_None(t *testing.T) {
	cfg := FieldConfig{Mode: "none"}
	if cfg.Enabled() {
		t.Fatal("expected Enabled() false for mode none")
	}
}
