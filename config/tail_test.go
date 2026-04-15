package config

import (
	"testing"
	"time"
)

func TestDefaultTailConfig(t *testing.T) {
	tc := DefaultTailConfig()
	if tc.Enabled {
		t.Error("expected Enabled to be false by default")
	}
	if tc.PollInterval != 200*time.Millisecond {
		t.Errorf("unexpected default PollInterval: %v", tc.PollInterval)
	}
	if tc.MaxLines != 0 {
		t.Errorf("unexpected default MaxLines: %d", tc.MaxLines)
	}
}

func TestTailValidate_Disabled(t *testing.T) {
	tc := TailConfig{Enabled: false, PollInterval: -1, MaxLines: -99}
	if err := tc.Validate(); err != nil {
		t.Errorf("disabled config should always be valid, got: %v", err)
	}
}

func TestTailValidate_InvalidPollInterval(t *testing.T) {
	tc := TailConfig{Enabled: true, PollInterval: 0, MaxLines: 0}
	if err := tc.Validate(); err == nil {
		t.Error("expected error for zero PollInterval")
	}
}

func TestTailValidate_NegativePollInterval(t *testing.T) {
	tc := TailConfig{Enabled: true, PollInterval: -5 * time.Millisecond, MaxLines: 0}
	if err := tc.Validate(); err == nil {
		t.Error("expected error for negative PollInterval")
	}
}

func TestTailValidate_NegativeMaxLines(t *testing.T) {
	tc := TailConfig{Enabled: true, PollInterval: 100 * time.Millisecond, MaxLines: -1}
	if err := tc.Validate(); err == nil {
		t.Error("expected error for negative MaxLines")
	}
}

func TestTailValidate_ValidConfig(t *testing.T) {
	cases := []TailConfig{
		{Enabled: true, PollInterval: 100 * time.Millisecond, MaxLines: 0},
		{Enabled: true, PollInterval: 500 * time.Millisecond, MaxLines: 100},
	}
	for _, tc := range cases {
		if err := tc.Validate(); err != nil {
			t.Errorf("expected valid config, got: %v", err)
		}
	}
}
