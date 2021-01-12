package main

import (
	"flag"
	"log"
	"os"

	"github.com/hexops/cmder"
)

// commands contains all registered subcommands.
var commands cmder.Commander

// Our help text for this command.
// Consult "go help" for inspiration on how to word yours.
var usageText = `kurl is a tool that makes HTTP requests.

Usage:

	kurl <command> [arguments]

The commands are:

	get    perform HTTP GET requests

Use "kurl <command> -h" for more information about a command.
`

func main() {
	// Configure logging if desired.
	log.SetFlags(0)
	log.SetPrefix("")

	commands.Run(flag.CommandLine, "kurl", usageText, os.Args[1:])
}
