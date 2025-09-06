package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/shuymn/json-stringify-cli/cli"
)

const (
	ExitCodeOK = iota
	ExitCodeErr
)

func main() {
	os.Exit(run())
}

func run() int {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "Enable debug log")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "Error: json-path argument is required\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <json-path>\n", os.Args[0])
		flag.PrintDefaults()
		return ExitCodeErr
	}

	fp := args[0]

	c := cli.New(fp, os.Stdout)
	if err := c.Run(); err != nil {
		format := "%v\n"
		if debug {
			// print a log message with stack trace when debug mode is enable
			format = "%+v\n"
		}
		fmt.Fprintf(os.Stdout, format, err)
		return ExitCodeErr
	}
	return ExitCodeOK
}
