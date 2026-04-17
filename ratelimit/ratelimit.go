// Package ratelimit provides line-rate throttling for log output.
package ratelimit

import (
	"errors"
	"time"
)

// Mode controls whether rate limiting is active.
type Mode string

const (
	ModeNone   Mode = "none"
	ModeLines  Mode = "lines"
)

// ParseMode converts a string to a Mode.
func ParseMode(s string) (Mode, error) {
	switch Mode(s) {
	case ModeNone, ModeLines:
		return Mode(s), nil
	}
	return "", errors.New("ratelimit: unknown mode: " + s)
}

// Limiter throttles output to at most N lines per second.
type Limiter struct {
	mode     Mode
	interval time.Duration
	last     time.Time
	count    int
	limit    int
}

// New creates a Limiter. limit is max lines per second; ignored when mode is none.
func New(mode Mode, limit int) (*Limiter, error) {
	if mode == ModeNone {
		return &Limiter{mode: mode}, nil
	}
	if limit <= 0 {
		return nil, errors.New("ratelimit: limit must be > 0")
	}
	return &Limiter{
		mode:     mode,
		limit:    limit,
		interval: time.Second,
		last:     time.Now(),
	}, nil
}

// Allow returns true if the line should be emitted, false if it should be dropped.
func (l *Limiter) Allow() bool {
	if l.mode == ModeNone {
		return true
	}
	now := time.Now()
	if now.Sub(l.last) >= l.interval {
		l.last = now
		l.count = 0
	}
	if l.count < l.limit {
		l.count++
		return true
	}
	return false
}
