package highlight_test

import (
	"strings"
	"testing"

	"github.com/yourorg/logslice/highlight"
)

func TestNew_NoPattern(t *testing.T) {
	h, err := highlight.New("", highlight.Cyan, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := h.Apply("hello"); got != "hello" {
		t.Errorf("expected unchanged line, got %q", got)
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	_, err := highlight.New("[", highlight.Red, true)
	if err == nil {
		t.Fatal("expected error for invalid regex")
	}
}

func TestApply_Disabled(t *testing.T) {
	h, _ := highlight.New("ERROR", highlight.Red, false)
	line := "this is an ERROR line"
	if got := h.Apply(line); got != line {
		t.Errorf("expected unchanged line when disabled, got %q", got)
	}
}

func TestApply_MatchesPattern(t *testing.T) {
	h, _ := highlight.New("ERROR", highlight.Red, true)
	result := h.Apply("prefix ERROR suffix")
	if !strings.Contains(result, "ERROR") {
		t.Error("result should still contain the matched text")
	}
	if !strings.Contains(result, highlight.Red) {
		t.Error("result should contain colour escape code")
	}
	if !strings.Contains(result, highlight.Reset) {
		t.Error("result should contain reset escape code")
	}
}

func TestApply_NoMatch(t *testing.T) {
	h, _ := highlight.New("ERROR", highlight.Red, true)
	line := "everything is fine"
	if got := h.Apply(line); got != line {
		t.Errorf("expected unchanged line when no match, got %q", got)
	}
}

func TestStrip_RemovesEscapes(t *testing.T) {
	input := highlight.Red + highlight.Bold + "hello" + highlight.Reset
	if got := highlight.Strip(input); got != "hello" {
		t.Errorf("Strip() = %q, want %q", got, "hello")
	}
}

func TestParseColour(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"red", highlight.Red},
		{"RED", highlight.Red},
		{"yellow", highlight.Yellow},
		{"cyan", highlight.Cyan},
		{"unknown", highlight.Cyan},
		{"", highlight.Cyan},
	}
	for _, tc := range cases {
		if got := highlight.ParseColour(tc.input); got != tc.want {
			t.Errorf("ParseColour(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}
