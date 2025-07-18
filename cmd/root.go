package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a simple command-line task manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
