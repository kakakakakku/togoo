package command

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
)

func CmdDone(c *cli.Context) {

	if len(c.Args()) != 1 {
		return
	}

	id := c.Args()[0]
	fmt.Printf("Task %s is done\n", id)

	db, err := sql.Open("sqlite3", "./togoo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE todos SET is_done = 1 WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}
}
