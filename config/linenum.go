package config

import (
	"fmt"
	"github.com/user/logslice/linenum"
)

// LineNumConfig holds line number filter settings.
type LineNumConfig struct {
	Mode  string
	From  int
	To    int
	Lines []int
}

// DefaultLineNumConfig returns a disabled line number config.
func DefaultLineNumConfig() LineNumConfig {
	return LineNumConfig{Mode: "none"}
}

// Validate checks the LineNumConfig for consistency.
func (c LineNumConfig) Validate() error {
	m, err := linenum.ParseMode(c.Mode)
	if err != nil {
		return err
	}
	if m == linenum.ModeRange && c.From < 1 {
		return fmt.Errorf("linenum: from must be >= 1")
	}
	if m == linenum.ModeRange && c.To > 0 && c.To < c.From {
		return fmt.Errorf("linenum: to must be >= from")
	}
	if m == linenum.ModeList && len(c.Lines) == 0 {
		return fmt.Errorf("linenum: list mode requires at least one line number")
	}
	return nil
}

// Enabled returns true when line number filtering is active.
func (c LineNumConfig) Enabled() bool {
	return c.Mode != "" && c.Mode != "none"
}
