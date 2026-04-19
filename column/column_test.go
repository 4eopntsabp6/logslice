package column_test

import (
	"testing"

	"github.com/yourorg/logslice/column"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct {
		input string
		want  column.Mode
	}{
		{"", column.ModeNone},
		{"none", column.ModeNone},
		{"extract", column.ModeExtract},
	}
	for _, c := range cases {
		got, err := column.ParseMode(c.input)
		if err != nil || got != c.want {
			t.Errorf("ParseMode(%q) = %v, %v; want %v, nil", c.input, got, err, c.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := column.ParseMode("split")
	if err == nil {
		t.Error("expected error for unknown mode")
	}
}

func TestNew_ModeNoneNoDelimiter(t *testing.T) {
	_, err := column.New(column.ModeNone, "", 0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNew_MissingDelimiter(t *testing.T) {
	_, err := column.New(column.ModeExtract, "", 0)
	if err == nil {
		t.Error("expected error for empty delimiter")
	}
}

func TestNew_NegativeIndex(t *testing.T) {
	_, err := column.New(column.ModeExtract, ",", -1)
	if err == nil {
		t.Error("expected error for negative index")
	}
}

func TestApply_ModeNone(t *testing.T) {
	e, _ := column.New(column.ModeNone, "", 0)
	got := e.Apply("a,b,c")
	if got != "a,b,c" {
		t.Errorf("got %q; want %q", got, "a,b,c")
	}
}

func TestApply_ExtractFirst(t *testing.T) {
	e, _ := column.New(column.ModeExtract, ",", 0)
	got := e.Apply("alpha,beta,gamma")
	if got != "alpha" {
		t.Errorf("got %q; want %q", got, "alpha")
	}
}

func TestApply_ExtractMiddle(t *testing.T) {
	e, _ := column.New(column.ModeExtract, "|", 1)
	got := e.Apply("foo|bar|baz")
	if got != "bar" {
		t.Errorf("got %q; want %q", got, "bar")
	}
}

func TestApply_IndexOutOfRange(t *testing.T) {
	e, _ := column.New(column.ModeExtract, ",", 10)
	got := e.Apply("a,b")
	if got != "a,b" {
		t.Errorf("expected original line, got %q", got)
	}
}
