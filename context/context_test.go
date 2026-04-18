package context

import (
	"testing"
)

func TestNew_ModeNone(t *testing.T) {
	b, err := New(ModeNone, 0, 0)
	if err != nil || b == nil {
		t.Fatal("expected valid buffer")
	}
}

func TestNew_NegativeBefore(t *testing.T) {
	_, err := New(ModeBefore, -1, 0)
	if err == nil {
		t.Fatal("expected error for negative before")
	}
}

func TestNew_NegativeAfter(t *testing.T) {
	_, err := New(ModeAfter, 0, -1)
	if err == nil {
		t.Fatal("expected error for negative after")
	}
}

func TestBefore_CapturesLines(t *testing.T) {
	b, _ := New(ModeBefore, 2, 0)
	b.Feed("line1")
	b.Feed("line2")
	b.Feed("line3")
	got := b.Before()
	if len(got) != 2 || got[0] != "line2" || got[1] != "line3" {
		t.Fatalf("unexpected before lines: %v", got)
	}
}

func TestBefore_FewerThanWindow(t *testing.T) {
	b, _ := New(ModeBefore, 3, 0)
	b.Feed("only")
	got := b.Before()
	if len(got) != 1 || got[0] != "only" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestAfter_EmitsPostMatch(t *testing.T) {
	b, _ := New(ModeAfter, 0, 2)
	b.OnMatch()
	if !b.Feed("a1") {
		t.Fatal("expected after line 1")
	}
	if !b.Feed("a2") {
		t.Fatal("expected after line 2")
	}
	if b.Feed("a3") {
		t.Fatal("should not emit after window exhausted")
	}
}

func TestBoth_BeforeAndAfter(t *testing.T) {
	b, _ := New(ModeBoth, 2, 1)
	b.Feed("pre1")
	b.Feed("pre2")
	lines := b.Before()
	if len(lines) != 2 {
		t.Fatalf("expected 2 before lines, got %d", len(lines))
	}
	b.OnMatch()
	if !b.Feed("post1") {
		t.Fatal("expected post-match line")
	}
	if b.Feed("post2") {
		t.Fatal("should not emit second after line")
	}
}

func TestModeNone_FeedAlwaysFalse(t *testing.T) {
	b, _ := New(ModeNone, 0, 0)
	b.OnMatch()
	if b.Feed("x") {
		t.Fatal("ModeNone should never emit context")
	}
}
