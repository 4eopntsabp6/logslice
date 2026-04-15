package output_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/user/logslice/output"
)

func TestNewFormatter_Valid(t *testing.T) {
	formats := []output.Format{output.FormatPlain, output.FormatJSON, output.FormatCSV}
	for _, f := range formats {
		_, err := output.NewFormatter(f, &bytes.Buffer{})
		if err != nil {
			t.Errorf("expected no error for format %q, got %v", f, err)
		}
	}
}

func TestNewFormatter_Invalid(t *testing.T) {
	_, err := output.NewFormatter(output.Format("xml"), &bytes.Buffer{})
	if err == nil {
		t.Error("expected error for unsupported format, got nil")
	}
}

func TestWriteLine_Plain(t *testing.T) {
	var buf bytes.Buffer
	f, _ := output.NewFormatter(output.FormatPlain, &buf)
	_ = f.WriteLine(1, "hello world")
	if got := strings.TrimSpace(buf.String()); got != "hello world" {
		t.Errorf("plain: got %q, want %q", got, "hello world")
	}
}

func TestWriteLine_JSON(t *testing.T) {
	var buf bytes.Buffer
	f, _ := output.NewFormatter(output.FormatJSON, &buf)
	_ = f.WriteLine(3, "log entry")
	got := strings.TrimSpace(buf.String())
	if !strings.Contains(got, `"index":3`) {
		t.Errorf("json: missing index field in %q", got)
	}
	if !strings.Contains(got, "log entry") {
		t.Errorf("json: missing line content in %q", got)
	}
}

func TestWriteLine_CSV(t *testing.T) {
	var buf bytes.Buffer
	f, _ := output.NewFormatter(output.FormatCSV, &buf)
	_ = f.WriteLine(2, "some,data")
	got := strings.TrimSpace(buf.String())
	if !strings.HasPrefix(got, "2,") {
		t.Errorf("csv: expected line to start with index, got %q", got)
	}
}

func TestParseFormat_Valid(t *testing.T) {
	cases := map[string]output.Format{
		"plain": output.FormatPlain,
		"JSON":  output.FormatJSON,
		"Csv":   output.FormatCSV,
	}
	for input, want := range cases {
		got, err := output.ParseFormat(input)
		if err != nil {
			t.Errorf("ParseFormat(%q): unexpected error %v", input, err)
		}
		if got != want {
			t.Errorf("ParseFormat(%q): got %q, want %q", input, got, want)
		}
	}
}

func TestParseFormat_Invalid(t *testing.T) {
	_, err := output.ParseFormat("toml")
	if err == nil {
		t.Error("expected error for unknown format, got nil")
	}
}
