package command

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"database/sql"

	"github.com/codegangsta/cli"
)

// CmdUpdate update task.
func CmdUpdate(c *cli.Context) {
	if len(c.Args()) <= 2 {
		fmt.Println("[ERROR] Must set task id and title")
		return
	}

	id := c.Args()[0]
	title := strings.Join(c.Args()[1:], " ")

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	now := strconv.FormatInt(time.Now().Unix(), 10)

	_, err = db.Exec("UPDATE todos SET title = '" + title + "', updated_at = " + now + " WHERE id = " + id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated task id = " + id)
}
