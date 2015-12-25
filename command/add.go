package command

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"strconv"
	"time"
)

// CmdAdd add new task.
func CmdAdd(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("[ERROR] Must set task title")
		return
	}

	title := c.Args()[0]
	fmt.Println("Added new task :", title)

	db, err := sql.Open("sqlite3", db_path())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("INSERT INTO todos(title, is_done, created_at, updated_at) values('" + title + "', 0, " + now + ", " + now + ")")
	if err != nil {
		log.Fatal(err)
	}
}
