package config

import "testing"

func TestDefaultMergeConfig(t *testing.T) {
	c := DefaultMergeConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Separator != " " {
		t.Errorf("expected separator ' ', got %q", c.Separator)
	}
}

func TestMergeValidate_Disabled(t *testing.T) {
	c := MergeConfig{Mode: "none"}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMergeValidate_InvalidMode(t *testing.T) {
	c := MergeConfig{Mode: "bogus"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestMergeValidate_ContinuationNoPattern(t *testing.T) {
	c := MergeConfig{Mode: "continuation", Pattern: ""}
	if err := c.Validate(); err == nil {
		t.Error("expected error when pattern is missing")
	}
}

func TestMergeValidate_ContinuationWithPattern(t *testing.T) {
	c := MergeConfig{Mode: "continuation", Pattern: `^\s+`, Separator: " "}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestMergeEnabled_None(t *testing.T) {
	c := MergeConfig{Mode: "none"}
	if c.Enabled() {
		t.Error("expected Enabled() == false for mode none")
	}
}

func TestMergeEnabled_Active(t *testing.T) {
	c := MergeConfig{Mode: "continuation", Pattern: `^\s+`}
	if !c.Enabled() {
		t.Error("expected Enabled() == true for active mode")
	}
}
