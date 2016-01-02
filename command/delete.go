package command

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codegangsta/cli"
)

// CmdDelete list all tasks.
func CmdDelete(c *cli.Context) {

	isAllMode := c.String("all")

	var id string
	if len(c.Args()) == 1 {
		id = c.Args()[0]
	}

	var s string
	var m string
	if isAllMode == "true" {
		s = "DELETE FROM todos"
		m = "Deleted all tasks"
	} else {
		s = "DELETE FROM todos WHERE id = " + id
		m = "Deleted task id = " + id
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
