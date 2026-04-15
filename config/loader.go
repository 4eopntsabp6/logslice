package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"logslice/parser"
)

// BindFlags registers all logslice flags onto cmd and returns a *Config
// populated lazily when the command runs. Call cfg.Validate() inside
// the command's RunE before use.
func BindFlags(cmd *cobra.Command) *Config {
	cfg := &Config{}

	cmd.Flags().StringVarP(&cfg.InputPath, "file", "f", "", "input log file (default: stdin)")
	cmd.Flags().StringVarP(&cfg.OutputDest, "output", "o", "stdout", "output destination: stdout, stderr, or a file path")
	cmd.Flags().StringVar(&cfg.OutputFormat, "format", "plain", "output format: plain, json, csv")
	cmd.Flags().StringVar(&cfg.Pattern, "pattern", "", "regex pattern to filter log lines")
	cmd.Flags().BoolVarP(&cfg.Numbered, "numbered", "n", false, "prefix each line with its original line number")
	cmd.Flags().BoolVarP(&cfg.Summary, "summary", "s", false, "print a summary after processing")

	// Raw string flags for time arguments; resolved in LoadTimes.
	var fromStr, toStr string
	cmd.Flags().StringVar(&fromStr, "from", "", "start of (RFC3339 or relative, e.g. -1h)")
	cmd.Flags().StringVar(&toStr, "to", "", "end of time window (RFC3339 or relative)")

	// Use PersistentPreRunE to parse time strings before RunE fires.
	prev := cmd.PersistentPreRunE
	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if prev != nil {
			if err := prev(cmd, args); err != nil {
				return err
			}
		}
		return loadTimes(cfg, fromStr, toStr)
	}

	return cfg
}

// loadTimes parses the raw --from / --to strings into cfg.From and cfg.To.
func loadTimes(cfg *Config, fromStr, toStr string) error {
	if fromStr != "" {
		t, err := parser.ParseTimeArg(fromStr)
		if err != nil {
			return fmt.Errorf("--from: %w", err)
		}
		cfg.From = t
	}
	if toStr != "" {
		t, err := parser.ParseTimeArg(toStr)
		if err != nil {
			return fmt.Errorf("--to: %w", err)
		}
		cfg.To = t
	}
	return nil
}
