package config

import (
	"fmt"

	"github.com/user/logslice/bookmark"
)

// BookmarkConfig holds configuration for the bookmark feature.
type BookmarkConfig struct {
	Mode    string
	Tag     string
	Indices []int
	Pattern string
}

// DefaultBookmarkConfig returns a BookmarkConfig with safe defaults.
func DefaultBookmarkConfig() BookmarkConfig {
	return BookmarkConfig{
		Mode: "none",
		Tag:  "[bookmark] ",
	}
}

// Validate checks that the BookmarkConfig is consistent.
func (c BookmarkConfig) Validate() error {
	mode, err := bookmark.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if mode == bookmark.ModeNone {
		return nil
	}
	if c.Tag == "" {
		return fmt.Errorf("bookmark: tag must not be empty when mode is active")
	}
	switch mode {
	case bookmark.ModeIndex:
		if len(c.Indices) == 0 {
			return fmt.Errorf("bookmark: indices required for index mode")
		}
	case bookmark.ModeRegex:
		if c.Pattern == "" {
			return fmt.Errorf("bookmark: pattern required for regex mode")
		}
	}
	return nil
}

// Enabled returns true when bookmarking is active.
func (c BookmarkConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
