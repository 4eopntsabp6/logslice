package config

import "testing"

func TestDefaultFieldFilterConfig(t *testing.T) {
	c := DefaultFieldFilterConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode none, got %q", c.Mode)
	}
}

func TestFieldFilterValidate_Disabled(t *testing.T) {
	c := DefaultFieldFilterConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFieldFilterValidate_InvalidMode(t *testing.T) {
	c := FieldFilterConfig{Mode: "fuzzy"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestFieldFilterValidate_MissingKey(t *testing.T) {
	c := FieldFilterConfig{Mode: "exact", Key: "", Value: "error"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for missing key")
	}
}

func TestFieldFilterValidate_MissingValue(t *testing.T) {
	c := FieldFilterConfig{Mode: "exact", Key: "level", Value: ""}
	if err := c.Validate(); err == nil {
		t.Error("expected error for missing value")
	}
}

func TestFieldFilterValidate_ValidCombinations(t *testing.T) {
	cases := []FieldFilterConfig{
		{Mode: "exact", Key: "level", Value: "error"},
		{Mode: "prefix", Key: "msg", Value: "conn"},
		{Mode: "regex", Key: "level", Value: "^(warn|error)$"},
	}
	for _, c := range cases {
		if err := c.Validate(); err != nil {
			t.Errorf("unexpected error for %+v: %v", c, err)
		}
	}
}

func TestFieldFilterEnabled(t *testing.T) {
	if DefaultFieldFilterConfig().Enabled() {
		t.Error("expected disabled")
	}
	c := FieldFilterConfig{Mode: "exact", Key: "level", Value: "error"}
	if !c.Enabled() {
		t.Error("expected enabled")
	}
}
