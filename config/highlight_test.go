package config_test

import (
	"testing"

	"github.com/yourorg/logslice/config"
)

func TestHighlightValidate_Disabled(t *testing.T) {
	h := config.HighlightConfig{Enabled: false}
	if err := h.Validate(); err != nil {
		t.Errorf("expected no error when disabled, got %v", err)
	}
}

func TestHighlightValidate_EnabledNoPattern(t *testing.T) {
	h := config.HighlightConfig{Enabled: true, Colour: "red"}
	if err := h.Validate(); err == nil {
		t.Error("expected error when pattern is missing")
	}
}

func TestHighlightValidate_InvalidColour(t *testing.T) {
	h := config.HighlightConfig{Enabled: true, Pattern: "ERROR", Colour: "purple"}
	if err := h.Validate(); err == nil {
		t.Error("expected error for unsupported colour")
	}
}

func TestHighlightValidate_ValidCombinations(t *testing.T) {
	colours := []string{"red", "yellow", "cyan", "RED", "Yellow"}
	for _, c := range colours {
		h := config.HighlightConfig{Enabled: true, Pattern: "WARN", Colour: c}
		if err := h.Validate(); err != nil {
			t.Errorf("unexpected error for colour %q: %v", c, err)
		}
	}
}

func TestHighlightValidate_DisabledIgnoresInvalidColour(t *testing.T) {
	h := config.HighlightConfig{Enabled: false, Pattern: "", Colour: "magenta"}
	if err := h.Validate(); err != nil {
		t.Errorf("disabled config should not validate colour, got %v", err)
	}
}
