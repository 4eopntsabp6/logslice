package tail_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/yourorg/logslice/tail"
)

func writeTempTailFile(t *testing.T, initial string) *os.File {
	t.Helper()
	f, err := os.CreateTemp(t.TempDir(), "tail-*.log")
	if err != nil {
		t.Fatalf("create temp file: %v", err)
	}
	if initial != "" {
		if _, err := f.WriteString(initial); err != nil {
			t.Fatalf("write initial content: %v", err)
		}
	}
	return f
}

func TestNew_DefaultPollInterval(t *testing.T) {
	tr := tail.New("/tmp/x.log", 0)
	if tr == nil {
		t.Fatal("expected non-nil Tailer")
	}
}

func TestFollow_MissingFile(t *testing.T) {
	tr := tail.New("/nonexistent/path/file.log", tail.DefaultPollInterval)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := tr.Follow(ctx)
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestFollow_EmitsNewLines(t *testing.T) {
	f := writeTempTailFile(t, "")
	defer f.Close()

	tr := tail.New(f.Name(), 20*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	lines, err := tr.Follow(ctx)
	if err != nil {
		t.Fatalf("Follow: %v", err)
	}

	// Append a line after Follow has started.
	time.Sleep(30 * time.Millisecond)
	if _, err := f.WriteString("hello tail\n"); err != nil {
		t.Fatalf("write: %v", err)
	}

	select {
	case got := <-lines:
		if got != "hello tail" {
			t.Errorf("got %q, want %q", got, "hello tail")
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timed out waiting for line")
	}
}

func TestFollow_CancelStopsChannel(t *testing.T) {
	f := writeTempTailFile(t, "")
	defer f.Close()

	tr := tail.New(f.Name(), 20*time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())

	lines, err := tr.Follow(ctx)
	if err != nil {
		t.Fatalf("Follow: %v", err)
	}

	cancel()

	select {
	case _, ok := <-lines:
		if ok {
			t.Error("expected channel to be closed after cancel")
		}
	case <-time.After(500 * time.Millisecond):
		t.Fatal("channel was not closed after context cancel")
	}
}

func TestFollow_EmitsMultipleLines(t *testing.T) {
	f := writeTempTailFile(t, "")
	defer f.Close()

	tr := tail.New(f.Name(), 20*time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	lines, err := tr.Follow(ctx)
	if err != nil {
		t.Fatalf("Follow: %v", err)
	}

	time.Sleep(30 * time.Millisecond)
	if _, err := f.WriteString("line one\nline two\n"); err != nil {
		t.Fatalf("write: %v", err)
	}

	want := []string{"line one", "line two"}
	for _, w := range want {
		select {
		case got := <-lines:
			if got != w {
				t.Errorf("got %q, want %q", got, w)
			}
		case <-time.After(1 * time.Second):
			t.Fatalf("timed out waiting for line %q", w)
		}
	}
}
