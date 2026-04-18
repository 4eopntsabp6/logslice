package config

import "testing"

func TestDefaultLineNumConfig(t *testing.T) {
	c := DefaultLineNumConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode=none, got %q", c.Mode)
	}
	if c.Enabled() {
		t.Error("expected Enabled() = false")
	}
}

func TestLineNumValidate_Disabled(t *testing.T) {
	c := DefaultLineNumConfig()
	if err := c.Validate(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestLineNumValidate_InvalidMode(t *testing.T) {
	c := LineNumConfig{Mode: "bogus"}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid mode")
	}
}

func TestLineNumValidate_RangeFromZero(t *testing.T) {
	c := LineNumConfig{Mode: "range", From: 0}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for from=0")
	}
}

func TestLineNumValidate_RangeToBeforeFrom(t *testing.T) {
	c := LineNumConfig{Mode: "range", From: 10, To: 5}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for to < from")
	}
}

func TestLineNumValidate_ListEmpty(t *testing.T) {
	c := LineNumConfig{Mode: "list", Lines: []int{}}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for empty list")
	}
}

func TestLineNumValidate_ValidCombinations(t *testing.T) {
	cases := []LineNumConfig{
		{Mode: "range", From: 1, To: 10},
		{Mode: "range", From: 5},
		{Mode: "list", Lines: []int{1, 3, 5}},
	}
	for _, c := range cases {
		if err := c.Validate(); err != nil {
			t.Errorf("unexpected error for %+v: %v", c, err)
		}
		if !c.Enabled() {
			t.Errorf("expected Enabled()=true for %+v", c)
		}
	}
}
