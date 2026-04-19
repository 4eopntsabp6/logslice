package config

import "testing"

func TestDefaultBookmarkConfig(t *testing.T) {
	c := DefaultBookmarkConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Tag == "" {
		t.Error("expected non-empty default tag")
	}
}

func TestBookmarkValidate_Disabled(t *testing.T) {
	c := DefaultBookmarkConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBookmarkValidate_InvalidMode(t *testing.T) {
	c := BookmarkConfig{Mode: "bad"}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func TestBookmarkValidate_IndexNoIndices(t *testing.T) {
	c := BookmarkConfig{Mode: "index", Tag: ">>>"}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for missing indices")
	}
}

func TestBookmarkValidate_RegexNoPattern(t *testing.T) {
	c := BookmarkConfig{Mode: "regex", Tag: ">>>"}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for missing pattern")
	}
}

func TestBookmarkValidate_ValidIndex(t *testing.T) {
	c := BookmarkConfig{Mode: "index", Tag: "[*] ", Indices: []int{1, 3}}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBookmarkValidate_ValidRegex(t *testing.T) {
	c := BookmarkConfig{Mode: "regex", Tag: "[!] ", Pattern: "ERROR"}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestBookmarkEnabled(t *testing.T) {
	if DefaultBookmarkConfig().Enabled() {
		t.Error("default should not be enabled")
	}
	c := BookmarkConfig{Mode: "regex", Tag: ">>>", Pattern: "ERR"}
	if !c.Enabled() {
		t.Error("expected enabled")
	}
}
