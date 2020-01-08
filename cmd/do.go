package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Done  task")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
