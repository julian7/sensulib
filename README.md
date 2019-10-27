# sensulib

Base library for writing go checks.

## Goals

There are three goals for this project:

1. write a go library to help writing go check plugins.
2. provide a [magefile](https://magefile.org/) library to help publishing assets for Sensu-go.

## Usage

### Root command error handling

Check handlers use `panic` instead of `os.Exit` for breaking execution, for better testability. However, the main command has to recover from panic to properly exit with the provided exit code:

```golang
package main

import "github.com/julian7/sensulib"

func main() {
    defer sensulib.Recover()

    if err := rootCmd().Execute(); err != nil {
        sensulib.HandleError(err)
    }
}
```

### Build a sensu check command

This snippet creates a command with a configuration struct, therefore all the parameters will be readily available during run. `checkCmd()` returns a `*cobra.Command`, which then can be called or added to another command.

```golang
package main

import (
    "github.com/julian7/sensulib"
    "github.com/spf13/cobra"
)

type checkConfig struct {
    param1 string
    param2 float64
}

func (conf *checkConfig) Run(cmd *cobra.Command, []args string) error {
    // interesting things
    return nil
}

func checkCmd() *cobra.Command {
    config := &checkConfig{}
    cmd := sensulib.Command(
        config,
        "use",
        "short desc",
        "long desc",
    )
    flags := cmd.Flags()
    flags.StringVarP(&config.param1, "param1", "p", "", "first parameter")
    flags.Float64VarP(&config.param2, "param2, "P", 0.0, "second parameter")

    return cmd
}
```

### Errors

There are two error classes in sensulib: a single one, and an aggregate one. Both can be used to return from commands.

Single errors:

- `sensulib.Ok(error)`: no real error, it should report success.
- `sensulib.Warn(error)`: report Warning-level (1) error.
- `sensulib.Crit(error)`: report Critical-level (2) error.
- `sensulib.Unknown(error)`: report Unknown-level (3) error.

Aggregate errors:

- `sensulib.NewErrors()`: create a new aggregate.
- `sensulib.Errors.Add(err)`: add a new single error.
- `sensulib.Errors.Return(err)`: return with aggregated errors, or a default one, if specified.

### Humanize

Some numbers need to be converted to be human-readable. Some functions:

- `sensulib.SizeToHuman(uint64)`: converts bytes into IEC-formatted, rounded string.
- `sensulib.PercentToHuman(percent float64, precision int)`: converts percentage into a nice, rounded string.

## Legal

This project is licensed under [Blue Oak Model License v1.0.0](https://blueoakcouncil.org/license/1.0.0). It is not registered either at OSI or GNU, therefore GitHub is widely looking at the other direction. However, this is the license I'm most happy with: you can read and understand it with no legal degree, and there are no hidden or cryptic meanings in it.

The project is also governed with [Contributor Covenant](https://contributor-covenant.org/)'s [Code of Conduct](https://www.contributor-covenant.org/version/1/4/) in mind. I'm not copying it here, as a pledge for taking the verbatim version by the word, and we are not going to modify it in any way.

## Any issues?

Open a ticket, perhaps a pull request. We support [GitHub Flow](https://guides.github.com/introduction/flow/). You might want to [fork](https://guides.github.com/activities/forking/) this project first.
