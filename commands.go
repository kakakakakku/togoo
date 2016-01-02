package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kakakakakku/togoo/command"
)

// GlobalFlags is global option for all command.
var GlobalFlags = []cli.Flag{}

var forceFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "force, f",
		Usage: "Force initialize even if database already existed.",
	},
}

var allFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "all, a",
		Usage: "List all tasks.",
	},
}

// Commands are slice of command.
var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "",
		Action: command.CmdInit,
		Flags:  forceFlags,
	},
	{
		Name:   "add",
		Usage:  "",
		Action: command.CmdAdd,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  allFlags,
	},
	{
		Name:   "done",
		Usage:  "",
		Action: command.CmdDone,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound is custom error.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
