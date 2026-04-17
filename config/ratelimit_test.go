package config_test

import (
	"testing"

	"github.com/user/logslice/config"
)

func TestDefaultRateLimitConfig(t *testing.T) {
	c := config.DefaultRateLimitConfig()
	if c.Mode != "none" {
		t.Errorf("expected mode none, got %s", c.Mode)
	}
	if c.Limit != 0 {
		t.Errorf("expected limit 0, got %d", c.Limit)
	}
}

func TestRateLimitValidate_ValidModes(t *testing.T) {
	for _, tc := range []struct {
		mode  string
		limit int
	}{
		{"none", 0},
		{"lines", 10},
	} {
		c := config.RateLimitConfig{Mode: tc.mode, Limit: tc.limit}
		if err := c.Validate(); err != nil {
			t.Errorf("unexpected error for mode=%s: %v", tc.mode, err)
		}
	}
}

func TestRateLimitValidate_InvalidMode(t *testing.T) {
	c := config.RateLimitConfig{Mode: "burst", Limit: 5}
	if err := c.Validate(); err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestRateLimitValidate_LinesNoLimit(t *testing.T) {
	c := config.RateLimitConfig{Mode: "lines", Limit: 0}
	if err := c.Validate(); err == nil {
		t.Error("expected error for lines mode with limit=0")
	}
}

func TestRateLimitEnabled(t *testing.T) {
	none := config.RateLimitConfig{Mode: "none"}
	if none.Enabled() {
		t.Error("expected Enabled=false for none mode")
	}
	lines := config.RateLimitConfig{Mode: "lines", Limit: 5}
	if !lines.Enabled() {
		t.Error("expected Enabled=true for lines mode")
	}
}
