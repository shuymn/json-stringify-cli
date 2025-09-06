# json-stringify-cli

CLI tool that reads a JSON file, compacts it (removes whitespace), and outputs a single JSON string literal containing that JSON. Useful when you need the JSON embedded as a string value.

## Installation

```bash
go install github.com/shuymn/json-stringify-cli/cmd/json-stringify@latest
```

Requires Go 1.21+.

## Usage

```bash
json-stringify [--debug] <json-path>
```

- `--debug`: print errors with stack traces.
- `<json-path>`: path to the input JSON file (required).

## Examples

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

Output (to stdout):

```
{"color":"red","value":"#f00"}
```

Note: the tool prints a JSON string literal, so quotes inside are escaped. For an array like `[100, 500, 300, 200, 400]`, the output is:

```
[100,500,300,200,400]
```

## Exit Status

- `0` on success.
- `1` on error (e.g., invalid JSON or file not found). Use `--debug` to include a stack trace in the error message.

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
