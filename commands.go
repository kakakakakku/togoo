package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kakakakakku/togoo/command"
)

var GlobalFlags = []cli.Flag{}

var ForceFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "force, f",
		Usage: "Force initialize even if database already existed.",
	},
}

var Commands = []cli.Command{
	{
		Name:   "init",
		Usage:  "",
		Action: command.CmdInit,
		Flags:  ForceFlags,
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
		Flags:  []cli.Flag{},
	},
	{
		Name:   "done",
		Usage:  "",
		Action: command.CmdDone,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
