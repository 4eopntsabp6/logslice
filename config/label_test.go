package config

import "testing"

func TestDefaultLabelConfig(t *testing.T) {
	c := DefaultLabelConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
}

func TestLabelValidate_Disabled(t *testing.T) {
	c := DefaultLabelConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLabelValidate_InvalidMode(t *testing.T) {
	c := LabelConfig{Mode: "wrap"}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestLabelValidate_EnabledNoLabel(t *testing.T) {
	c := LabelConfig{Mode: "prefix", Label: "", Pattern: "error"}
	if err := c.Validate(); err == nil {
		t.Error("expected error when label text is empty")
	}
}

func TestLabelValidate_EnabledNoPattern(t *testing.T) {
	c := LabelConfig{Mode: "append", Label: "TAG", Pattern: ""}
	if err := c.Validate(); err == nil {
		t.Error("expected error when pattern is empty")
	}
}

func TestLabelValidate_Valid(t *testing.T) {
	cases := []LabelConfig{
		{Mode: "prefix", Label: "ERR", Pattern: "error"},
		{Mode: "append", Label: "WARN", Pattern: "warn"},
	}
	for _, c := range cases {
		if err := c.Validate(); err != nil {
			t.Errorf("unexpected error for %+v: %v", c, err)
		}
	}
}

func TestLabelEnabled_None(t *testing.T) {
	c := DefaultLabelConfig()
	if c.Enabled() {
		t.Error("expected Enabled() = false for mode none")
	}
}

func TestLabelEnabled_Active(t *testing.T) {
	c := LabelConfig{Mode: "prefix", Label: "X", Pattern: "x"}
	if !c.Enabled() {
		t.Error("expected Enabled() = true")
	}
}
