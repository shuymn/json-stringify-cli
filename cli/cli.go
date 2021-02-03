package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/xerrors"
)

type CLI struct {
	path string
}

func New(path string) *CLI {
	return &CLI{path: path}
}

func (c *CLI) Run() error {
	b, err := ioutil.ReadFile(c.path)
	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	var buf bytes.Buffer
	if err = json.Compact(&buf, b); err != nil {
		return xerrors.Errorf("%v", err)
	}

	b, err = json.Marshal(buf.String())
	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	fmt.Print(string(b))

	return nil
}
