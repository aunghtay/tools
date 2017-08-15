package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

func processArgs(arguments map[string]interface{}) {
	if arguments["run"].(bool) {
		fmt.Printf("Running Command %v ", arguments["<command>"])
		fmt.Printf("with Log Level %v\n", arguments["--log-level"])
	}
}

func main() {
	usage := `Command for launching program.
    
Usage:
    docopts-go-example -h | --help
    docopts-go-example run
        [--log-level=<log-level>]
        <command>

Options:
    -h --help                  Show this help message and exit.
    --version                  Show version and exit.

Logging Options:
    --log-level=<log-level>    Log level [default: info].`

	arguments, _ := docopt.Parse(usage, nil, true, "1.0.0", false)

	processArgs(arguments)

}
