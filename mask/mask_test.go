package mask

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"", ModeNone},
		{"none", ModeNone},
		{"redact", ModeRedact},
		{"partial", ModePartial},
		{"REDACT", ModeRedact},
	}
	for _, c := range cases {
		got, err := ParseMode(c.input)
		if err != nil || got != c.want {
			t.Errorf("ParseMode(%q) = %v, %v; want %v, nil", c.input, got, err, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("scramble")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoPattern(t *testing.T) {
	m, err := New(ModeNone, "")
	if err != nil || m == nil {
		t.Fatalf("unexpected: %v %v", m, err)
	}
}

func TestNew_MissingPattern(t *testing.T) {
	_, err := New(ModeRedact, "")
	if err == nil {
		t.Error("expected error when pattern is empty")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := New(ModeRedact, "[invalid")
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}

func TestApply_ModeNone(t *testing.T) {
	m, _ := New(ModeNone, "")
	got := m.Apply("token=abc123")
	if got != "token=abc123" {
		t.Errorf("unexpected: %q", got)
	}
}

func TestApply_Redact(t *testing.T) {
	m, _ := New(ModeRedact, `token=\S+`)
	got := m.Apply("auth token=abc123 ok")
	want := "auth [REDACTED] ok"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestApply_Partial(t *testing.T) {
	m, _ := New(ModePartial, `\b[0-9]{6,}\b`)
	got := m.Apply("id=123456")
	if got == "id=123456" {
		t.Error("expected masking to occur")
	}
	if got[:5] != "id=12" {
		t.Errorf("expected prefix preserved, got %q", got)
	}
}
