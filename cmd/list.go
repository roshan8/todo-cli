package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks in your todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Listing tasks")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
