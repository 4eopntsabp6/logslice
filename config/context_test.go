package config

import "testing"

func TestDefaultContextConfig(t *testing.T) {
	c := DefaultContextConfig()
	if c.Enabled {
		t.Fatal("expected disabled by default")
	}
	if c.Mode != "none" {
		t.Fatalf("expected mode 'none', got %q", c.Mode)
	}
}

func TestContextValidate_Disabled(t *testing.T) {
	c := DefaultContextConfig()
	if err := c.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestContextValidate_InvalidMode(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "sideways", Before: 1, After: 1}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func TestContextValidate_BeforeZero(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "before", Before: 0}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error when before=0 with mode 'before'")
	}
}

func TestContextValidate_AfterZero(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "after", After: 0}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error when after=0 with mode 'after'")
	}
}

func TestContextValidate_ValidBoth(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "both", Before: 2, After: 2}
	if err := c.Validate(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestContextEnabled_None(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "none"}
	if c.ContextEnabled() {
		t.Fatal("expected false for mode 'none'")
	}
}

func TestContextEnabled_Active(t *testing.T) {
	c := ContextConfig{Enabled: true, Mode: "both", Before: 1, After: 1}
	if !c.ContextEnabled() {
		t.Fatal("expected true")
	}
}
