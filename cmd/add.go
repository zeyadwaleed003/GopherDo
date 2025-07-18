package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeyadwaleed003/GopherDo/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if len(task) == 0 {
			fmt.Println("Please provide a task to add 😠")
			return
		}

		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}

		fmt.Printf("Added \"%v\" to your task list.\n", task)
	},
}

// This go init function runs automatically before the main() function
func init() {
	rootCmd.AddCommand(addCmd)
}
