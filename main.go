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
	fmt.Println(homeDir)
	dbPath := filepath.Join(homeDir, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
