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
	flag.BoolVar(&debug, "debug", false, "Enable verbose error formatting (when available)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] [json-path]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nRead JSON from a file or stdin and output as a JSON string.\n")
		fmt.Fprintf(os.Stderr, "If json-path is omitted or '-', read from stdin.\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// If no args are provided and stdin is a TTY (i.e., not piped),
	// show help instead of blocking on stdin.
	if len(flag.Args()) == 0 {
		if fi, err := os.Stdin.Stat(); err == nil {
			if (fi.Mode() & os.ModeCharDevice) != 0 {
				flag.Usage()
				return ExitCodeErr
			}
		}
	}

	args := flag.Args()
	var fp string
	if len(args) < 1 {
		// No arguments with piped stdin: read from stdin
		fp = ""
	} else {
		fp = args[0]
	}

	c := cli.New(fp, os.Stdin, os.Stdout)
	if err := c.Run(); err != nil {
		format := "%v\n"
		if debug {
			// print a log message with stack trace when debug mode is enable
			format = "%+v\n"
		}
		_, err2 := fmt.Fprintf(os.Stderr, format, err)
		if err2 != nil {
			err = fmt.Errorf("failed to write to stderr: %w", err2)
			panic(err) // idk what to do if writing to stderr fails.
		}
		return ExitCodeErr
	}
	return ExitCodeOK
}
