// Package pipeline wires together the input, filter, slicer, and output
// stages into a single reusable execution unit.
package pipeline

import (
	"fmt"
	"io"

	"github.com/yourorg/logslice/config"
	"github.com/yourorg/logslice/filter"
	"github.com/yourorg/logslice/input"
	"github.com/yourorg/logslice/output"
	"github.com/yourorg/logslice/slicer"
	"github.com/yourorg/logslice/stats"
)

// Pipeline holds all wired-up components needed to run a log-slice operation.
type Pipeline struct {
	cfg       *config.Config
	reader    io.ReadCloser
	formatter output.Formatter
	dest      io.WriteCloser
	filter    *filter.RegexFilter
	stats     *stats.Stats
}

// New constructs a Pipeline from the supplied Config, opening all required
// resources. The caller must call Close when finished.
func New(cfg *config.Config) (*Pipeline, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	reader, err := input.NewFileReader(cfg.InputPath)
	if err != nil {
		return nil, fmt.Errorf("open input: %w", err)
	}

	dest, err := output.OpenDestination(cfg.OutputPath)
	if err != nil {
		reader.Close()
		return nil, fmt.Errorf("open output: %w", err)
	}

	fmt, err := output.NewFormatter(cfg.Format)
	if err != nil {
		reader.Close()
		dest.Close()
		return nil, fmt.Errorf("create formatter: %w", err)
	}

	var rf *filter.RegexFilter
	if cfg.Pattern != "" {
		rf, err = filter.NewRegexFilter(cfg.Pattern)
		if err != nil {
			reader.Close()
			dest.Close()
			return nil, fmt.Errorf("compile pattern: %w", err)
		}
	}

	return &Pipeline{
		cfg:       cfg,
		reader:    reader,
		formatter: fmt,
		dest:      dest,
		filter:    rf,
		stats:     stats.New(),
	}, nil
}

// Run executes the full slice pipeline and returns collected statistics.
func (p *Pipeline) Run() (*stats.Stats, error) {
	results, err := slicer.Slice(p.reader, p.cfg, p.filter, p.stats)
	if err != nil {
		return nil, fmt.Errorf("slice: %w", err)
	}

	if err := slicer.Write(p.dest, p.formatter, results, p.cfg.Numbered); err != nil {
		return nil, fmt.Errorf("write: %w", err)
	}

	return p.stats, nil
}

// Close releases all resources held by the pipeline.
func (p *Pipeline) Close() error {
	err1 := p.reader.Close()
	err2 := p.dest.Close()
	if err1 != nil {
		return err1
	}
	return err2
}
