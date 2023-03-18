package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
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
		return errors.Errorf("%v", err)
	}

	var buf bytes.Buffer
	if err = json.Compact(&buf, b); err != nil {
		return errors.Errorf("%v", err)
	}

	b, err = json.Marshal(buf.String())
	if err != nil {
		return errors.Errorf("%v", err)
	}

	fmt.Print(string(b))

	return nil
}
