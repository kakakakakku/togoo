package command

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

// CmdList list all tasks.
func CmdList(c *cli.Context) {
	if len(c.Args()) != 0 {
		return
	}

	isAllMode := c.String("all")

	var s string
	if isAllMode == "true" {
		s = "SELECT id, title, is_done FROM todos"
	} else {
		s = "SELECT id, title, is_done FROM todos WHERE is_done = 0"
	}

	db, err := sql.Open("sqlite3", dbPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(s)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	data := [][]string{}

	for rows.Next() {
		var id int
		var title string
		var isDone int
		rows.Scan(&id, &title, &isDone)
		data = append(data, []string{strconv.Itoa(id), title, doneLabel(isDone)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Title", "Status"})
	table.SetBorder(true)
	table.AppendBulk(data)
	table.Render()
}

func doneLabel(isDone int) string {
	if isDone == 0 {
		return "-"
	}
	return "Done"
}
