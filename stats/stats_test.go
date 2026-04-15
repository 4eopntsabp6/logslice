package stats_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/logslice/stats"
)

func TestNew_InitializesStartTime(t *testing.T) {
	before := time.Now()
	s := stats.New()
	after := time.Now()

	if s.StartTime.Before(before) || s.StartTime.After(after) {
		t.Errorf("StartTime %v not in expected range [%v, %v]", s.StartTime, before, after)
	}
}

func TestRecordLine_Matched(t *testing.T) {
	s := stats.New()
	s.RecordLine(true, 42)

	if s.LinesRead != 1 {
		t.Errorf("expected LinesRead=1, got %d", s.LinesRead)
	}
	if s.LinesMatched != 1 {
		t.Errorf("expected LinesMatched=1, got %d", s.LinesMatched)
	}
	if s.LinesDropped != 0 {
		t.Errorf("expected LinesDropped=0, got %d", s.LinesDropped)
	}
	if s.BytesRead != 42 {
		t.Errorf("expected BytesRead=42, got %d", s.BytesRead)
	}
}

func TestRecordLine_NotMatched(t *testing.T) {
	s := stats.New()
	s.RecordLine(false, 10)

	if s.LinesDropped != 1 {
		t.Errorf("expected LinesDropped=1, got %d", s.LinesDropped)
	}
	if s.LinesMatched != 0 {
		t.Errorf("expected LinesMatched=0, got %d", s.LinesMatched)
	}
}

func TestMatchRate_Zero(t *testing.T) {
	s := stats.New()
	if s.MatchRate() != 0.0 {
		t.Errorf("expected MatchRate=0.0 for empty stats, got %f", s.MatchRate())
	}
}

func TestMatchRate_Half(t *testing.T) {
	s := stats.New()
	s.RecordLine(true, 5)
	s.RecordLine(false, 5)

	if s.MatchRate() != 50.0 {
		t.Errorf("expected MatchRate=50.0, got %f", s.MatchRate())
	}
}

func TestFinish_SetsEndTime(t *testing.T) {
	s := stats.New()
	s.Finish()

	if s.EndTime.IsZero() {
		t.Error("expected EndTime to be set after Finish")
	}
}

func TestDuration_AfterFinish(t *testing.T) {
	s := stats.New()
	time.Sleep(2 * time.Millisecond)
	s.Finish()

	if s.Duration() < time.Millisecond {
		t.Errorf("expected duration >= 1ms, got %v", s.Duration())
	}
}

func TestString_ContainsFields(t *testing.T) {
	s := stats.New()
	s.RecordLine(true, 100)
	s.RecordLine(false, 50)
	s.Finish()

	str := s.String()
	for _, want := range []string{"read=", "matched=", "dropped=", "bytes=", "duration=", "match_rate="} {
		if !strings.Contains(str, want) {
			t.Errorf("String() missing field %q, got: %s", want, str)
		}
	}
}
