package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type CLI struct {
	path   string
	stdout io.Writer
}

func New(path string, stdout io.Writer) *CLI {
	return &CLI{
		path:   path,
		stdout: stdout,
	}
}

func (c *CLI) Run() error {
	b, err := os.ReadFile(c.path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var buf bytes.Buffer
	if err = json.Compact(&buf, b); err != nil {
		return fmt.Errorf("failed to compact JSON: %w", err)
	}

	b, err = json.Marshal(buf.String())
	if err != nil {
		return fmt.Errorf("failed to marshal JSON string: %w", err)
	}

	_, err = fmt.Fprint(c.stdout, string(b))
	if err != nil {
		return fmt.Errorf("failed to write to stdout: %w", err)
	}

	return nil
}
