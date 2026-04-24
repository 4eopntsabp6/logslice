// Package entropy implements Shannon entropy-based line filtering for logslice.
//
// Lines are scored using the standard Shannon entropy formula (bits per symbol).
// Two modes are supported:
//
//   - above: keep lines whose entropy is at or above the threshold.
//     Useful for surfacing high-randomness lines such as encoded payloads
//     or cryptographic tokens.
//
//   - below: keep lines whose entropy is at or below the threshold.
//     Useful for filtering out noisy, highly random content and retaining
//     human-readable log messages.
//
// A threshold of 0 disables the filter (mode "none").
package entropy
