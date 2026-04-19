package pipeline

import (
	"fmt"
	"io"

	"github.com/user/logslice/config"
	"github.com/user/logslice/slicer"
	"github.com/user/logslice/stats"
)

// Runner executes a configured pipeline end-to-end.
type Runner struct {
	cfg      *config.Config
	pipeline *Pipeline
	reporter *stats.Reporter
}

// NewRunner constructs a Runner from the given config, writing output to w.
func NewRunner(cfg *config.Config, w io.Writer) (*Runner, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config must not be nil")
	}
	p, err := New(cfg)
	if err != nil {
		return nil, fmt.Errorf("pipeline init: %w", err)
	}
	return &Runner{
		cfg:      cfg,
		pipeline: p,
		reporter: stats.NewReporter(w),
	}, nil
}

// Run executes the pipeline: slices input, writes results, and reports stats.
func (r *Runner) Run(out io.Writer) error {
	if out == nil {
		return fmt.Errorf("output writer must not be nil")
	}

	results, st, err := r.pipeline.Execute()
	if err != nil {
		return fmt.Errorf("pipeline execute: %w", err)
	}

	wOpts := slicer.WriteOptions{
		Numbered: r.cfg.Numbered,
	}
	if err := slicer.Write(out, results, wOpts); err != nil {
		return fmt.Errorf("write results: %w", err)
	}

	if r.cfg.ShowStats {
		r.reporter.Report(st)
	}

	return nil
}
