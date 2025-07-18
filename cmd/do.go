package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zeyadwaleed003/GopherDo/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do is the command for marking a task as done.",
	Run: func(cmd *cobra.Command, args []string) {
		ids := []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err == nil {
				ids = append(ids, id)
			}
		}

		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		for i, id := range ids {
			if id <= 0 || id > len(tasks) {
				continue
			}

			task := tasks[i]
			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark %v as completed. Error: %v\n", id, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
