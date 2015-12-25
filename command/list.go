package command

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
)

// CmdList list all tasks.
func CmdList(c *cli.Context) {
	if len(c.Args()) != 0 {
		return
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, title, is_done FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Printf("| %-3s | %-50s | %-7s |\n", "---", "--------------------------------------------------", "-------")
	fmt.Printf("| %-3s | %-50s | %-7s |\n", "No", "Title", "Status")
	fmt.Printf("| %-3s | %-50s | %-7s |\n", "---", "--------------------------------------------------", "-------")

	for rows.Next() {
		var id int
		var title string
		var isDone int
		rows.Scan(&id, &title, &isDone)
		fmt.Printf("| %3d | %-50s | %-7s |\n", id, title, doneLabel(isDone))
	}

	fmt.Printf("| %-3s | %-50s | %-7s |\n", "---", "--------------------------------------------------", "-------")
}

func doneLabel(isDone int) string {
	if isDone == 0 {
		return "-"
	}
	return "Done"
}
