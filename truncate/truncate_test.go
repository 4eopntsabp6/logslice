package truncate

import (
	"strings"
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  Mode
	}{
		{"", ModeNone},
		{"none", ModeNone},
		{"bytes", ModeBytes},
		{"runes", ModeRunes},
	}
	for _, c := range cases {
		got, err := ParseMode(c.input)
		if err != nil {
			t.Errorf("ParseMode(%q) unexpected error: %v", c.input, err)
		}
		if got != c.want {
			t.Errorf("ParseMode(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("chars")
	if err == nil {
		t.Error("expected error for unknown mode, got nil")
	}
}

func TestNew_InvalidLimit(t *testing.T) {
	_, err := New(ModeBytes, 0, "")
	if err == nil {
		t.Error("expected error for limit=0, got nil")
	}
	_, err = New(ModeRunes, -1, "")
	if err == nil {
		t.Error("expected error for negative limit, got nil")
	}
}

func TestNew_ModeNoneIgnoresLimit(t *testing.T) {
	_, err := New(ModeNone, 0, "")
	if err != nil {
		t.Errorf("unexpected error for ModeNone with limit=0: %v", err)
	}
}

func TestApply_ModeNone(t *testing.T) {
	tr, _ := New(ModeNone, 0, "")
	line := strings.Repeat("a", 200)
	if got := tr.Apply(line); got != line {
		t.Error("ModeNone should not modify line")
	}
}

func TestApply_Bytes_ShortLine(t *testing.T) {
	tr, _ := New(ModeBytes, 50, "")
	line := "short"
	if got := tr.Apply(line); got != line {
		t.Errorf("expected %q, got %q", line, got)
	}
}

func TestApply_Bytes_LongLine(t *testing.T) {
	tr, _ := New(ModeBytes, 10, "---")
	line := "hello world this is a long line"
	got := tr.Apply(line)
	if !strings.HasSuffix(got, "---") {
		t.Errorf("expected suffix '---', got %q", got)
	}
	if len(got) > 10+len("---") {
		t.Errorf("result too long: %d bytes", len(got))
	}
}

func TestApply_Runes_LongLine(t *testing.T) {
	tr, _ := New(ModeRunes, 5, "...")
	line := "abcdefghij"
	got := tr.Apply(line)
	if got != "abcde..." {
		t.Errorf("expected %q, got %q", "abcde...", got)
	}
}

func TestApply_Runes_Unicode(t *testing.T) {
	tr, _ := New(ModeRunes, 3, "")
	line := "日本語テスト"
	got := tr.Apply(line)
	// first 3 runes + DefaultSuffix
	if got != "日本語"+DefaultSuffix {
		t.Errorf("unexpected result: %q", got)
	}
}

func TestApply_DefaultSuffix(t *testing.T) {
	tr, _ := New(ModeBytes, 4, "")
	got := tr.Apply("hello world")
	if !strings.HasSuffix(got, DefaultSuffix) {
		t.Errorf("expected default suffix %q in %q", DefaultSuffix, got)
	}
}
