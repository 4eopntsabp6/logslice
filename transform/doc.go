// Package transform provides line-level text transformations for log output.
//
// Supported modes:
//
//   - none      — no transformation (default)
//   - upper     — convert line to uppercase
//   - lower     — convert line to lowercase
//   - trim      — strip leading and trailing whitespace
//
// Example usage:
//
//	tr := transform.New(transform.ModeUpper)
//	fmt.Println(tr.Apply("info: started")) // INFO: STARTED
package transform
