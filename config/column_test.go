package config_test

import (
	"testing"

	"github.com/yourorg/logslice/config"
)

func TestDefaultColumnConfig(t *testing.T) {
	c := config.DefaultColumnConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.Delimiter != "," {
		t.Errorf("expected delimiter ',', got %q", c.Delimiter)
	}
}

func TestColumnValidate_Disabled(t *testing.T) {
	c := config.DefaultColumnConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestColumnValidate_InvalidMode(t *testing.T) {
	c := config.DefaultColumnConfig()
	c.Mode = "split"
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestColumnValidate_ExtractNoDelimiter(t *testing.T) {
	c := config.ColumnConfig{Mode: "extract", Delimiter: "", Index: 0}
	if err := c.Validate(); err == nil {
		t.Error("expected error for missing delimiter")
	}
}

func TestColumnValidate_ExtractNegativeIndex(t *testing.T) {
	c := config.ColumnConfig{Mode: "extract", Delimiter: ",", Index: -1}
	if err := c.Validate(); err == nil {
		t.Error("expected error for negative index")
	}
}

func TestColumnValidate_ValidExtract(t *testing.T) {
	c := config.ColumnConfig{Mode: "extract", Delimiter: "|", Index: 2}
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestColumnEnabled(t *testing.T) {
	none := config.DefaultColumnConfig()
	if none.Enabled() {
		t.Error("expected Enabled() false for mode none")
	}
	active := config.ColumnConfig{Mode: "extract", Delimiter: ",", Index: 0}
	if !active.Enabled() {
		t.Error("expected Enabled() true for mode extract")
	}
}
