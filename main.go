package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/zeyadwaleed003/GopherDo/cmd"
	"github.com/zeyadwaleed003/GopherDo/db"
)

func main() {
	homeDir, _ := homedir.Dir()
	dbPath := filepath.Join(homeDir, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
