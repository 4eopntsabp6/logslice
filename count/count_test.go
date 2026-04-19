package count

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct{ in string; want Mode }{
		{"", ModeNone},
		{"none", ModeNone},
		{"limit", ModeLimit},
		{"skip", ModeSkip},
	}
	for _, tc := range cases {
		got, err := ParseMode(tc.in)
		if err != nil || got != tc.want {
			t.Errorf("ParseMode(%q) = %v, %v; want %v, nil", tc.in, got, err, tc.want)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("bogus")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNew_InvalidN(t *testing.T) {
	_, err := New(ModeLimit, 0)
	if err == nil {
		t.Fatal("expected error for n=0")
	}
}

func TestNew_ModeNoneIgnoresN(t *testing.T) {
	_, err := New(ModeNone, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestKeep_ModeNone_AllowsAll(t *testing.T) {
	c, _ := New(ModeNone, 0)
	for i := 0; i < 100; i++ {
		if !c.Keep("line") {
			t.Fatal("ModeNone should allow all lines")
		}
	}
}

func TestKeep_ModeLimit(t *testing.T) {
	c, _ := New(ModeLimit, 3)
	results := []bool{c.Keep("a"), c.Keep("b"), c.Keep("c"), c.Keep("d")}
	expect := []bool{true, true, true, false}
	for i, r := range results {
		if r != expect[i] {
			t.Errorf("line %d: got %v want %v", i+1, r, expect[i])
		}
	}
}

func TestKeep_ModeSkip(t *testing.T) {
	c, _ := New(ModeSkip, 2)
	results := []bool{c.Keep("a"), c.Keep("b"), c.Keep("c"), c.Keep("d")}
	expect := []bool{false, false, true, true}
	for i, r := range results {
		if r != expect[i] {
			t.Errorf("line %d: got %v want %v", i+1, r, expect[i])
		}
	}
}

func TestDone_ModeLimit(t *testing.T) {
	c, _ := New(ModeLimit, 2)
	c.Keep("a")
	if c.Done() {
		t.Fatal("should not be done after 1 of 2")
	}
	c.Keep("b")
	if !c.Done() {
		t.Fatal("should be done after 2 of 2")
	}
}
