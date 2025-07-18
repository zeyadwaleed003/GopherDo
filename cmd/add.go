package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Adding the task: \"%v\" to the task list.\n", task)
	},
}

// This go init function runs automatically before the main() function
func init() {
	rootCmd.AddCommand(addCmd)
}
