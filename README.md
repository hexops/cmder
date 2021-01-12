# cmder - lightweight Go pattern for writing CLIs <a href="https://hexops.com"><img align="right" alt="Hexops logo" src="https://raw.githubusercontent.com/hexops/media/master/readme.svg"></img></a>

<a href="https://pkg.go.dev/github.com/hexops/cmder"><img src="https://pkg.go.dev/badge/badge/github.com/hexops/cmder.svg" alt="Go Reference" align="right"></a>
  
[![Go CI](https://github.com/hexops/cmder/workflows/Go%20CI/badge.svg)](https://github.com/hexops/cmder/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/hexops/cmder)](https://goreportcard.com/report/github.com/hexops/cmder)

Cmder is a ~100 LOC pattern that has been used as the foundation of all the Go CLIs [I've](https://twitter.com/slimsag) written (including [the Sourcegraph CLI](https://github.com/sourcegraph/src-cli/blob/1af97e4f78819ffd042ef000d964090dbb65268f/cmd/src/cmd.go#L1-L123).)

I've [often just suggested others simply copy the pattern](https://twitter.com/slimsag/status/1330924665544404994) as it is so lightweight. Now you can import it as a Go package, which helps to document the pattern you're using, and avoids temptation to make it more complex than needed.

* Mimics what the official `go` tool does internally.
* Merely builds upon the `flag` package to support subcommands.
* Supports subcommand flags, subcommands of subcommands, etc.

## Example

An example/fictional HTTP tool called `kurl` is provided in the `example/kurl` directory.

## Usage

The idea is simple, import the package:

```Go
import "github.com/hexops/cmder"
```

Declare a list of your subcommands:

```Go
// commands contains all registered subcommands.
var commands cmder.Commander
```

Append a few subcommands like so (usually one `init` function per subcommand):

```Go
flagSet := flag.NewFlagSet("foo", flag.ExitOnError)
commands = append(commands, &cmder.Command{
    FlagSet: flagSet,
    Handler: func(args []string) error {
        _ = flagSet.Parse(args)
        return nil
    },
})
```

In your `main` function, call:

```Go
commands.Run(flag.CommandLine, commandName, usageText, os.Args[1:])
```

* Consult `go help` for inspiration on how to write your `usageText`.
* Register subcommand flags by using e.g. `flagSet.Bool` (as you would've if using the Go `flag` package otherwise.)
* Need subcommands in your subcommands? Declare another set of `commands` and simply call the `Run` method inside your subcommand `Handler`.

Consult the [API documentation](https://pkg.go.dev/github.com/hexops/cmder) for more information.

## Project status

We're open to considering improvements, but since this pattern has been in use in various CLIs over the past 3-4 years, we likely won't make any major changes to the API or introduce new features. The aim is to keep it minimal and simple.
