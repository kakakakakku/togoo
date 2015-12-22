package command

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func CmdAdd(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("[ERROR] Must set task title")
		return
	}

	title := c.Args()[0]
	fmt.Println("Added new task :", title)
}
