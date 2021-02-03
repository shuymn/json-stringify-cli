package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/shuymn/json-stringify-cli/cli"
)

const (
	ExitCodeOK = iota
	ExitCodeErr
)

func main() {
	os.Exit(_main())
}

func _main() int {
	debug := kingpin.Flag("debug", "Enable debug log").Bool()
	fp := kingpin.Arg("json-path", "Specify json file path").Required().String()

	kingpin.Parse()

	c := cli.New(*fp)
	if err := c.Run(); err != nil {
		format := "%v\n"
		if *debug {
			// print a log message with stack trace when debug mode is enable
			format = "%+v\n"
		}
		fmt.Printf(format, err)
		return ExitCodeErr
	}
	return ExitCodeOK
}
