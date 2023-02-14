package main

import (
	"flag"
	"fmt"

	"github.com/hexops/cmder"
)

func init() {
	const usage = `
Examples:

  Perform an HTTPS GET request on a URL:

    $ kurl get https://google.com

  Include verbose output:

    $ kurl get --verbose https://google.com
`

	// Parse flags for our subcommand.
	flagSet := flag.NewFlagSet("get", flag.ExitOnError)
	verboseFlag := flagSet.Bool("verbose", true, "include verbose information")

	// Handles calls to our subcommand.
	handler := func(args []string) error {
		_ = flagSet.Parse(args)

		// do something with args and *verboseFlag
		fmt.Println("subcommand called with", flagSet.Args(), "verbose?", *verboseFlag)

		// return &cmder.UsageError{} if usage text should be printed.

		// return &cmder.ExitCodeError{ExitStatus: ...} if a specific exit
		// code should be returned.

		// return any Go error if the command failed.
		return nil
	}

	// Register the command.
	commands = append(commands, &cmder.Command{
		FlagSet: flagSet,
		Aliases: []string{"fetch"},
		Handler: handler,
		UsageFunc: func() {
			fmt.Fprintf(flag.CommandLine.Output(), "Usage of 'kurl %s':\n", flagSet.Name())
			flagSet.PrintDefaults()
			fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usage)
		},
	})
}
