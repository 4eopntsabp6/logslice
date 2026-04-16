package config

import "fmt"

// FieldConfig controls field extraction behaviour.
type FieldConfig struct {
	Mode string
	Key  string
}

// DefaultFieldConfig returns a FieldConfig with extraction disabled.
func DefaultFieldConfig() FieldConfig {
	return FieldConfig{
		Mode: "none",
	}
}

// Validate checks that the FieldConfig is coherent.
func (f FieldConfig) Validate() error {
	switch f.Mode {
	case "none":
		return nil
	case "json", "kv":
		if f.Key == "" {
			return fmt.Errorf("field: key is required when mode is %q", f.Mode)
		}
		return nil
	default:
		return fmt.Errorf("field: unknown mode %q (want none|json|kv)", f.Mode)
	}
}

// Enabled reports whether field extraction is active.
func (f FieldConfig) Enabled() bool {
	return f.Mode != "none"
}
