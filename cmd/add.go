package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/todo-cli/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)

		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task", task)
		// for i, arg := range args {
		// 	fmt.Println(i, ":", arg)
		// }
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
