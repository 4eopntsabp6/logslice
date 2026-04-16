package config

import (
	"testing"
)

func TestDefaultSampleConfig(t *testing.T) {
	cfg := DefaultSampleConfig()
	if cfg.Mode != "none" {
		t.Fatalf("expected mode none, got %q", cfg.Mode)
	}
}

func TestSampleValidate_ValidModes(t *testing.T) {
	for _, mode := range []string{"none", "nth", "random"} {
		cfg := SampleConfig{Mode: mode, N: 2}
		if err := cfg.Validate(); err != nil {
			t.Fatalf("unexpected error for mode %q: %v", mode, err)
		}
	}
}

func TestSampleValidate_InvalidMode(t *testing.T) {
	cfg := SampleConfig{Mode: "burst", N: 1}
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func TestSampleValidate_ZeroN(t *testing.T) {
	cfg := SampleConfig{Mode: "nth", N: 0}
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error when N is zero for active mode")
	}
}

func TestSampleEnabled_None(t *testing.T) {
	cfg := SampleConfig{Mode: "none"}
	if cfg.Enabled() {
		t.Fatal("expected Enabled() false for mode none")
	}
}

func TestSampleEnabled_Active(t *testing.T) {
	cfg := SampleConfig{Mode: "nth", N: 3}
	if !cfg.Enabled() {
		t.Fatal("expected Enabled() true for mode nth")
	}
}
