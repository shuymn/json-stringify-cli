# json-stringify-cli

CLI tool that reads JSON from a file or stdin, compacts it (removes whitespace), and outputs a single JSON string literal containing that JSON. Useful when you need the JSON embedded as a string value.

## Features

- Read JSON from files or standard input
- Compact JSON output (removes unnecessary whitespace)
- Output as JSON string literal with proper escaping
- Zero external dependencies (uses only Go standard library)
- Helpful error messages with optional debug mode

## Installation

```bash
go install github.com/shuymn/json-stringify-cli/cmd/json-stringify@latest
```

Requires Go 1.21.6+.

## Usage

```bash
json-stringify [--debug] [json-path]
```

- `--debug`: enable verbose error formatting when available (no stack traces by default)
- `[json-path]`: path to the input JSON file (optional)
  - If omitted or `-`, reads from stdin
  - If provided, reads from the specified file

## Examples

### Reading from a file

Given `sample.json`:

```json
{
  "color": "red",
  "value": "#f00"
}
```

Run:

```bash
json-stringify sample.json
```

Output:

```
"{\"color\":\"red\",\"value\":\"#f00\"}"
```

### Reading from stdin

Using pipe:

```bash
echo '{"name": "test", "value": 123}' | json-stringify
```

Output:

```
"{\"name\":\"test\",\"value\":123}"
```

Using redirection:

```bash
json-stringify < sample.json
```

Or explicitly with `-`:

```bash
cat sample.json | json-stringify -
```

### Processing API responses

```bash
curl -s https://api.example.com/data | json-stringify
```

### Note on output format

The tool outputs a JSON string literal, so internal quotes are escaped. For example:

- Input: `{"key": "value"}`
- Output: `"{\"key\":\"value\"}"`

- Input: `[1, 2, 3]`
- Output: `"[1,2,3]"`

## Exit Status

- `0` on success
- `1` on error (e.g., invalid JSON, file not found, or stdin read error)
  - Errors are written to stderr
  - Use `--debug` to enable verbose error formatting when supported by underlying errors; this tool does not add stack traces by itself

## Build From Source

```bash
make            # builds bin/json-stringify
```

Or directly:

```bash
go build -ldflags="-s -w" -o bin/json-stringify cmd/json-stringify/main.go
```

## Test

```bash
make test       # runs `go test -race ./...`
```

## License

MIT
