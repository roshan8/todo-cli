package main

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/todo-cli/cmd"
	"github.com/todo-cli/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	cmd.RootCmd.Execute()
}
