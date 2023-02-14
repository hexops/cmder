package cmder

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Command is a subcommand handler and its associated flag set.
type Command struct {
	// FlagSet is the flag set for the command.
	FlagSet *flag.FlagSet

	// Aliases for the command name.
	Aliases []string

	// Handler is the function that is invoked to handle this command.
	//
	// The error types *UsageError and *ExitCodeError have special meaning, consult their docs for
	// details.
	Handler func(args []string) error

	// A flagSet.Usage function to invoke when e.g. the -h flag is specified. If nil, a default one
	// is used.
	UsageFunc func()
}

// matches tells if the given name matches this command or one of its aliases.
func (c *Command) matches(name string) bool {
	if name == c.FlagSet.Name() {
		return true
	}
	for _, alias := range c.Aliases {
		if name == alias {
			return true
		}
	}
	return false
}

// Commander represents a command with a list of subcommands.
type Commander []*Command

// Run runs a subcommand of the command described by the input flagSet (e.g. flag.CommandLine).
//
// cmdName and usageText should describe your command, not the subcommand. Consult "go help" for
// inspiration when writing your own usageText.
//
// A special "help" command is registered automatically, which acts the same as the `-h` flag.
func (c Commander) Run(flagSet *flag.FlagSet, cmdName, usageText string, args []string) {
	// Parse flags.
	flagSet.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usageText)
	}
	if !flagSet.Parsed() {
		// We assume flag.ExitOnError is in use.
		if err := flagSet.Parse(args); err != nil {
			panic(fmt.Sprintf("command should use flag.ExitOnError: error: %s", err))
		}
	}

	// Print usage if the command is "help".
	if flagSet.Arg(0) == "help" || flagSet.NArg() == 0 {
		flagSet.Usage()
		os.Exit(0)
	}

	// Configure default usage funcs for commands.
	for _, cmd := range c {
		cmd := cmd
		if cmd.UsageFunc != nil {
			cmd.FlagSet.Usage = cmd.UsageFunc
			continue
		}
		cmd.FlagSet.Usage = func() {
			fmt.Fprintf(flag.CommandLine.Output(), "Usage of '%s %s':\n", cmdName, cmd.FlagSet.Name())
			cmd.FlagSet.PrintDefaults()
		}
	}

	// Find the subcommand to execute.
	name := flagSet.Arg(0)
	for _, cmd := range c {
		if !cmd.matches(name) {
			continue
		}

		// Execute the subcommand.
		if err := cmd.Handler(flagSet.Args()[1:]); err != nil {
			if _, ok := err.(*UsageError); ok {
				log.Println(err)
				cmd.FlagSet.Usage()
				os.Exit(2)
			}
			if e, ok := err.(*ExitCodeError); ok {
				if e.Err != nil {
					log.Println(e)
				}
				os.Exit(e.ExitCode)
			}
			log.Fatal(err)
		}
		os.Exit(0)
	}
	log.Printf("%s: unknown subcommand %q", cmdName, name)
	log.Fatalf("Run '%s help' for usage.", cmdName)
}

// UsageError is an error type that subcommands can return in order to signal
// that a usage error has occurred.
type UsageError struct {
	// Err is the error to log when exiting.
	Err error
}

// Error implements the error interface.
func (e *UsageError) Error() string { return e.Err.Error() }

// ExitCodeError is an error type that subcommands can return in order to
// specify the exact exit code.
type ExitCodeError struct {
	// Err is the error to log when exiting.
	Err error

	// ExitCode is the exit status code to use in the call to os.Exit.
	ExitCode int
}

// Error implements the error interface.
func (e *ExitCodeError) Error() string { return e.Err.Error() }
