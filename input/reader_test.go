package input_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/yourorg/logslice/input"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.log")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writeTempFile: %v", err)
	}
	return path
}

func TestNewFileReader_Valid(t *testing.T) {
	path := writeTempFile(t, "line one\nline two\nline three\n")
	r, err := input.NewFileReader(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer r.Close()

	expected := []string{"line one", "line two", "line three"}
	for i, want := range expected {
		got, ok := r.ReadLine()
		if !ok {
			t.Fatalf("line %d: expected more lines", i+1)
		}
		if got != want {
			t.Errorf("line %d: got %q, want %q", i+1, got, want)
		}
		if r.LineNum() != i+1 {
			t.Errorf("LineNum: got %d, want %d", r.LineNum(), i+1)
		}
	}
	_, ok := r.ReadLine()
	if ok {
		t.Error("expected no more lines")
	}
	if err := r.Err(); err != nil {
		t.Errorf("unexpected scanner error: %v", err)
	}
}

func TestNewFileReader_Missing(t *testing.T) {
	_, err := input.NewFileReader("/nonexistent/path/file.log")
	if err == nil {
		t.Fatal("expected error for missing file, got nil")
	}
}

func TestNewFileReader_Empty(t *testing.T) {
	path := writeTempFile(t, "")
	r, err := input.NewFileReader(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer r.Close()

	_, ok := r.ReadLine()
	if ok {
		t.Error("expected no lines from empty file")
	}
	if r.LineNum() != 0 {
		t.Errorf("LineNum: got %d, want 0", r.LineNum())
	}
}

// TestNewFileReader_NoTrailingNewline verifies that a file whose last line
// lacks a trailing newline is still read correctly.
func TestNewFileReader_NoTrailingNewline(t *testing.T) {
	path := writeTempFile(t, "alpha\nbeta\ngamma")
	r, err := input.NewFileReader(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer r.Close()

	expected := []string{"alpha", "beta", "gamma"}
	for i, want := range expected {
		got, ok := r.ReadLine()
		if !ok {
			t.Fatalf("line %d: expected more lines", i+1)
		}
		if got != want {
			t.Errorf("line %d: got %q, want %q", i+1, got, want)
		}
	}
	_, ok := r.ReadLine()
	if ok {
		t.Error("expected no more lines after last line without newline")
	}
}

func TestNewStdinReader_NotNil(t *testing.T) {
	r := input.NewStdinReader()
	if r == nil {
		t.Fatal("expected non-nil LineReader for stdin")
	}
	// Do not read; just verify Close does not panic.
	if err := r.Close(); err != nil {
		t.Errorf("Close: unexpected error: %v", err)
	}
}
