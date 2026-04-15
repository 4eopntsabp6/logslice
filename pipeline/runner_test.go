package pipeline_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/user/logslice/config"
	"github.com/user/logslice/pipeline"
)

func writeRunnerLog(t *testing.T, lines []string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "run.log")
	content := strings.Join(lines, "\n") + "\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write temp log: %v", err)
	}
	return path
}

func TestNewRunner_Valid(t *testing.T) {
	path := writeRunnerLog(t, []string{
		"2024-01-01T10:00:00Z INFO starting",
		"2024-01-01T10:01:00Z INFO running",
	})
	cfg := &config.Config{Input: path, Format: "plain"}
	var buf bytes.Buffer
	r, err := pipeline.NewRunner(cfg, &buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil runner")
	}
}

func TestNewRunner_MissingInput(t *testing.T) {
	cfg := &config.Config{Input: "/nonexistent/file.log", Format: "plain"}
	var buf bytes.Buffer
	_, err := pipeline.NewRunner(cfg, &buf)
	if err == nil {
		t.Fatal("expected error for missing input")
	}
}

func TestRunner_Run_WritesOutput(t *testing.T) {
	path := writeRunnerLog(t, []string{
		"2024-01-01T10:00:00Z INFO hello",
		"2024-01-01T10:01:00Z INFO world",
	})
	cfg := &config.Config{
		Input:  path,
		Format: "plain",
		From:   time.Time{},
		To:     time.Time{},
	}
	var statBuf, outBuf bytes.Buffer
	r, err := pipeline.NewRunner(cfg, &statBuf)
	if err != nil {
		t.Fatalf("NewRunner: %v", err)
	}
	if err := r.Run(&outBuf); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if outBuf.Len() == 0 {
		t.Error("expected non-empty output")
	}
}

func TestRunner_Run_ShowStats(t *testing.T) {
	path := writeRunnerLog(t, []string{
		"2024-01-01T10:00:00Z INFO entry",
	})
	cfg := &config.Config{
		Input:     path,
		Format:    "plain",
		ShowStats: true,
	}
	var statBuf, outBuf bytes.Buffer
	r, err := pipeline.NewRunner(cfg, &statBuf)
	if err != nil {
		t.Fatalf("NewRunner: %v", err)
	}
	if err := r.Run(&outBuf); err != nil {
		t.Fatalf("Run: %v", err)
	}
	if statBuf.Len() == 0 {
		t.Error("expected stats output when ShowStats=true")
	}
}
