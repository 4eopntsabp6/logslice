package output_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/logslice/output"
)

func TestOpenDestination_Stdout(t *testing.T) {
	w, err := output.OpenDestination(output.DestStdout, "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer w.Close()
	if w == nil {
		t.Error("expected non-nil writer for stdout")
	}
}

func TestOpenDestination_File(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.log")

	w, err := output.OpenDestination(output.DestFile, path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, writeErr := w.Write([]byte("test line\n"))
	w.Close()

	if writeErr != nil {
		t.Fatalf("write error: %v", writeErr)
	}
	data, _ := os.ReadFile(path)
	if string(data) != "test line\n" {
		t.Errorf("file content mismatch: got %q", data)
	}
}

func TestOpenDestination_File_EmptyPath(t *testing.T) {
	_, err := output.OpenDestination(output.DestFile, "")
	if err == nil {
		t.Error("expected error for empty file path")
	}
}

func TestOpenDestination_Unknown(t *testing.T) {
	_, err := output.OpenDestination(output.Destination("s3"), "")
	if err == nil {
		t.Error("expected error for unknown destination")
	}
}

func TestParseDestination_Valid(t *testing.T) {
	for _, s := range []string{"stdout", "file"} {
		_, err := output.ParseDestination(s)
		if err != nil {
			t.Errorf("ParseDestination(%q): unexpected error %v", s, err)
		}
	}
}

func TestParseDestination_Invalid(t *testing.T) {
	_, err := output.ParseDestination("stderr")
	if err == nil {
		t.Error("expected error for unknown destination")
	}
}
