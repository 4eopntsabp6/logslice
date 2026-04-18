package linenum

import (
	"testing"
)

func TestParseMode_Valid(t *testing.T) {
	cases := []struct{ in string; want Mode }{
		{"", ModeNone},
		{"none", ModeNone},
		{"range", ModeRange},
		{"list", ModeList},
	}
	for _, c := range cases {
		got, err := ParseMode(c.in)
		if err != nil || got != c.want {
			t.Errorf("ParseMode(%q) = %v, %v", c.in, got, err)
		}
	}
}

func TestParseMode_Invalid(t *testing.T) {
	_, err := ParseMode("unknown")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestNew_ModeNone_AllowsAll(t *testing.T) {
	f, err := New(ModeNone, 0, 0, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, n := range []int{1, 50, 999} {
		if !f.Keep(n) {
			t.Errorf("expected Keep(%d) = true", n)
		}
	}
}

func TestNew_RangeInvalidFrom(t *testing.T) {
	_, err := New(ModeRange, 0, 0, nil)
	if err == nil {
		t.Fatal("expected error for from < 1")
	}
}

func TestNew_RangeToBeforeFrom(t *testing.T) {
	_, err := New(ModeRange, 5, 3, nil)
	if err == nil {
		t.Fatal("expected error for to < from")
	}
}

func TestKeep_Range(t *testing.T) {
	f, err := New(ModeRange, 3, 6, nil)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ n int; want bool }{
		{1, false}, {2, false}, {3, true}, {5, true}, {6, true}, {7, false},
	}
	for _, c := range cases {
		if got := f.Keep(c.n); got != c.want {
			t.Errorf("Keep(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}

func TestKeep_List(t *testing.T) {
	f, err := New(ModeList, 0, 0, []int{2, 4, 7})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ n int; want bool }{
		{1, false}, {2, true}, {3, false}, {4, true}, {7, true}, {8, false},
	}
	for _, c := range cases {
		if got := f.Keep(c.n); got != c.want {
			t.Errorf("Keep(%d) = %v, want %v", c.n, got, c.want)
		}
	}
}
