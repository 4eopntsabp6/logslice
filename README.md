# logslice

A CLI tool for extracting and filtering structured log ranges from large files by time window or regex.

---

## Installation

```bash
go install github.com/yourusername/logslice@latest
```

Or build from source:

```bash
git clone https://github.com/yourusername/logslice.git
cd logslice
go build -o logslice .
```

---

## Usage

```bash
# Extract logs between two timestamps
logslice --from "2024-01-15T10:00:00" --to "2024-01-15T10:30:00" app.log

# Filter logs matching a regex pattern
logslice --pattern "ERROR|WARN" app.log

# Combine time window and pattern filter
logslice --from "2024-01-15T10:00:00" --to "2024-01-15T11:00:00" --pattern "user_id=42" app.log

# Write output to a file
logslice --from "2024-01-15T10:00:00" --to "2024-01-15T10:30:00" app.log -o output.log
```

### Flags

| Flag        | Description                              |
|-------------|------------------------------------------|
| `--from`    | Start of the time window (RFC3339)       |
| `--to`      | End of the time window (RFC3339)         |
| `--pattern` | Regex pattern to match log lines         |
| `-o`        | Output file (defaults to stdout)         |
| `--format`  | Log timestamp format (default: RFC3339)  |

---

## Contributing

Pull requests and issues are welcome. Please open an issue before submitting large changes.

---

## License

MIT © 2024 yourusername