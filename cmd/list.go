package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/zeyadwaleed003/GopherDo/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete :)")
			return
		}

		for i, task := range tasks {
			fmt.Printf("%v. %v\n", i+1, task.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
