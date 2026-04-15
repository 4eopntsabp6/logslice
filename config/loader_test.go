package config

import (
	"testing"
	"time"
)

func TestLoadTimes_Empty(t *testing.T) {
	cfg := &Config{}
	if err := loadTimes(cfg, "", ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.From.IsZero() || !cfg.To.IsZero() {
		t.Error("expected zero times when no args provided")
	}
}

func TestLoadTimes_ValidRFC3339(t *testing.T) {
	cfg := &Config{}
	const ts = "2024-01-15T10:00:00Z"
	if err := loadTimes(cfg, ts, ""); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected, _ := time.Parse(time.RFC3339, ts)
	if !cfg.From.Equal(expected) {
		t.Errorf("From mismatch: got %v, want %v", cfg.From, expected)
	}
}

func TestLoadTimes_InvalidFrom(t *testing.T) {
	cfg := &Config{}
	if err := loadTimes(cfg, "not-a-time", ""); err == nil {
		t.Fatal("expected error for invalid --from, got nil")
	}
}

func TestLoadTimes_InvalidTo(t *testing.T) {
	cfg := &Config{}
	if err := loadTimes(cfg, "", "bad-value"); err == nil {
		t.Fatal("expected error for invalid --to, got nil")
	}
}

func TestLoadTimes_BothValid(t *testing.T) {
	cfg := &Config{}
	fromStr := "2024-01-15T08:00:00Z"
	toStr := "2024-01-15T18:00:00Z"
	if err := loadTimes(cfg, fromStr, toStr); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.From.IsZero() || cfg.To.IsZero() {
		t.Error("expected non-zero From and To")
	}
	if !cfg.To.After(cfg.From) {
		t.Error("expected To to be after From")
	}
}
