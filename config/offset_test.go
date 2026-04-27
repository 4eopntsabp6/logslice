package config

import "testing"

func TestDefaultOffsetConfig(t *testing.T) {
	c := DefaultOffsetConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode 'none', got %q", c.Mode)
	}
	if c.N != 0 {
		t.Errorf("expected N=0, got %d", c.N)
	}
}

func TestOffsetValidate_ValidModes(t *testing.T) {
	cases := []struct {
		mode string
		n    int
	}{
		{"none", 0},
		{"skip", 5},
		{"start", 10},
	}
	for _, tc := range cases {
		c := OffsetConfig{Mode: tc.mode, N: tc.n}
		if err := c.Validate(); err != nil {
			t.Errorf("Validate(%q, %d): unexpected error: %v", tc.mode, tc.n, err)
		}
	}
}

func TestOffsetValidate_InvalidMode(t *testing.T) {
	c := OffsetConfig{Mode: "bad", N: 1}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestOffsetValidate_ZeroN(t *testing.T) {
	for _, mode := range []string{"skip", "start"} {
		c := OffsetConfig{Mode: mode, N: 0}
		if err := c.Validate(); err == nil {
			t.Errorf("expected error for mode=%q with N=0", mode)
		}
	}
}

func TestOffsetValidate_NegativeN(t *testing.T) {
	for _, mode := range []string{"skip", "start"} {
		c := OffsetConfig{Mode: mode, N: -1}
		if err := c.Validate(); err == nil {
			t.Errorf("expected error for mode=%q with N=-1", mode)
		}
	}
}

func TestOffsetEnabled_None(t *testing.T) {
	c := OffsetConfig{Mode: "none", N: 0}
	if c.Enabled() {
		t.Error("expected Enabled()=false for mode 'none'")
	}
}

func TestOffsetEnabled_Active(t *testing.T) {
	for _, mode := range []string{"skip", "start"} {
		c := OffsetConfig{Mode: mode, N: 3}
		if !c.Enabled() {
			t.Errorf("expected Enabled()=true for mode %q", mode)
		}
	}
}
