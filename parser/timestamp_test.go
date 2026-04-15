package parser

import (
	"testing"
	"time"
)

func TestExtractTimestamp(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantOK  bool
		wantUTC string // expected time in RFC3339 UTC, empty if wantOK==false
	}{
		{
			name:    "RFC3339 with Z",
			line:    `2024-03-15T10:22:01Z level=info msg="service started"`,
			wantOK:  true,
			wantUTC: "2024-03-15T10:22:01Z",
		},
		{
			name:    "RFC3339 with offset",
			line:    `2024-03-15T10:22:01+02:00 ERROR connection refused`,
			wantOK:  true,
			wantUTC: "2024-03-15T08:22:01Z",
		},
		{
			name:    "datetime with space separator",
			line:    `2024-03-15 10:22:01 WARN retrying request`,
			wantOK:  true,
			wantUTC: "2024-03-15T10:22:01Z",
		},
		{
			name:   "no timestamp",
			line:   `this log line has no timestamp`,
			wantOK: false,
		},
		{
			name:   "empty line",
			line:   ``,
			wantOK: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := ExtractTimestamp(tc.line)
			if ok != tc.wantOK {
				t.Fatalf("ExtractTimestamp(%q) ok=%v, want %v", tc.line, ok, tc.wantOK)
			}
			if tc.wantOK {
				want, _ := time.Parse(time.RFC3339, tc.wantUTC)
				if !got.UTC().Equal(want.UTC()) {
					t.Errorf("ExtractTimestamp(%q) = %v, want %v", tc.line, got.UTC(), want.UTC())
				}
			}
		})
	}
}

func TestParseTimeArg(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"2024-03-15T10:22:01Z", false},
		{"2024-03-15 10:22:01", false},
		{"2024-03-15T10:22:01", false},
		{"not-a-time", true},
		{"", true},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			_, err := ParseTimeArg(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("ParseTimeArg(%q) error=%v, wantErr=%v", tc.input, err, tc.wantErr)
			}
		})
	}
}
