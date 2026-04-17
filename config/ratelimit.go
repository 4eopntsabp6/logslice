package config

import (
	"errors"

	"github.com/user/logslice/ratelimit"
)

// RateLimitConfig holds rate-limiting settings.
type RateLimitConfig struct {
	Mode  string `mapstructure:"ratelimit-mode"`
	Limit int    `mapstructure:"ratelimit-limit"`
}

// DefaultRateLimitConfig returns a disabled rate-limit config.
func DefaultRateLimitConfig() RateLimitConfig {
	return RateLimitConfig{
		Mode:  "none",
		Limit: 0,
	}
}

// Validate checks the RateLimitConfig for correctness.
func (c *RateLimitConfig) Validate() error {
	if _, err := ratelimit.ParseMode(c.Mode); err != nil {
		return err
	}
	if ratelimit.Mode(c.Mode) == ratelimit.ModeLines && c.Limit <= 0 {
		return errors.New("config: ratelimit limit must be > 0 when mode is lines")
	}
	return nil
}

// Enabled returns true when rate limiting is active.
func (c *RateLimitConfig) Enabled() bool {
	return ratelimit.Mode(c.Mode) != ratelimit.ModeNone
}
