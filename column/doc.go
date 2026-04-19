// Package column provides column extraction for log lines.
//
// It splits each line by a configurable delimiter and returns the value
// at a given zero-based column index. When mode is "none" the line is
// passed through unchanged.
//
// Supported modes:
//
//	none    – pass-through, no extraction performed
//	extract – split by delimiter and return column at index
//
// Example usage:
//
//	e, err := column.New(column.ModeExtract, ",", 2)
//	if err != nil { ... }
//	result := e.Apply(line)
package column
