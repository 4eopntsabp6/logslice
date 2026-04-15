package config

import (
	"testing"
	"time"
)

func TestValidate_Defaults(t *testing.T) {
	c := &Config{}
	if err := c.Validate(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if c.OutputFormat != "plain" {
		t.Errorf("expected default format 'plain', got %q", c.OutputFormat)
	}
	if c.OutputDest != "stdout" {
		t.Errorf("expected default dest 'stdout', got %q", c.OutputDest)
	}
}

func TestValidate_InvalidFormat(t *testing.T) {
	c := &Config{OutputFormat: "xml"}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error for invalid format, got nil")
	}
}

func TestValidate_ToBeforeFrom(t *testing.T) {
	now := time.Now()
	c := &Config{
		From: now,
		To:   now.Add(-time.Hour),
	}
	if err := c.Validate(); err == nil {
		t.Fatal("expected error when To is before From, got nil")
	}
}

func TestValidate_EqualFromTo(t *testing.T) {
	now := time.Now()
	c := &Config{From: now, To: now}
	if err := c.Validate(); err != nil {
		t.Fatalf("expected no error for equal From/To, got %v", err)
	}
}

func TestHasTimeWindow_NoWindow(t *testing.T) {
	c := &Config{}
	if c.HasTimeWindow() {
		t.Error("expected HasTimeWindow to be false when no times set")
	}
}

func TestHasTimeWindow_FromOnly(t *testing.T) {
	c := &Config{From: time.Now()}
	if !c.HasTimeWindow() {
		t.Error("expected HasTimeWindow to be true when From is set")
	}
}

func TestHasTimeWindow_ToOnly(t *testing.T) {
	c := &Config{To: time.Now()}
	if !c.HasTimeWindow() {
		t.Error("expected HasTimeWindow to be true when To is set")
	}
}

func TestHasTimeWindow_Both(t *testing.T) {
	now := time.Now()
	c := &Config{From: now, To: now.Add(time.Hour)}
	if !c.HasTimeWindow() {
		t.Error("expected HasTimeWindow to be true when both times set")
	}
}
