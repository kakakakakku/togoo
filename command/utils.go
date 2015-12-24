package command

import (
	"github.com/mitchellh/go-homedir"
)

func db_path() string {
	home, err := homedir.Dir()
	if err != nil {
	}
	return home + "/togoo.db"
}
