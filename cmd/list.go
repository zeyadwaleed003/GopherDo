package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks in your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("List command is working.\n")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
