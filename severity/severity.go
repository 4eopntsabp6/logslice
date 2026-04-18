package severity

import (
	"fmt"
	"strings"
)

// Mode controls how severity filtering is applied.
type Mode string

const (
	ModeNone  Mode = "none"
	ModeMin   Mode = "min"
	ModeExact Mode = "exact"
)

// Level represents a log severity level.
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

var levelNames = map[string]Level{
	"debug": LevelDebug,
	"info":  LevelInfo,
	"warn":  LevelWarn,
	"error": LevelError,
	"fatal": LevelFatal,
}

// ParseLevel parses a string into a Level.
func ParseLevel(s string) (Level, error) {
	l, ok := levelNames[strings.ToLower(s)]
	if !ok {
		return 0, fmt.Errorf("unknown severity level: %q", s)
	}
	return l, nil
}

// ParseMode parses a string into a Mode.
func ParseMode(s string) (Mode, error) {
	switch Mode(strings.ToLower(s)) {
	case ModeNone, ModeMin, ModeExact:
		return Mode(strings.ToLower(s)), nil
	}
	return "", fmt.Errorf("unknown severity mode: %q", s)
}

// Filter holds the severity filter configuration.
type Filter struct {
	mode  Mode
	level Level
}

// New creates a new Filter. For ModeNone, level is ignored.
func New(mode Mode, levelStr string) (*Filter, error) {
	if mode == ModeNone {
		return &Filter{mode: mode}, nil
	}
	l, err := ParseLevel(levelStr)
	if err != nil {
		return nil, err
	}
	return &Filter{mode: mode, level: l}, nil
}

// Allow returns true if the line should be kept based on its detected level.
func (f *Filter) Allow(line string) bool {
	if f.mode == ModeNone {
		return true
	}
	detected := detectLevel(line)
	switch f.mode {
	case ModeMin:
		return detected >= f.level
	case ModeExact:
		return detected == f.level
	}
	return true
}

func detectLevel(line string) Level {
	upper := strings.ToUpper(line)
	for _, name := range []string{"FATAL", "ERROR", "WARN", "INFO", "DEBUG"} {
		if strings.Contains(upper, name) {
			l, _ := ParseLevel(strings.ToLower(name))
			return l
		}
	}
	return LevelDebug
}
