package command

import (
	"database/sql"
	"github.com/codegangsta/cli"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func CmdInit(c *cli.Context) {
	os.Remove(db_path())

	db, err := sql.Open("sqlite3", db_path())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
		CREATE TABLE todos (
			id integer NOT NULL PRIMARY KEY,
			title text NOT NULL,
			is_done integer NOT NULL,
			created_at integer NOT NULL,
			updated_at integer NOT NULL
		);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
}
