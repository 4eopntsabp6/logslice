package field_test

import (
	"testing"

	"github.com/user/logslice/field"
)

func TestParseMode_Valid(t *testing.T) {
	for _, m := range []string{"json", "kv", "none"} {
		out, err := field.ParseMode(m)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", m, err)
		}
		if out != m {
			t.Errorf("ParseMode(%q) = %q, want %q", m, out, m)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := field.ParseMode("xml")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoKey(t *testing.T) {
	_, err := field.New("none", "")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNew_MissingKey(t *testing.T) {
	_, err := field.New("json", "")
	if err == nil {
		t.Error("expected error for empty key with json mode")
	}
}

func TestExtract_ModeNone(t *testing.T) {
	e, _ := field.New("none", "")
	if got := e.Extract(`{"level":"info"}`); got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestExtract_JSON_Found(t *testing.T) {
	e, _ := field.New("json", "level")
	got := e.Extract(`{"level":"error","msg":"fail"}`)
	if got != "error" {
		t.Errorf("expected %q, got %q", "error", got)
	}
}

func TestExtract_JSON_Missing(t *testing.T) {
	e, _ := field.New("json", "level")
	got := e.Extract(`{"msg":"ok"}`)
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestExtract_JSON_InvalidLine(t *testing.T) {
	e, _ := field.New("json", "level")
	if got := e.Extract("not json"); got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}

func TestExtract_KV_Found(t *testing.T) {
	e, _ := field.New("kv", "status")
	got := e.Extract("ts=2024-01-01 status=200 path=/health")
	if got != "200" {
		t.Errorf("expected %q, got %q", "200", got)
	}
}

func TestExtract_KV_Missing(t *testing.T) {
	e, _ := field.New("kv", "status")
	got := e.Extract("ts=2024-01-01 path=/health")
	if got != "" {
		t.Errorf("expected empty, got %q", got)
	}
}
