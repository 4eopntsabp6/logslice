package pipeline_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/yourorg/logslice/config"
	"github.com/yourorg/logslice/pipeline"
)

func writeTempLog(t *testing.T, lines []string) string {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "log-*.txt")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	defer f.Close()
	f.WriteString(strings.Join(lines, "\n") + "\n")
	return f.Name()
}

func baseConfig(inputPath, outputPath string) *config.Config {
	return &config.Config{
		InputPath:  inputPath,
		OutputPath: outputPath,
		Format:     "plain",
	}
}

func TestNew_ValidConfig(t *testing.T) {
	lines := []string{
		"2024-01-01T10:00:00Z INFO starting",
		"2024-01-01T10:01:00Z DEBUG tick",
	}
	input := writeTempLog(t, lines)
	output := filepath.Join(t.TempDir(), "out.txt")

	p, err := pipeline.New(baseConfig(input, output))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer p.Close()
}

func TestNew_MissingInput(t *testing.T) {
	cfg := baseConfig("/no/such/file.log", "stdout")
	_, err := pipeline.New(cfg)
	if err == nil {
		t.Fatal("expected error for missing input file")
	}
}

func TestNew_InvalidPattern(t *testing.T) {
	lines := []string{"2024-01-01T10:00:00Z INFO ok"}
	input := writeTempLog(t, lines)
	cfg := baseConfig(input, "stdout")
	cfg.Pattern = "[invalid"

	_, err := pipeline.New(cfg)
	if err == nil {
		t.Fatal("expected error for invalid regex pattern")
	}
}

func TestRun_ReturnsStats(t *testing.T) {
	lines := []string{
		"2024-01-01T10:00:00Z INFO hello",
		"2024-01-01T10:02:00Z INFO world",
	}
	input := writeTempLog(t, lines)
	output := filepath.Join(t.TempDir(), "out.txt")

	cfg := baseConfig(input, output)
	from := time.Date(2024, 1, 1, 9, 59, 0, 0, time.UTC)
	to := time.Date(2024, 1, 1, 10, 5, 0, 0, time.UTC)
	cfg.From = &from
	cfg.To = &to

	p, err := pipeline.New(cfg)
	if err != nil {
		t.Fatalf("pipeline.New: %v", err)
	}
	defer p.Close()

	st, err := p.Run()
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	if st == nil {
		t.Fatal("expected non-nil stats")
	}
	if st.TotalLines() == 0 {
		t.Error("expected at least one line processed")
	}
}
