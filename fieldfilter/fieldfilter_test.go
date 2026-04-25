package fieldfilter

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []string{"none", "exact", "prefix", "regex"}
	for _, c := range cases {
		_, err := ParseMode(c)
		if err != nil {
			t.Errorf("expected valid mode %q, got error: %v", c, err)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("bogus")
	if err == nil {
		t.Error("expected error for invalid mode")
	}
}

func TestNew_ModeNoneNoKey(t *testing.T) {
	f, err := New(ModeNone, "", "")
	if err != nil || f == nil {
		t.Fatal("expected filter with no error")
	}
}

func TestNew_MissingKey(t *testing.T) {
	_, err := New(ModeExact, "", "val")
	if err == nil {
		t.Error("expected error for missing key")
	}
}

func TestNew_MissingValue(t *testing.T) {
	_, err := New(ModeExact, "level", "")
	if err == nil {
		t.Error("expected error for missing value")
	}
}

func TestNew_InvalidRegex(t *testing.T) {
	_, err := New(ModeRegex, "level", "[invalid")
	if err == nil {
		t.Error("expected error for invalid regex")
	}
}

func TestMatch_ModeNone(t *testing.T) {
	f, _ := New(ModeNone, "", "")
	if !f.Match("anything") {
		t.Error("expected match for mode none")
	}
}

func TestMatch_Exact(t *testing.T) {
	f, _ := New(ModeExact, "level", "error")
	if !f.Match("ts=2024-01-01 level=error msg=oops") {
		t.Error("expected match")
	}
	if f.Match("ts=2024-01-01 level=info msg=ok") {
		t.Error("expected no match")
	}
}

func TestMatch_Prefix(t *testing.T) {
	f, _ := New(ModePrefix, "msg", "conn")
	if !f.Match("level=info msg=connected") {
		t.Error("expected prefix match")
	}
	if f.Match("level=info msg=timeout") {
		t.Error("expected no match")
	}
}

func TestMatch_Regex(t *testing.T) {
	f, _ := New(ModeRegex, "level", "^(error|warn)$")
	if !f.Match("level=warn msg=slow") {
		t.Error("expected regex match for warn")
	}
	if f.Match("level=info msg=ok") {
		t.Error("expected no match for info")
	}
}

func TestMatch_EmptyLine(t *testing.T) {
	modes := []struct {
		mode  Mode
		key   string
		value string
	}{
		{ModeExact, "level", "error"},
		{ModePrefix, "level", "err"},
		{ModeRegex, "level", "^error$"},
	}
	for _, m := range modes {
		f, _ := New(m.mode, m.key, m.value)
		if f.Match("") {
			t.Errorf("expected no match for empty line with mode %v", m.mode)
		}
	}
}
