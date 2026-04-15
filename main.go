package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// Version is set at build time via ldflags
	Version = "dev"

	// Flags
	startTime  string
	endTime    string
	pattern    string
	timestamp  string
	outputFile string
	followMode bool
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "logslice [flags] <logfile>",
	Short: "Extract and filter structured log ranges from large files",
	Long: `logslice is a CLI tool for extracting and filtering structured log ranges
from large files by time window or regex pattern.

Examples:
  # Extract logs between two timestamps
  logslice --start "2024-01-15  "2024-01-15 11:00:00" app.log

  # logs matchingice --pattern "ERROR|WARN" app.log

  # time window with pattern filter
  logslice --start "2024-01-15 10:00:00" --end "2024-01-15 11:00:00" --pattern "ERROR" app.log

  # Write output to a file
  logslice --start "2024-01-15 10:00:00" --output filtered.log app.log`,
	Version:      Version,
	Args:         cobra.ExactArgs(1),
	RunE:         runSlice,
	SilenceUsage: true,
}

func init() {
	rootCmd.Flags().StringVarP(&startTime, "start", "s", "", "Start of time window (e.g. \"2024-01-15 10:00:00\")")
	rootCmd.Flags().StringVarP(&endTime, "end", "e", "", "End of time window (e.g. \"2024-01-15 11:00:00\")")
	rootCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "Regex pattern to filter log lines")
	rootCmd.Flags().StringVarP(&timestamp, "timestamp-format", "t", "", "Custom timestamp format (Go time layout or named preset)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Write output to file instead of stdout")
	rootCmd.Flags().BoolVarP(&followMode, "follow", "f", false, "Follow log file for new matching entries (like tail -f)")
}

func runSlice(cmd *cobra.Command, args []string) error {
	logFile := args[0]

	// Validate that at least one filter is provided
	if startTime == "" && endTime == "" && pattern == "" {
		return fmt.Errorf("at least one of --start, --end, or --pattern must be specified")
	}

	// Validate the log file exists
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		return fmt.Errorf("log file not found: %s", logFile)
	}

	// Build slicer configuration
	cfg := &SliceConfig{
		InputFile:       logFile,
		StartTime:       startTime,
		EndTime:         endTime,
		Pattern:         pattern,
		TimestampFormat: timestamp,
		OutputFile:      outputFile,
		Follow:          followMode,
	}

	slicer, err := NewSlicer(cfg)
	if err != nil {
		return fmt.Errorf("failed to initialize slicer: %w", err)
	}
	defer slicer.Close()

	if err := slicer.Run(); err != nil {
		return fmt.Errorf("slice failed: %w", err)
	}

	return nil
}
