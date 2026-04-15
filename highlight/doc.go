// Package highlight provides optional ANSI terminal colour highlighting
// for logslice output.
//
// When a user supplies --highlight-pattern and --highlight-colour flags,
// every matched substring in a log line is wrapped with the corresponding
// ANSI escape sequence before the line is written to the output destination.
//
// Usage:
//
//	h, err := highlight.New(pattern, highlight.ParseColour(colourName), enabled)
//	if err != nil {
//		// handle invalid regex
//	}
//	formattedLine := h.Apply(rawLine)
//
// Highlighting is automatically skipped when writing to non-terminal
// destinations (files, pipes) if the caller sets enabled=false.
//
// Strip can be used to remove all ANSI codes from a string, which is
// useful when writing to file destinations or running tests.
package highlight
