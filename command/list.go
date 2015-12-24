package command

import (
	"database/sql"
	"fmt"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CmdList(c *cli.Context) {
	if len(c.Args()) != 0 {
		return
	}

	db, err := sql.Open("sqlite3", "./togoo.db")
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
		var is_done int
		rows.Scan(&id, &title, &is_done)
		fmt.Printf("| %3d | %-50s | %-7s |\n", id, title, done_label(is_done))
	}

	fmt.Printf("| %-3s | %-50s | %-7s |\n", "---", "--------------------------------------------------", "-------")
}

func done_label(is_done int) string {
	if is_done == 0 {
		return "-"
	} else {
		return "Done"
	}
}
