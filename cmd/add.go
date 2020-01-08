package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to your task", task)
		// for i, arg := range args {
		// 	fmt.Println(i, ":", arg)
		// }
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
